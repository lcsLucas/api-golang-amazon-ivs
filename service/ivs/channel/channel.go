package ivs

import (
	"context"
	"errors"
	authaws "golang-ivs/auth/aws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ivs"
)

func switchErrors(err error) string {
	var msgError string

	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case ivs.ErrCodeResourceNotFoundException:
			msgError = "Recurso não encontrado"
		case ivs.ErrCodeAccessDeniedException:
			msgError = "Acesso negado"
		case ivs.ErrCodeConflictException:
			msgError = "Erro de conflito"
		case ivs.ErrCodeValidationException:
			msgError = "Erro de validação"
		default:
			msgError = aerr.Message()
		}
	}

	return msgError
}

type serviceChannel struct{}

func NewServiceChannel() ServiceIVSChannel {
	return &serviceChannel{}
}

func (serviceChannel) CreateChannel(ctx context.Context) error {
	return errors.New("not implemented")
}
func (serviceChannel) GetChannel(ctx context.Context, arnChannel string) (*ivs.GetChannelOutput, error) {
	sess, err := authaws.NewSession(nil)
	if err != nil {
		return nil, err
	}

	svc := ivs.New(sess, aws.NewConfig())

	channel, err := svc.GetChannel(&ivs.GetChannelInput{
		Arn: &arnChannel,
	})
	if err != nil {
		return nil, errors.New(switchErrors(err))
	}

	return channel, nil
}
func (serviceChannel) ListChannels(ctx context.Context) ([]*ivs.ChannelSummary, error) {

	sess, err := authaws.NewSession(nil)
	if err != nil {
		return nil, err
	}

	svc := ivs.New(sess, aws.NewConfig())

	channels, err := svc.ListChannels(&ivs.ListChannelsInput{})
	if err != nil {
		return nil, errors.New(switchErrors(err))
	}

	return channels.Channels, nil
}
func (serviceChannel) UpdateChannel(ctx context.Context) error {
	return errors.New("not implemented")
}
func (serviceChannel) DeleteChannels(ctx context.Context) error {
	return errors.New("not implemented")
}
func (serviceChannel) StatusService(ctx context.Context) error {
	return errors.New("not implemented")
}

func (serviceChannel) ListStreamKey(ctx context.Context, arnChannel string) ([]*ivs.StreamKeySummary, error) {

	sess, err := authaws.NewSession(nil)
	if err != nil {
		return nil, err
	}

	svc := ivs.New(sess, aws.NewConfig())

	keys, err := svc.ListStreamKeys(&ivs.ListStreamKeysInput{
		ChannelArn: &arnChannel,
	})
	if err != nil {
		return nil, errors.New(switchErrors(err))
	}

	return keys.StreamKeys, nil
}
func (serviceChannel) GetStreamKey(ctx context.Context, arnKey string) (*ivs.StreamKey, error) {

	sess, err := authaws.NewSession(nil)
	if err != nil {
		return nil, err
	}

	svc := ivs.New(sess, aws.NewConfig())

	key, err := svc.GetStreamKey(&ivs.GetStreamKeyInput{
		Arn: &arnKey,
	})
	if err != nil {
		return nil, errors.New(switchErrors(err))
	}

	return key.StreamKey, nil
}

func init() {

}
