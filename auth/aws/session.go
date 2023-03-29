package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewSession(config *aws.Config) (*session.Session, error) {

	if config == nil {
		config = &aws.Config{
			Region:      aws.String(os.Getenv("aws_region")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("aws_access_key_id"), os.Getenv("aws_secret_access_key"), ""),
		}
	}

	sess, err := session.NewSession(config)

	return sess, err
}
