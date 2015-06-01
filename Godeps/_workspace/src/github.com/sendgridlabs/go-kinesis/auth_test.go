package kinesis

import (
	"os"
	"testing"
)

func TestGetSecretKey(t *testing.T) {
	auth := NewAuth("BAD_ACCESS_KEY", "BAD_SECRET_KEY")

	if auth.GetAccessKey() != "BAD_ACCESS_KEY" {
		t.Error("incorrect value for auth#accessKey")
	}
}

func TestGetAccessKey(t *testing.T) {
	auth := NewAuth("BAD_ACCESS_KEY", "BAD_SECRET_KEY")

	if auth.GetSecretKey() != "BAD_SECRET_KEY" {
		t.Error("incorrect value for auth#secretKey")
	}
}

func TestNewAuthFromEnv(t *testing.T) {
	os.Setenv(ACCESS_ENV_KEY, "asdf")
	os.Setenv(SECRET_ENV_KEY, "asdf")

	auth, _ := NewAuthFromEnv()

	if auth.GetAccessKey() != "asdf" {
		t.Error("Expected AccessKey to be inferred as \"asdf\"")
	}

	if auth.GetSecretKey() != "asdf" {
		t.Error("Expected SecretKey to be inferred as \"asdf\"")
	}

	os.Setenv(ACCESS_ENV_KEY, "") // Use Unsetenv with go1.4
	os.Setenv(SECRET_ENV_KEY, "") // Use Unsetenv with go1.4
}
