package engine

import (
	"io"

	"github.com/Shopify/go-lua"
	"golang.org/x/net/context"
)

var _ Program = &LuaProgram{}

type LuaProgram struct {
	vm *lua.State

	// program state
	configured bool
	steps      []string

	info  func(*lua.State) int
	fatal func(*lua.State) int

	Stdout io.Writer
}

func Lua(source io.Reader, out io.Writer) (*LuaProgram, error) {
	l := lua.NewState()

	prgm := &LuaProgram{
		vm:     l,
		Stdout: out,
	}
	// setup the `step` hooks
	configureLua(prgm, l)
	l.Register("info", func(l *lua.State) int { return prgm.info(l) })
	l.Register("fatal", func(l *lua.State) int { return prgm.fatal(l) })

	// load the source
	if err := l.Load(source, "", ""); err != nil {
		return prgm, err
	}
	// invoke the program to prepare the steps
	if err := l.ProtectedCall(0, 0, 0); err != nil {
		return prgm, err
	}

	prgm.configured = true

	return prgm, nil
}

func (prgm *LuaProgram) Execute(ctx context.Context) error {

	l := prgm.vm

	runNext := true
	currentStep := "<not a step>"

	prgm.info = func(l *lua.State) int {
		msg := lua.CheckString(l, 1)
		fmt.Fprintf(prgm.Stdout, "[INFO] %s: %s", currentStep, msg)
		return 0
	}
	prgm.fatal = func(l *lua.State) int {
		msg := lua.CheckString(l, 1)
		fmt.Fprintf(prgm.Stdout, "[FATAL] %s: %s", currentStep, msg)
		return 0
	}

	reporter := func(stepName string) bool {
		currentStep = stepName
		select {
		default:
			return runNext
		case <-ctx.Done():
			return false
		}
	}

	return prgm.runSteps(reporter)
}

func (prgm *LuaProgram) runSteps(reporter func(step string) bool) error {
	l := prgm.vm

	// bring the step table on the stack
	l.Global("step")
	// prepare a table to hold the results of calling
	// each step
	l.NewTable()
	defer l.Pop(2) // cleanup the step + table

	for _, stepName := range prgm.steps {
		if !reporter(stepName) {
			// stop running
			break
		}

		i := l.Top()
		// pull the func at `stepName` out of the table
		l.Field(-i, stepName)
		// copy the table that holds results of the last step
		// as an argument for the func about to be invoked
		l.PushValue(-i)
		// remove the old copy of the result
		l.Remove(-(i + 1))
		// now that we have the:
		//   - argument
		//   - function
		//   - step-table
		// we can invoke the function with the argument
		err := l.ProtectedCall(1, 1, 0) // 1 argument, with 1 return value
		if err != nil {
			return &StepError{Step: stepName, Err: err}
		}
		if l.Top() != 2 {
			lua.Errorf(l, "step %q needs to return exactly 1 argument", stepName)
		}
	}
	return nil
}

func (prgm *LuaProgram) registerStep(l *lua.State) int {
	if prgm.configured {
		lua.Errorf(l, "step is immutable")
		return 0
	}
	stepName := lua.CheckString(l, 2)
	l.RawSet(1)
	prgm.steps = append(prgm.steps, stepName)
	return 0
}

func configureLua(prgm *LuaProgram, l *lua.State) {
	lua.NewMetaTable(l, "stepMetaTable")
	lua.SetFunctions(l, []lua.RegistryFunction{{
		"__newindex", prgm.registerStep,
	}}, 0)

	// create the `step` table, make it global, give it the
	// meta table that intercepts `__newindex` and then pop
	// them off the stack
	l.NewTable()
	l.PushValue(-1)
	l.SetGlobal("step")
	lua.SetMetaTableNamed(l, "stepMetaTable")
	l.Pop(2)
}