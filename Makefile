.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./fn/get-upload-url/get-upload-url
	rm -rf ./fn/get-download-url/get-download-url
	rm -rf ./fn/resize-image/resize-image
	rm -rf ./fn/signin/signin
	rm -rf ./fn/respond-to-challenge/respond-to-challengen
	
build:
	GOOS=linux GOARCH=amd64 go build -o fn/get-upload-url/get-upload-url ./fn/get-upload-url
	GOOS=linux GOARCH=amd64 go build -o fn/get-download-url/get-download-url ./fn/get-download-url
	GOOS=linux GOARCH=amd64 go build -o fn/resize-image/resize-image ./fn/resize-image
	GOOS=linux GOARCH=amd64 go build -o fn/signin/signin ./fn/signin
	GOOS=linux GOARCH=amd64 go build -o fn/respond-to-challenge/respond-to-challenge ./fn/respond-to-challenge