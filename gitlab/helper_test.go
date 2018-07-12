package gitlab

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

const mocksDir = "mocks"
const mocksMappingsDir = "mappings"
const mocksFilesDir = "__files"

type mockMapping struct {
	Request struct {
		Url    string `json:"url"`
		Method string `json:"method"`
	} `json:"request"`
	Response struct {
		Status       int               `json:"status"`
		BodyFileName string            `json:"bodyFileName"`
		Headers      map[string]string `json:"headers"`
	} `json:"response"`
}

type responseMock struct {
	mapping *mockMapping
	body    []byte
}

func loadResponseMock(t *testing.T, mappingFile string) *responseMock {
	t.Helper()

	_, currentFilename, _, ok := runtime.Caller(0)
	if !ok {
		t.Errorf("unable to determine caller")
		t.FailNow()
		return nil
	}

	baseDir, err := filepath.Abs(filepath.Join(filepath.Dir(currentFilename), ".."))
	if err != nil {
		t.Errorf("unable to determine base dir from %s\n%v", currentFilename, err)
		t.FailNow()
		return nil
	}

	mappingFilePath := filepath.Join(baseDir, mocksDir, mocksMappingsDir, mappingFile)
	mappingRaw, err := ioutil.ReadFile(mappingFilePath)
	if err != nil {
		t.Errorf("unable to load mock mapping file: %s\n%v", mappingFilePath, err)
		t.FailNow()
		return nil
	}

	mapping := new(mockMapping)
	err = json.Unmarshal(mappingRaw, &mapping)
	if err != nil {
		t.Errorf("unable to unmarshal mock mapping file: %s\n%v\n\ncontent:\n%s", mappingFilePath, err, mappingRaw)
		t.FailNow()
		return nil
	}

	bodyFilePath := filepath.Join(baseDir, mocksDir, mocksFilesDir, mapping.Response.BodyFileName)
	body, err := ioutil.ReadFile(bodyFilePath)
	if err != nil {
		t.Errorf("unable to load mock body file: %s\n%v", bodyFilePath, err)
		t.FailNow()
		return nil
	}

	return &responseMock{
		mapping: mapping,
		body:    body,
	}
}

func mockServerFromMapping(t *testing.T, mappingFile string) (*httptest.Server, *Gitlab) {
	mock := loadResponseMock(t, mappingFile)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range mock.mapping.Response.Headers {
			w.Header().Add(k, v)
		}

		w.WriteHeader(mock.mapping.Response.Status)
		w.Write([]byte(mock.body))
	})

	ts := httptest.NewServer(handler)
	gitlab := NewGitlab(ts.URL, "", "")

	return ts, gitlab
}
