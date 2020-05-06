package awssvc_test

import (
	"os"
	"serverless-demo/awssvc"
	"testing"
)

func newKmsClient() *awssvc.KmsClient {

	keyID := os.Getenv("KMS_KEY_ID")

	cfg := &awssvc.KmsConfig{
		KeyID: keyID,
	}

	return awssvc.NewKmsClient(cfg)
}

func TestListKeys(t *testing.T)  {

	kmsClient := newKmsClient()

	resp, err := kmsClient.ListKeys()
	if err != nil {
		t.Fatalf("ListKeys failed: %v", err)
	}
	t.Logf("ListKeys: %v", resp)
}

func TestEncrypt(t *testing.T)  {

	content := []byte("plain data")

	kmsClient := newKmsClient()

	resp, err := kmsClient.Encrypt(content)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}
	t.Logf("Encrypt: %v", resp)
}