package aws

import (
	"fmt"

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

	result, _ := svc.ListBuckets(nil)

  // Info for bucket
	for _, bucket := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(bucket.Name), aws.TimeValue(bucket.CreationDate))
	}
	return svc, nil
}

func StringAws(item string) *string {
	return aws.String(item)
}
