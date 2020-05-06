package awssvc

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Client() *S3Client {

	awsCfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	client := s3.New(awsCfg)

	return &S3Client{
		client: client,
	}
}

type S3PreSignURLRequest struct {
	Bucket      string
	Key         string
	ContentType string
	Metadata    map[string]string
	TTL         time.Duration
}

type S3Client struct {
	client *s3.Client
}

func (api *S3Client) GetPutObjectPreSignURL(reqData *S3PreSignURLRequest) (string, error) {

	input := &s3.PutObjectInput{
		Bucket:      aws.String(reqData.Bucket),
		Key:         aws.String(reqData.Key),
		ContentType: aws.String(reqData.ContentType),
		Metadata:    reqData.Metadata,
	}

	req := api.client.PutObjectRequest(input)

	return req.Presign(10 * time.Minute)
}

func (api *S3Client) GetPutObjectPreSignURLHeaders(reqData *S3PreSignURLRequest) (string, http.Header, error) {

	input := &s3.PutObjectInput{
		Bucket:      aws.String(reqData.Bucket),
		Key:         aws.String(reqData.Key),
		ContentType: aws.String(reqData.ContentType),
		Metadata:    reqData.Metadata,
	}

	req := api.client.PutObjectRequest(input)

	req.NotHoist = true
	url, headers, err := req.PresignRequest(reqData.TTL)
	return url, headers, err
}

func (api *S3Client) GetObject(bucket, key string) (*s3.GetObjectResponse, error) {

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	req := api.client.GetObjectRequest(input)

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *S3Client) PutObject(body io.Reader, bucket, key, contentType string, metadata map[string]string) (*s3.PutObjectResponse, error) {

	input := &s3.PutObjectInput{
		Body:        aws.ReadSeekCloser(body),
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
		Metadata:    metadata,
	}

	req := api.client.PutObjectRequest(input)

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
