package worker

import (
	"log"
	"myhttp/testhelper"
	"sync"
	"testing"
)

func TestMakeHttpRequest(t *testing.T) {
	ts, err := testhelper.StartTestServer("127.0.0.1:32001", "hello")
	if err != nil {
		log.Fatal(err)
	}
	defer ts.Close()

	_, err = makeHttpRequest(ts.URL)
	if err != nil {
		t.Errorf("failed to make request from makeHttpRequest to url: %v", ts.URL)
	}
}

func TestMakeHttpRequestFailure(t *testing.T) {
	_, err := makeHttpRequest("http://www.ganeshdip.dumbare.com")
	if err == nil {
		t.Error("failed to make request from makeHttpRequest to url: http://www.ganeshdip.dumbare.com")
	}
}

func TestFixSiteName(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"google.com", "http://google.com"},
		{"www.fb.com", "http://www.fb.com"},
		{"http://www.google.com", "http://www.google.com"},
	}
	for _, test := range tests {
		out := fixSiteName(test.input)
		if out != test.output {
			t.Errorf("Failed to fixSiteName, expected: <%v>,got:<%v>", test.output, out)
		}
	}
}

func Example_printResult() {
	printResult("hello.com", []byte("hello"))
	// Output:
	// hello.com 5d41402abc4b2a76b9719d911017c592
}

func Example_startWorker() {
	var wg sync.WaitGroup
	j := make(chan string, 1)

	wg.Add(1)
	go startWorker(j, &wg)

	ts, err := testhelper.StartTestServer("127.0.0.1:32001", "hello")
	if err != nil {
		log.Fatal(err)
	}
	defer ts.Close()

	j <- ts.URL
	close(j)

	wg.Wait()
	// Output:
	// http://127.0.0.1:32001 b1946ac92492d2347c6235b4d2611184
}

func Example_startWorkerFailure() { // This will output nothing
	var wg sync.WaitGroup
	j := make(chan string, 1)

	wg.Add(1)
	go startWorker(j, &wg)

	j <- "http://www.ganeshdip.dumbare.com"
	close(j)

	wg.Wait()
	// Output:
	//
}

func ExampleProcessSites() {
	sites := []string{}
	ts, err := testhelper.StartTestServer("127.0.0.1:32001", "hello")
	if err != nil {
		log.Fatal(err)
	}
	defer ts.Close()

	sites = append(sites, ts.URL)
	err = ProcessSites(sites, 1)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// http://127.0.0.1:32001 b1946ac92492d2347c6235b4d2611184
}

// send 2 sites and start 1 worker
func ExampleProcessSites_multiple() {
	sites := []string{}
	ts1, err := testhelper.StartTestServer("127.0.0.1:32001", "hello")
	if err != nil {
		log.Fatal(err)
	}
	defer ts1.Close()

	ts2, err := testhelper.StartTestServer("127.0.0.1:32002", "world")
	if err != nil {
		log.Fatal(err)
	}
	defer ts2.Close()

	sites = append(sites, ts1.URL, ts2.URL)
	err = ProcessSites(sites, 1)
	if err != nil {
		log.Fatal(err)
	}
	// Unordered output:
	// http://127.0.0.1:32001 b1946ac92492d2347c6235b4d2611184
	// http://127.0.0.1:32002 591785b794601e212b260e25925636fd
}

// send 2 sites and start 2 workers in parallel
func ExampleProcessSites_multipleParallel() {
	sites := []string{}
	ts1, err := testhelper.StartTestServer("127.0.0.1:32001", "hello")
	if err != nil {
		log.Fatal(err)
	}
	defer ts1.Close()

	ts2, err := testhelper.StartTestServer("127.0.0.1:32002", "world")
	if err != nil {
		log.Fatal(err)
	}
	defer ts2.Close()

	sites = append(sites, ts1.URL, ts2.URL)
	err = ProcessSites(sites, 2)
	if err != nil {
		log.Fatal(err)
	}
	// Unordered output:
	// http://127.0.0.1:32001 b1946ac92492d2347c6235b4d2611184
	// http://127.0.0.1:32002 591785b794601e212b260e25925636fd
}
