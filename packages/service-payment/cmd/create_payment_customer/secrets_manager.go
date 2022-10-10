package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type SecretsManager struct {
	client *ssm.SSM
}

func initialize() *SecretsManager {
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

func (sm *SecretsManager) getParam() (*ssm.GetParameterOutput, error) {
	param, err := sm.client.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("/dev/service-payment/stripe"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
		return nil, err
	}
	fmt.Println("PARAM: ", param)
	return param, nil
}