package awssvc

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"time"
)

func NewSsmClient() *SsmClient {

	awsCfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	client := ssm.New(awsCfg)

	return &SsmClient{
		client: client,
	}
}

type SsmClient struct {
	client *ssm.Client
}

func (api *SsmClient) GetParameter(name string) (string, error) {

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	input := &ssm.GetParameterInput{
		Name: &name,
	}

	req := api.client.GetParameterRequest(input)

	resp, err := req.Send(ctx)
	if err != nil {
		return "", err
	}
	return *resp.Parameter.Value, nil
}

func (api *SsmClient) GetDecryptedParameter(name string) (string, error) {

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	decrypt := true
	input := &ssm.GetParameterInput{
		Name:           &name,
		WithDecryption: &decrypt,
	}

	req := api.client.GetParameterRequest(input)

	resp, err := req.Send(ctx)
	if err != nil {
		return "", err
	}
	return *resp.Parameter.Value, nil
}

func (api *SsmClient) GetParameters(names ...string) (map[string]string, error) {

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	input := &ssm.GetParametersInput{
		Names: names,
	}

	req := api.client.GetParametersRequest(input)

	resp, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}

	result := map[string]string{}
	for _, param := range resp.Parameters {
		result[*param.Name] = *param.Value
	}
	return result, nil
}
