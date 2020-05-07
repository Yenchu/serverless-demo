.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./fn/get-upload-url/get-upload-url
	rm -rf ./fn/get-download-url/get-download-url
	rm -rf ./fn/resize-image/resize-image
	
build:
	GOOS=linux GOARCH=amd64 go build -o fn/get-upload-url/get-upload-url ./fn/get-upload-url
	GOOS=linux GOARCH=amd64 go build -o fn/get-download-url/get-download-url ./fn/get-download-url
	GOOS=linux GOARCH=amd64 go build -o fn/resize-image/resize-image ./fn/resize-image