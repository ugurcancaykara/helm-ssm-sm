package internal

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"log"
	"strings"
)

func GetSSMParam(paramPath string, region string) (string, error) {
	var cfg aws.Config
	var err error
	if region != "" {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
			return "", nil

		}
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
			return "", nil

		}
	}
	svc := ssm.NewFromConfig(cfg)
	param, err := svc.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           &paramPath,
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Fatal(err)
		return "", nil
	}
	value := *param.Parameter.Value
	return value, nil
}

func ssmFunc(args ...string) string {
	path := args[0]
	var region string
	if len(args) > 1 {
		region = strings.Split(args[1], "=")[1]
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	value, err := GetSSMParam(path, region)
	if err != nil {
		log.Fatalln(err)
	}
	return value
}
