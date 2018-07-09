package gogitlab

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func Stub(filename string) (*httptest.Server, *Gitlab) {
	var err error

	stub := []byte("")

	if len(filename) > 0 {
		stub, err = ioutil.ReadFile(filename)

		if err != nil {
			panic(err)
		}
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	gitlab := NewGitlab(ts.URL, "", "")

	return ts, gitlab
}
