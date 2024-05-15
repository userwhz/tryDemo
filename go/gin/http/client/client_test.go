package test

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	r, err := http.Post("http://localhost:8080/ping", "", nil)
	if err != nil {
		t.Error(err)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	t.Logf("%s", body)
	t.Error("test")
}
