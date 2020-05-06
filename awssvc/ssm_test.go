package awssvc_test

import (
	"serverless-demo/awssvc"
	"testing"
)

func TestGetDecryptedParameter(t *testing.T) {

	ssmClient := awssvc.NewSsmClient()

	val, err := ssmClient.GetDecryptedParameter("/applications/ServerlessDemo/CloudFront/PrivateKey")
	if err != nil {
		t.Fatalf("GetDecryptedParameter failed: %v", err)
	}
	t.Logf("GetDecryptedParameter: %v", val)

}

func TestGetParameters(t *testing.T) {

	ssmClient := awssvc.NewSsmClient()

	vals, err := ssmClient.GetParameters("/applications/ServerlessDemo/CloudFront/KeyId", "/applications/ServerlessDemo/CloudFront/PrivateKey")
	if err != nil {
		t.Fatalf("GetParameters failed: %v", err)
	}
	t.Logf("GetParameters: %v", vals)

}
