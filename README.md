## Serverless image uploading and downloading

This is a demo application for image uploading and downloading.
It uses S3 prs-signed URL to upload to S3, and CloudFront signed URL to download from CloudFront.

### Architecture Design

![Architecture Diagram](https://github.com/yenchu/serverless-demo/raw/master/images/architecture.png)

### Build and Deployment

This project uses SAM to build and deploy, please follow the AWS document
 [Installing the AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
 to install.

### Configuration

Before deployment, please configure the following items.

##### Enable API Gateway Logging

If your API Gateway doesn't have permissions to write CloudWatch logs, please enable it on API Gateway console. 

You can also enable it by deploying the CloudFormation file `cloudformation/apigateway-global-settings.yaml`.
 
##### Generate CloudFront Key Pair

Because signing CloudFront URL needs a key pair, you need to use root account to create one.

On AWS console, select `My Security Credentials` in menu bar, select `CloudFront key pairs`, click `Create New Key Pair`.
Then you need to store the generate key ID and private key in SSM parameter store. 

##### Store Key Pair in SSM Parameter Store

To store CloudFront key ID in SSM parameter store, go to `Systems Manager` console, select `Parameter Store`,
 save it with name `/applications/ServerlessDemo/CloudFront/KeyId` and type `String`.
 
To store private key in parameter store, save it with name `/applications/ServerlessDemo/CloudFront/PrivateKey` and type `SecureString`.

### Test

You can use cURL to test these APIs.

##### Get S3 PreSigned URL for Upload

To get a S3 pre-signed URL for uploading, you need to provide file name and content type.
If you want to resize the image, you can provide width and height you want.

```
curl -X POST -H "Content-Type: application/json" https://{API_GATEWAY_ENDPOINT}/get-upload-url
 -d '{"file": "{FILE_NAME}", "contentType": "image/jpg", "width": 2048, "height": 1024}'
```

You will get the following response, the `url` is a S3 pre-signed URL you can use to upload file to S3. 
Please note the response headers need to be passed back to S3 when uploading file: 

```json
{
  "headers": {
    "content-type": "image/jpg",
    "x-amz-meta-height": "1024",
    "x-amz-meta-width": "2048"
  },
  "url": "{UPLOAD_URL}"
}
```

##### Upload to S3

To upload file to S3 using pre-signed URL, you need to pass back the headers you got from calling get-upload-url API.

``` 
curl -X PUT -H "content-type: image/jpg" -H "x-amz-meta-heigh: 1024" -H "x-amz-meta-width: 2048" {UPLOAD_URL}
 --data-binary '@{PATH_TO_IMAGE}.jpg'
```

##### Get CloudFront Signed URL for Download

To get a CloudFront signed URL for downloading, you need to provide file name you want to download.

```
curl -X POST -H "Content-Type: application/json" https://{API_GATEWAY_ENDPOINT}/get-download-url
 -d '{"file": "{FILE_NAME}"'
```

You will get the following response, the `url` is a CloudFront signed URL you can use to download file from CloudFront. 

```json
{
  "url": "{DOWNLOAD_URL}"
}
```

##### Download from CloudFront

```
curl -X GET {DOWNLOAD_URL}
```
