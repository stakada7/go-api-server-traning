package main

import (
	"testing"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Run("Index", func(t *testing T) {
		t.Parallel()

		s: httptest.NewServer(http.HandleFunc(index()))
		defer s.Close()

		res, err := http.Get(s.URL)
		assert.NoError(t, err)
		assert.Equal(t, "text/plain", res.Header.Get("Content-Type"))
		assert.Equal(t, 200, res.StatusCode)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, "", string(body))
	})

}
