package service_test

import (
	"os"
	"serverless-demo/service"
	"testing"
)

func TestGetDownloadURL(t *testing.T) {

	domain := os.Getenv("CF_DOMAIN_NAME")

	file := "test.jpg"

	downloadSvc := service.NewDownloadService()

	resp, err := downloadSvc.GetDownloadURL(domain, file)
	if err != nil {
		t.Fatalf("GetDownloadURL failed: %v", err)
	}
	t.Logf("GetDownloadURL: %v", resp)
}
