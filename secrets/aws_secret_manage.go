package secrets

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"os"
)

type SecretData struct {
	ApikeyProviderA string
	ApikeyProviderB string
}

var (
	secretName = os.Getenv("SECRET_NAME")
	region     = os.Getenv("AWS_REGION")
)

var DefaultSecret = SecretData{
	ApikeyProviderB: "ApikeyProviderB",
	ApikeyProviderA: "ApikeyProviderA",
}

func GetSecret() (secretData SecretData) {
	secretData = DefaultSecret
	svc := secretsmanager.New(
		session.New(),
		aws.NewConfig().WithRegion(region),
	)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return
	}

	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	}

	err = json.Unmarshal([]byte(secretString), &secretData)
	if err != nil {
		return
	}

	return secretData
}
