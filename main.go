// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import (
	"flag"
	"log"
	"myhttp/arguments"
	"myhttp/worker"
	"os"
)

const (
	parallelFlagName  = "parallel"
	maxParallelWorker = 10
)

func main() {

	numOfParallelWorkers := flag.Uint(parallelFlagName, maxParallelWorker, "no of jobs to be run in parallel")
	flag.Parse()

	args := os.Args[1:]
	sites, err := arguments.GetSitesFromArgs(args, parallelFlagName)
	if err != nil {
		log.Printf("error occured while calling GetSitesFromArgs: %+v", err)
		return
	}

	err = worker.ProcessSites(sites, int(*numOfParallelWorkers))
	if err != nil {
		log.Printf("error occured while calling ProcessSites: %+v", err)
		return
	}
}
