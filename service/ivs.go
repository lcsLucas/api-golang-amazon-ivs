package service

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ivs"
	"github.com/joho/godotenv"
)

func NewStream(ctx context.Context) (interface{}, error) {
	ResponseStream := struct {
		Name string `json:"name"`
		URL  string `json:"url"`
		Key  string `json:"key"`
	}{}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("aws_region")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("aws_access_key_id"), os.Getenv("aws_secret_access_key"), ""),
	})

	if err != nil {
		return nil, err
	}

	// Create a IVS client with additional configuration
	svc := ivs.New(sess, aws.NewConfig())

	channels, err := svc.ListChannels(&ivs.ListChannelsInput{})
	if err != nil {
		return nil, err
	}

	fmt.Println(channels)
	fmt.Println(err)

	if channels != nil && len(channels.Channels) > 0 {
		arn := channels.Channels[0].Arn

		co, err := svc.GetChannel(&ivs.GetChannelInput{
			Arn: arn,
		})

		if err != nil {
			return nil, err
		}

		ResponseStream.Name = *co.Channel.Name
		ResponseStream.URL = fmt.Sprintf("rtmps://%s:443/app/", *co.Channel.IngestEndpoint)

		keys, err := svc.ListStreamKeys(&ivs.ListStreamKeysInput{
			ChannelArn: arn,
		})
		if err != nil {
			return nil, err
		}

		if len(keys.StreamKeys) > 0 {

			arnKey := keys.StreamKeys[0].Arn

			key, err := svc.GetStreamKey(&ivs.GetStreamKeyInput{
				Arn: arnKey,
			})
			if err != nil {
				return nil, err
			}

			ResponseStream.Key = *key.StreamKey.Value

		}

		return ResponseStream, nil

	}

	return []interface{}{}, nil
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
