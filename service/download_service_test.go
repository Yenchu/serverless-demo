package service_test

import (
	"os"
	"serverless-demo/model"
	"serverless-demo/service"
	"testing"
)

func TestGetDownloadURL(t *testing.T) {

	req := &model.GetDownloadURLRequest{
		Scheme: "https",
		Domain: os.Getenv("CF_DOMAIN_NAME"),
		File:   "test.jpg",
	}

	downloadSvc, err := service.NewDownloadService()
	if err != nil {
		t.Fatalf("NewDownloadService failed: %v", err)
	}

	resp, err := downloadSvc.GetDownloadURL(req)
	if err != nil {
		t.Fatalf("GetDownloadURL failed: %v", err)
	}
	t.Logf("GetDownloadURL: %v", resp)
}
