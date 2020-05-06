package service

import (
	"serverless-demo/awssvc"
	"serverless-demo/model"
	"strings"
	"time"
)

const S3PreSignedURLTTL = 10 * time.Minute

func NewUploadService() *UploadService {

	return &UploadService{
		s3Client: awssvc.NewS3Client(),
	}
}

type UploadService struct {
	s3Client *awssvc.S3Client
}

func (svc *UploadService) GetUploadURL(bucket, key, contentType string, metadata map[string]string) (*model.GetUploadURLResponse, error) {

	reqData := &awssvc.S3PreSignURLRequest{
		Bucket:      bucket,
		Key:         key,
		ContentType: contentType,
		Metadata:    metadata,
		TTL:         S3PreSignedURLTTL,
	}

	url, headers, err := svc.s3Client.GetPutObjectPreSignURLHeaders(reqData)
	if err != nil {
		return nil, err
	}

	rtHeaders := map[string]string{}
	for name, values := range headers {
		value := strings.Join(values, ",")
		rtHeaders[name] = value
	}

	return &model.GetUploadURLResponse{
		URL:     url,
		Headers: rtHeaders,
	}, nil
}
