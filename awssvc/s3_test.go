package awssvc_test

import (
	"serverless-demo/awssvc"
	"testing"
	"time"
)

func TestGetPutObjectPreSignURL(t *testing.T) {

	metadata := map[string]string{
		"width": "800",
		"height": "600",
	}

	reqData := &awssvc.S3PreSignURLRequest{
		Bucket: "abc",
		Key: "test.jpg",
		Metadata: metadata,
		TTL: 10 * time.Minute,
	}

	s3Client := awssvc.NewS3Client()

	signedURL, err := s3Client.GetPutObjectPreSignURL(reqData)
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

	reqData := &awssvc.S3PreSignURLRequest{
		Bucket: "abc",
		Key: "test.jpg",
		Metadata: metadata,
		TTL: 10 * time.Minute,
	}

	s3Client := awssvc.NewS3Client()

	url, headers, err := s3Client.GetPutObjectPreSignURLHeaders(reqData)
	if err != nil {
		t.Fatalf("GetPutObjectPreSignURLHeaders failed: %v", err)
	}
	t.Logf("GetPutObjectPreSignURLHeaders:\nurl=%s\nheaders=%v", url, headers)
}