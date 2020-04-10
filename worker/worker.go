package worker

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func ProcessSites(sites []string, numOfParallelWorkers int) error {
	var wg sync.WaitGroup

	numJobs := len(sites)
	jobs := make(chan string, numJobs)

	// start workers in parallel
	for w := 1; w <= numOfParallelWorkers; w++ {
		wg.Add(1)
		go startWorker(jobs, &wg)
	}

	// send jobs to workers
	for _, v := range sites {
		jobs <- v
	}
	close(jobs)

	// wait untill all the jobs are finished
	wg.Wait()
	return nil
}

func makeHttpRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var bodyBytes []byte
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("response code expected: %v, received: %v", http.StatusOK, resp.StatusCode)
	}

	return bodyBytes, nil
}

func startWorker(jobs <-chan string, wg *sync.WaitGroup) {
	for j := range jobs {
		url := fixSiteName(j)

		resp, err := makeHttpRequest(url)
		if err != nil {
			log.Printf("error occurred while calling makeHttpRequest for url %v: %v\n", url, err)
			continue // allow failing in http request as we can work on other inputs
		}

		printResult(url, resp)
	}

	wg.Done()
}

func printResult(url string, resp []byte) {
	fmt.Printf("%v %x\n", url, md5.Sum(resp))
}

func fixSiteName(site string) string {
	if !strings.HasPrefix(site, "http") {
		site = "http://" + site
	}
	return site
}
