package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/sign"
	"net/url"
	"serverless-demo/awssvc"
	"serverless-demo/model"
	"strings"
	"time"
)

const (
	SsmCFKeyID      = "/applications/ServerlessDemo/CloudFront/KeyId"
	SsmCFPrivateKey = "/applications/ServerlessDemo/CloudFront/PrivateKey"
	CFSignedURLTTL = 7 * 24 * time.Hour
)

func NewDownloadService() *DownloadService {

	ssmClient := awssvc.NewSsmClient()

	cfCfg, err := createCFConfig(ssmClient)
	if err != nil {
		panic("failed to create CloudFront config, " + err.Error())
	}

	cfClient := awssvc.NewCloudFrontClient(cfCfg)

	return &DownloadService{
		ssmClient: ssmClient,
		cfClient:  cfClient,
	}
}

func createCFConfig(ssmClient *awssvc.SsmClient) (*awssvc.CloudFrontConfig, error) {

	keyID, err := ssmClient.GetParameter(SsmCFKeyID)
	if err != nil {
		return nil, err
	}

	pkStr, err := ssmClient.GetDecryptedParameter(SsmCFPrivateKey)
	if err != nil {
		return nil, err
	}

	privKey, err := sign.LoadPEMPrivKey(strings.NewReader(pkStr))
	if err != nil {
		return nil, err
	}

	return &awssvc.CloudFrontConfig{
		KeyID:      keyID,
		PrivateKey: privKey,
	}, nil
}

type DownloadService struct {
	ssmClient *awssvc.SsmClient
	cfClient  *awssvc.CloudFrontClient
}

func (svc *DownloadService) GetDownloadURL(domain, file string) (*model.GetDownloadURLResponse, error) {

	urlObj := &url.URL{
		Scheme: "https",
		Host: domain,
		Path: file,
	}
	fmt.Printf("get download url: %s", urlObj.String())

	reqData := &awssvc.CFSignURLRequest{
		URL: urlObj.String(),
		TTL: CFSignedURLTTL,
	}

	url, err := svc.cfClient.GetSignURL(reqData)
	if err != nil {
		return nil, err
	}

	return &model.GetDownloadURLResponse{
		URL: url,
	}, nil
}
