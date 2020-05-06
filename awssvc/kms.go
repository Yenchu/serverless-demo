package awssvc

import (
	"context"
	"encoding/base64"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func NewKmsClient(cfg *KmsConfig) *KmsClient {

	awsCfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	client := kms.New(awsCfg)

	return &KmsClient{
		cfg:    cfg,
		client: client,
	}
}

type KmsConfig struct {
	KeyID string
}

type KmsClient struct {
	cfg    *KmsConfig
	client *kms.Client
}

func (api *KmsClient) ListKeys() (*kms.ListKeysResponse, error) {

	input := &kms.ListKeysInput{}

	req := api.client.ListKeysRequest(input)

	return req.Send(context.Background())
}

func (api *KmsClient) Encrypt(plainText []byte) (string, error) {

	input := &kms.EncryptInput{
		KeyId:     aws.String(api.cfg.KeyID),
		Plaintext: plainText,
	}

	req := api.client.EncryptRequest(input)

	resp, err := req.Send(context.Background())
	if err != nil {
		return "", err
	}

	encodedText := base64.StdEncoding.EncodeToString(resp.CiphertextBlob)
	return encodedText, nil
}
