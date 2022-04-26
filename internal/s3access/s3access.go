package s3access

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
)

func FetchS3(bucket, path string) ([]byte, error) {
	sess := session.Must(session.NewSession())
	downloader := s3manager.NewDownloader(sess)

	b := aws.NewWriteAtBuffer([]byte{})
	n, err := downloader.Download(b, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download file, %v", err)
	}
	log.Printf("file downloaded, %d bytes\n", n)
	return b.Bytes(), nil
}
