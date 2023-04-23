package internal

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/stretchr/testify/assert"

	"testing"
)

type SecretsManagerSecret struct {
	Value         *string
	DefaultValue  *string
	ExpectedValue *string
}

var (
	fakeSecretValue        = "fakeValue"
	fakeSecretOtherValue   = "fakeOtherValue"
	fakeSecretMissingValue = "fakeMissingValue"
	fakeSecretStore        = map[string]SecretsManagerSecret{
		"existing-secret": {
			Value:         &fakeSecretValue,
			DefaultValue:  &fakeSecretOtherValue,
			ExpectedValue: &fakeSecretValue,
		},
	}
)

func TestGetSecretsManagerSecret(t *testing.T) {
	// Mock Secrets Manager client
	mockSecretsManagerClient := &mockSecretsManagerClient{}

	// Initialize fake store
	fakeSecretsManagerStore := fakeSecretStore

	for k, v := range fakeSecretsManagerStore {
		expectedValueStr := "nil"
		if v.ExpectedValue != nil {
			expectedValueStr = *v.ExpectedValue
		}

		t.Logf("Key: %s should have value: %s", k, expectedValueStr)

		// Call GetSecret function
		value, err := mockSecretsManagerClient.GetSecretValue(k)
		if err != nil {
			fmt.Println("error:", err)
		}

		returnedValue := string(value.SecretBinary)

		if v.ExpectedValue == nil {
			fmt.Println("nil")
			assert.Equal(t, []byte(nil), value.SecretBinary)
		} else {
			assert.EqualValues(t, []byte(expectedValueStr), []byte(returnedValue))
		}
	}
}

type mockSecretsManagerClient struct{}

func (m *mockSecretsManagerClient) GetSecretValue(secretname string) (*secretsmanager.GetSecretValueOutput, error) {

	secret, ok := fakeSecretStore[secretname]
	if !ok {
		return nil, awserr.New(secretsmanager.ErrCodeResourceNotFoundException, "Secret not found", nil)
	}
	if secret.Value == nil {
		return &secretsmanager.GetSecretValueOutput{}, nil
	}

	// Return as byte array instead of string
	output := &secretsmanager.GetSecretValueOutput{
		SecretBinary: []byte(*secret.Value),
	}
	return output, nil
}
