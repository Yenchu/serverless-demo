package awssvc

import (
	"crypto/rsa"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/sign"
	"time"
)

func NewCloudFrontClient(cfg *CloudFrontConfig) *CloudFrontClient {

	awsCfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	client := cloudfront.New(awsCfg)

	return &CloudFrontClient{
		cfg:    cfg,
		client: client,
	}
}

type CFSignURLRequest struct {
	URL string
	TTL time.Duration
}

type CloudFrontConfig struct {
	KeyID      string
	PrivateKey *rsa.PrivateKey
}

type CloudFrontClient struct {
	cfg    *CloudFrontConfig
	client *cloudfront.Client
}

func (api *CloudFrontClient) GetSignURL(reqData *CFSignURLRequest) (string, error) {

	signer := sign.NewURLSigner(api.cfg.KeyID, api.cfg.PrivateKey)

	return signer.Sign(reqData.URL, time.Now().UTC().Add(reqData.TTL))
}
