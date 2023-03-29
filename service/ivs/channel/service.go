package ivs

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ivs"
)

type ServiceIVSChannel interface {
	CreateChannel(ctx context.Context) error
	GetChannel(ctx context.Context, arn string) (*ivs.GetChannelOutput, error)
	ListChannels(ctx context.Context) ([]*ivs.ChannelSummary, error)
	UpdateChannel(ctx context.Context) error
	DeleteChannels(ctx context.Context) error
	StatusService(ctx context.Context) error

	ListStreamKey(ctx context.Context, arnChannel string) ([]*ivs.StreamKeySummary, error)
	GetStreamKey(ctx context.Context, arnKey string) (*ivs.StreamKey, error)
}
