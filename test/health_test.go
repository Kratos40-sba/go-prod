//go:build e2e
// +build e2e

package test

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	log.Println("Running E2E test for health check endpoint")
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/health")
	if err != nil {
		t.Fail()
	}

	//log.Println(resp.StatusCode())
	assert.Equal(t, 200, resp.StatusCode())
	// cmd : go test ./... -tags=e2e -v
}
