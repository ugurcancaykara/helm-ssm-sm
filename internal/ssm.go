package internal

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"log"
)

func getSSMParam(paramPath string) any {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	svc := ssm.NewFromConfig(cfg)
	param, err := svc.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           &paramPath,
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Fatal(err)
	}
	value := *param.Parameter.Value
	return value
}
