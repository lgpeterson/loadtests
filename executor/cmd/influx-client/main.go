package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/lgpeterson/loadtests/executor/controller"
	"github.com/lgpeterson/loadtests/executor/persister"
)

func main() {
	log.SetFlags(0)

	addr := flag.String("addr", "localhost:50045", "the IP and port to coneect to")
	flag.Parse()

	testIflux(*addr)
}

func testIflux(ip string) {
	metrics := controller.NewMetricsGatherer()
	metrics.IncrHTTPGet("http://localhost/foo", 99, time.Millisecond)

	pass := os.Getenv("INFLUX_PWD")
	user := os.Getenv("INFLUX_USER")

	persister := &persister.InfluxPersister{}
	err := persister.SetupPersister(ip, user, pass, "ltm_metrics", true)
	if err != nil {
		log.Fatalf("Error creating influx persistor: %v", err)
	}
	err = persister.Persist("test_run", metrics)
	if err != nil {
		log.Fatalf("Error with influx persistor: %v", err)
	}
	count, err := persister.CountOccurrences("test_run", "GetRequestTable")
	if err != nil {
		log.Fatalf("Error with influx persistor getting count: %v", err)
	}
	log.Printf("Count is: %d", count)

}