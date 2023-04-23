package internal

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"gotest.tools/v3/assert"
	"testing"
)

type SSMParameter struct {
	value         *string
	defaultValue  *string
	expectedValue *string
}

var (
	fakeValue         = "production"
	fakeOtherValue    = "fake-other-value"
	fakeMissingValue  = "fake-missing-value"
	fakeSSMParameter1 = SSMParameter{
		value:         aws.String(fakeValue),
		defaultValue:  aws.String(fakeOtherValue),
		expectedValue: aws.String(fakeValue),
	}
	fakeSSMParameter2 = SSMParameter{
		value:         aws.String(""),
		defaultValue:  aws.String(fakeOtherValue),
		expectedValue: aws.String(fakeOtherValue),
	}
	fakeSSMParameter3 = SSMParameter{
		value:         nil,
		defaultValue:  aws.String(fakeOtherValue),
		expectedValue: nil,
	}
	fakeStore = map[string]SSMParameter{
		"/infra/environment": fakeSSMParameter1,
	}
)

func TestGetSSMParameter(t *testing.T) {
	for k, v := range fakeStore {
		expectedValueStr := "nil"
		if v.expectedValue != nil {
			expectedValueStr = *v.expectedValue
		}

		fmt.Println("v.expectedValue: " + *v.expectedValue)
		t.Logf("Key: %s should have value: %s", k, expectedValueStr)

		value, err := GetSSMParam(k, "eu-west-1")
		if err != nil {
			fmt.Println("error:", err)
		}
		assert.Equal(t, *v.expectedValue, value)
		if v.expectedValue == nil {
			assert.Error(t, err, "ParameterNotFound: Parameter does not exist in SSM")
		}

	}

}
