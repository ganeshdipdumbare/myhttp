package main

import (
	"log"
	"myhttp/testhelper"
	"os"
)

func Example_main_with_args() {

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	ts1, err := testhelper.StartTestServer("127.0.0.1:32004", "hello")
	if err != nil {
		log.Fatal(err)
	}
	defer ts1.Close()

	ts2, err := testhelper.StartTestServer("127.0.0.1:32005", "world")
	if err != nil {
		log.Fatal(err)
	}
	defer ts2.Close()

	ts3, err := testhelper.StartTestServer("127.0.0.1:32006", "welcome")
	if err != nil {
		log.Fatal(err)
	}
	defer ts3.Close()

	os.Args = []string{oldArgs[0], ts1.URL, ts2.URL, ts3.URL}
	main()
	// Unordered output:
	// http://127.0.0.1:32004 b1946ac92492d2347c6235b4d2611184
	// http://127.0.0.1:32006 0bb3c30dc72e63881db5005f1aa19ac3
	// http://127.0.0.1:32005 591785b794601e212b260e25925636fd
}
