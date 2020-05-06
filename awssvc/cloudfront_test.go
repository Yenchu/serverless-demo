package awssvc_test

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/sign"
	"io/ioutil"
	"os"
	"serverless-demo/awssvc"
	"strings"
	"testing"
)

func TestLoadFilePrivKey(t *testing.T) {

	content, err := ioutil.ReadFile(os.Getenv("PRIVATE_KEY_FILE"))
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	pkStr := string(content)

	privKey, err := sign.LoadPEMPrivKey(strings.NewReader(pkStr))
	if err != nil {
		t.Fatalf("LoadPrivKey failed: %v", err)
	}
	t.Logf("LoadPrivKey: %v", privKey)
}

func TestLoadSSMPrivKey(t *testing.T) {

	ssmClient := awssvc.NewSsmClient()

	pkStr, err := ssmClient.GetDecryptedParameter("/applications/ServerlessDemo/CloudFront/PrivateKey")
	if err != nil {
		t.Fatalf("GetDecryptedParameter failed: %v", err)
	}

	privKey, err := sign.LoadPEMPrivKey(strings.NewReader(pkStr))
	if err != nil {
		t.Fatalf("LoadPrivKey failed: %v", err)
	}
	t.Logf("LoadPrivKey: %v", privKey)
}