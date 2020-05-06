package main

import (
	"flag"
	"log"
	"myhttp/worker"
	"strconv"
)

const (
	parallelFlagName  = "parallel"
	maxParallelWorker = 10
)

func main() {

	numOfParallelWorkers := flag.Uint(parallelFlagName, maxParallelWorker, "no of jobs to be run in parallel")
	flag.Parse()

	sites := flag.Args()
	if len(sites) == 0 {
		log.Println("invalid no of args")
		return
	}

	flagPassed := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == parallelFlagName {
			flagPassed = true
		}
	})

	if !flagPassed { // fetch parallel worker from arg[0] otherwise leave it
		if parallel, err := strconv.Atoi(sites[0]); err == nil {
			*numOfParallelWorkers = (uint)(parallel)
			sites = sites[1:]
		}
	}

	err := worker.ProcessSites(sites, int(*numOfParallelWorkers))
	if err != nil {
		log.Printf("error occured while calling ProcessSites: %+v", err)
		return
	}
}
