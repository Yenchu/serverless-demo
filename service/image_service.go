package service

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"serverless-demo/awsapi"
	"strconv"
	"strings"
)

const (
	MetadataWidth  = "Width"
	MetadataHeight = "Height"
)

func NewImageService() *ImageService {

	return &ImageService{
		s3Api: awsapi.NewS3API(),
	}
}

type ImageService struct {
	s3Api *awsapi.S3API
}

func (svc *ImageService) ResizeImage(bucket, key string) error {

	object, err := svc.s3Api.GetObject(bucket, key)
	if err != nil {
		return err
	}

	metadata := object.Metadata
	log.WithFields(log.Fields{"metadata": metadata}).Info("Get object")

	var contentType string
	var imgType string

	if object.ContentType != nil {
		contentType = *object.ContentType
		ctArr := strings.Split(contentType, "/")
		imgType = ctArr[len(ctArr)-1]
	}

	if imgType == "" {
		return errors.Errorf("bucket %s key %s doesn't have valid content type: %s", bucket, key, contentType)
	}

	newWidth := 0
	width := metadata[MetadataWidth]
	if width != "" {
		if n, err := strconv.Atoi(width); err == nil {
			newWidth = n
		}
	}

	newHeight := 0
	height := metadata[MetadataHeight]
	if height != "" {
		if n, err := strconv.Atoi(height); err == nil {
			newHeight = n
		}
	}

	if newWidth <= 0 || newHeight <= 0 {
		fmt.Printf("bucket %s key %s doesn't have valid resizing width %v and height %v", bucket, key, width, height)
		return nil
	}

	resizedObj, err := svc.Resize(object.Body, imgType, newWidth, newHeight)
	if err != nil {
		return err
	}

	// for test: set to zero to avoid resizing loop
	metadata[MetadataWidth] = "0"
	metadata[MetadataHeight] = "0"

	resp, err := svc.s3Api.PutObject(bytes.NewReader(resizedObj), bucket, key, contentType, metadata)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"result": resp.String()}).Info("Put resized object")
	return nil
}

func (svc *ImageService) Resize(content io.Reader, imgType string, width, height int) ([]byte, error) {

	img, err := imaging.Decode(content)
	if err != nil {
		return nil, err
	}

	resizedImg := imaging.Resize(img, width, height, imaging.Lanczos)

	format, err := imaging.FormatFromExtension(imgType)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	err = imaging.Encode(buf, resizedImg, format)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
