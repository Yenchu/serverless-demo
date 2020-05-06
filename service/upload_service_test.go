package service_test

import (
	"serverless-demo/service"
	"testing"
)

func TestGetUploadURL(t *testing.T) {

	bucket := "abc"
	key := "test.jpg"
	contentType := "image/jpg"

	metadata := map[string]string{
		"width": "800",
		"height": "600",
	}

	uploadSvc := service.NewUploadService()

	resp, err := uploadSvc.GetUploadURL(bucket, key, contentType, metadata)
	if err != nil {
		t.Fatalf("GetUploadURL failed: %v", err)
	}
	t.Logf("GetUploadURL: %v", resp)
}
