package port

import "github.com/aws/aws-sdk-go/service/ssm"

type SecretsManager interface {
	GetParam() (*ssm.GetParameterOutput, error)
}
