package main

import (
	"bitbucket.org/christian-m/resizeimage/internal/resize"
	"bitbucket.org/christian-m/resizeimage/internal/s3access"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"image"
	"log"
	"os"
	"strconv"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := make(map[string]string)
	headers["Content-Type"] = fmt.Sprint("text/plain")

	log.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)

	bucketName := os.Getenv("BUCKET_NAME")
	defaultFolder := os.Getenv("DEFAULT_FOLDER")

	path := request.PathParameters["proxy"]
	filePath := fmt.Sprintf("%s/%s", defaultFolder, path)
	log.Printf("file path: %s", filePath)
	s3file, err := s3access.FetchS3(bucketName, filePath)
	if err != nil {
		errorBody := fmt.Sprintf("%s\nfor file: %s", err.Error(), filePath)
		return events.APIGatewayProxyResponse{StatusCode: 404, Headers: headers, Body: errorBody, IsBase64Encoded: false}, nil
	}

	width, err := strconv.Atoi(request.QueryStringParameters["w"])
	if err != nil {
		width = 0
	}
	height, err := strconv.Atoi(request.QueryStringParameters["h"])
	if err != nil {
		height = 0
	}
	border, err := strconv.Atoi(request.QueryStringParameters["b"])
	if err != nil {
		border = 0
	}
	log.Printf("path=%s;w=%d;h=%d;b=%d", path, width, height, border)

	var byteSlice []byte
	b := bytes.NewBuffer(byteSlice)
	img, f, err := image.Decode(bytes.NewBuffer(s3file))
	if err != nil {
		errorBody := fmt.Sprintf("error decoding image: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 404, Headers: headers, Body: errorBody, IsBase64Encoded: false}, nil
	}
	picSize := resize.PicSize{Width: width, Height: height}
	picSize.EnsureImageBounds(img)
	picSize.AddBorder(border)
	err = resize.Resize(picSize, f, img, b)
	if err != nil {
		errorBody := fmt.Sprintf("%s\nfor file: %s", err.Error(), filePath)
		return events.APIGatewayProxyResponse{StatusCode: 404, Headers: headers, Body: errorBody, IsBase64Encoded: false}, nil
	}
	headers["Content-Type"] = fmt.Sprintf("image/%s", f)
	imageBody := base64.StdEncoding.EncodeToString(b.Bytes())
	return events.APIGatewayProxyResponse{StatusCode: 200, Headers: headers, Body: imageBody, IsBase64Encoded: true}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
