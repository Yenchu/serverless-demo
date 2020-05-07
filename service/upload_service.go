package service

import (
	"serverless-demo/awsapi"
	"serverless-demo/model"
	"strconv"
	"strings"
	"time"
)

const S3PreSignedUrlTTL = 10 * time.Minute

func NewUploadService() *UploadService {

	return &UploadService{
		s3Api: awsapi.NewS3API(),
	}
}

type UploadService struct {
	s3Api *awsapi.S3API
}

func (svc *UploadService) GetUploadURL(req *model.GetUploadURLRequest) (*model.GetUploadURLResponse, error) {

	metadata := map[string]string{}

	width := req.Width
	if width > 0 {
		metadata["width"] = strconv.FormatInt(width, 10)
	}

	height := req.Height
	if height > 0 {
		metadata["height"] = strconv.FormatInt(height, 10)
	}

	reqData := &awsapi.S3PreSignURLRequest{
		Bucket:      req.Bucket,
		Key:         req.File,
		ContentType: req.ContentType,
		Metadata:    metadata,
		TTL:         S3PreSignedUrlTTL,
	}

	url, headers, err := svc.s3Api.GetPutObjectPreSignURLHeaders(reqData)
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
