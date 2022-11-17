package secrets

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type SecretsManager struct {
	client *ssm.SSM
}

func New() *SecretsManager {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("us-east-1")},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}
	client := ssm.New(sess, aws.NewConfig().WithRegion("us-east-1"))

	return &SecretsManager{
		client: client,
	}
}

func (sm *SecretsManager) GetParam() (*ssm.GetParameterOutput, error) {
	param, err := sm.client.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("/service-payment/stripe"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}
	return param, nil
}
