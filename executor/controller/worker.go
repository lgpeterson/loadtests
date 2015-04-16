package controller

import (
	"log"
	"strings"
	"sync"

	"github.com/lgpeterson/loadtests/executor/engine"
	"github.com/lgpeterson/loadtests/executor/executorGRPC"
	"golang.org/x/net/context"
)

type worker struct {
	Persist    Persister
	Command    *executorGRPC.CommandMessage
	Wait       *sync.WaitGroup
	JobChannel <-chan struct{}
	Done       <-chan struct{}
}

func (w *worker) execute() {
	w.Wait.Add(1)
	defer w.Wait.Done()
	for {
		select {
		case <-w.Done:
			return
		case <-w.JobChannel:
			scriptReader := strings.NewReader(w.Command.Script)
			prog, err := engine.Lua(scriptReader)
			if err != nil {
				// This should not be because the script did not compile, if it
				// did not compile it would be reported to the user before this
				log.Printf("Error creating lua script: %v", err)
				return
			}
			err = prog.Execute(context.Background())

			if err != nil {
				// I assume I can keep going if the lua script encoutered an error
				log.Printf("Error running lua script: %v", err)
				// TODO: Log error in influx?
				continue
			}

			w.Persist.Persist(w.Command.ScriptName, w.Command.URL, []byte{})
		}

	}
}
