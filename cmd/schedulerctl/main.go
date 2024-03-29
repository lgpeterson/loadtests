package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/codegangsta/cli"
	"github.com/dustin/go-humanize"
	"github.com/flynn/flynn/controller/name"
	"github.com/lgpeterson/loadtests/scheduler/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	appname = "schedulerctl"
	version = "devel"
	usage   = `
interact with the scheduler through the command line
`
)

func init() {
	name.SetSeed([]byte(time.Now().Format(time.UnixDate))[:10])
}

var (
	addrFlag = cli.StringFlag{Name: "addr", Usage: "address where the scheduler service can be reached"}
	tgtFlag  = cli.StringFlag{Name: "tgt", Usage: "target URL to execute the load test against"}

	scriptNameFlag    = cli.StringFlag{Name: "script.name", Usage: "name of the script"}
	scriptFileFlag    = cli.StringFlag{Name: "script.file", Usage: "if specified, the file where the source of the script can be found. Otherwise uses stdin"}
	scriptConfigFlag  = cli.StringFlag{Name: "script.config", Usage: "if specified, the file where the source of the config can be found.", Value: ""}
	runTimeFlag       = cli.DurationFlag{Name: "duration", Value: time.Minute, Usage: "how long to perform the load test for"}
	maxExecPerSecFlag = cli.IntFlag{Name: "max.exec.ps", Value: 100, Usage: "number of executions per second"}

	growthFactorFlag              = cli.Float64Flag{Name: "extra.growth.factor", Value: 1.5}
	timeBetweenGrowthFlag         = cli.DurationFlag{Name: "extra.time.between.growth", Value: time.Second}
	startingRequestsPerSecondFlag = cli.IntFlag{Name: "extra.starting.requests.ps", Value: 50}
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(appname + ": ")
	if err := newApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = appname
	app.Usage = strings.TrimSpace(usage)
	app.Version = version
	app.Author = "Antoine Grondin"
	app.Email = "antoinegrondin@gmail.com"

	var client pb.SchedulerClient
	app.Flags = []cli.Flag{
		addrFlag,
		tgtFlag,
		scriptNameFlag,
		scriptFileFlag,
		scriptConfigFlag,
		runTimeFlag,
		maxExecPerSecFlag,
		growthFactorFlag,
		timeBetweenGrowthFlag,
		startingRequestsPerSecondFlag,
	}
	app.Before = func(ctx *cli.Context) error {
		addr := ctx.GlobalString(addrFlag.Name)
		cc, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithTimeout(2*time.Second), grpc.WithBlock())
		if err != nil {
			return err
		}
		client = pb.NewSchedulerClient(cc)
		return nil
	}

	app.Action = func(ctx *cli.Context) {
		if ctx.GlobalString(scriptNameFlag.Name) == "" {
			log.Fatalf("param %q required to run script", scriptNameFlag.Name)
		}
		script, err := readFileOrStdin(ctx, scriptFileFlag)
		if err != nil {
			log.Fatal(err)
		}
		scriptConfig, err := readFileIfExists(ctx, scriptConfigFlag)
		if err != nil {
			log.Fatal(err)
		}

		in := &pb.LoadTestReq{
			Url:                       ctx.GlobalString(tgtFlag.Name),
			ScriptName:                ctx.GlobalString(scriptNameFlag.Name),
			Script:                    string(script),
			MaxRequestsPerSecond:      int32(ctx.GlobalInt(maxExecPerSecFlag.Name)),
			RunTime:                   int32(ctx.GlobalDuration(runTimeFlag.Name).Seconds()),
			GrowthFactor:              ctx.Float64(growthFactorFlag.Name),
			TimeBetweenGrowth:         ctx.Duration(timeBetweenGrowthFlag.Name).Seconds(),
			StartingRequestsPerSecond: int32(ctx.GlobalInt(maxExecPerSecFlag.Name)),
			ScriptConfig:              string(scriptConfig),
		}
		if in.StartingRequestsPerSecond == 0 {
			in.StartingRequestsPerSecond = in.MaxRequestsPerSecond
		}
		if in.GrowthFactor == 0 {
			in.GrowthFactor = 1
		}
		if in.TimeBetweenGrowth == 0 {
			in.TimeBetweenGrowth = 1
		}

		log.Printf("requesting %v load test at %drps on %q with script %q (%v)",
			time.Duration(in.RunTime)*time.Second,
			in.MaxRequestsPerSecond,
			in.Url,
			in.ScriptName,
			humanize.IBytes(uint64(len(in.Script))),
		)
		now := time.Now()
		srv, err := client.LoadTest(context.Background(), in)
		if err != nil {
			log.Fatalf("issuing load test request: %v", err)
		}
		for {
			res, err := srv.Recv()
			switch err {
			case io.EOF:
				log.Print("done")
				return
			default:
				log.Fatalf("waiting for response: %v", err)
			case nil:
			}
			switch {
			case res.GetPreparing() != nil:
				log.Printf("%s: load test is preparing %d workers...", time.Since(now), res.GetPreparing().Count)
			case res.GetStart() != nil:
				log.Printf("%s: load test started!", time.Since(now))
			case res.GetFinish() != nil:
				log.Printf("%s: load test finished!", time.Since(now))
			case res.GetError() != nil:
				log.Printf("%s: load test had an error: %v", time.Since(now), res.GetError().Error)
			default:
				log.Printf("%s: unexpected message: %#v", time.Since(now), res)
			}
		}
	}

	return app
}

func readFileOrStdin(ctx *cli.Context, fileFlag cli.StringFlag) ([]byte, error) {
	filename := ctx.GlobalString(fileFlag.Name)
	if filename == "" || filename == "-" {
		waiting := time.AfterFunc(time.Second, func() {
			log.Printf("reading script source from stdin, waiting for EOF...")
		})
		defer waiting.Stop()
		return ioutil.ReadAll(os.Stdin)
	}
	return readFile(filename)
}

func readFileIfExists(ctx *cli.Context, fileFlag cli.StringFlag) ([]byte, error) {
	filename := ctx.GlobalString(fileFlag.Name)
	if filename != "" {
		return readFile(filename)
	}
	return make([]byte, 0), nil
}

func readFile(filename string) ([]byte, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	return ioutil.ReadAll(fd)
}
