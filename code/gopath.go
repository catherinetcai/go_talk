package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Creates a new session from the AWS Go SDK to interface with S3
	sess := session.Must(session.NewSession())

	svc := s3.New(sess)
}
