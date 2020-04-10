package testhelper

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
)

func StartTestServer(url, respMsg string) (*httptest.Server, error) {
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respMsg)
	}))

	l, err := net.Listen("tcp", url)
	if err != nil {
		return nil, err
	}
	ts.Listener = l
	ts.Start()

	return ts, nil
}
