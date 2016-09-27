package handler

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/drewolson/testflight"
	"github.com/stretchr/testify/assert"
)

func TestMainHandler(t *testing.T) {
	testflight.WithServer(NewHandler(), func(r *testflight.Requester) {
		response := r.Get("/")

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, "Team Guru app.", response.Body)
	})
}

func TestHealthcheckHandler(t *testing.T) {
	testflight.WithServer(NewHandler(), func(r *testflight.Requester) {
		response := r.Get("/healtz")

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, "Yeah, I'm ready.", response.Body)
	})
}

func TestVersionHandler(t *testing.T) {
	testflight.WithServer(NewHandler(), func(r *testflight.Requester) {
		response := r.Get("/version")

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Regexp(t, regexp.MustCompile("[0-9]+\\.[0-9]+\\.[0-9]+"), response.Body)
	})
}
