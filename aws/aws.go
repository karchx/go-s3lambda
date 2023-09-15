package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateS3Session() (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: StringAws("us-east-1"),
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	return svc, nil
}

func StringAws(item string) *string {
	return aws.String(item)
}
