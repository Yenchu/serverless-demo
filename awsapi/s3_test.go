package awsapi_test

import (
	"serverless-demo/awsapi"
	"testing"
	"time"
)

func TestGetPutObjectPreSignURL(t *testing.T) {

	metadata := map[string]string{
		"width": "800",
		"height": "600",
	}

	reqData := &awsapi.S3PreSignURLRequest{
		Bucket: "abc",
		Key: "test.jpg",
		Metadata: metadata,
		TTL: 10 * time.Minute,
	}

	s3Api := awsapi.NewS3API()

	signedURL, err := s3Api.GetPutObjectPreSignURL(reqData)
	if err != nil {
		t.Fatalf("GetPutObjectPreSignURL failed: %v", err)
	}
	t.Logf("GetPutObjectPreSignURL: %s", signedURL)
}

func TestGetPutObjectPreSignURLHeaders(t *testing.T) {

	metadata := map[string]string{
		"width": "800",
		"height": "600",
	}

	reqData := &awsapi.S3PreSignURLRequest{
		Bucket: "abc",
		Key: "test.jpg",
		Metadata: metadata,
		TTL: 10 * time.Minute,
	}

	s3Api := awsapi.NewS3API()

	url, headers, err := s3Api.GetPutObjectPreSignURLHeaders(reqData)
	if err != nil {
		t.Fatalf("GetPutObjectPreSignURLHeaders failed: %v", err)
	}
	t.Logf("GetPutObjectPreSignURLHeaders:\nurl=%s\nheaders=%v", url, headers)
}