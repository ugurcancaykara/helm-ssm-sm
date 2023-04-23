package internal

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"log"
	"strings"
)

func smFunc(args ...string) any {
	key := args[0]
	var region string
	if len(args) > 1 {
		region = strings.Split(args[1], "=")[1]
	}

	value := GetSecret(key, region)
	return value
}

func GetSecret(secretName string, region string) any {
	var cfg aws.Config
	var err error
	if region != "" {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}
	}
	svc := secretsmanager.NewFromConfig(cfg)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return ""
	}

	return *result.SecretString

}
