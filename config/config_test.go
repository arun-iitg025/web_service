package config

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	// Set a temporary environment variable
	os.Setenv("TEST_ENV", "test_value")
	defer os.Unsetenv("TEST_ENV")

	// Test that GetEnv retrieves the environment variable
	if value := GetEnv("TEST_ENV", "default"); value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}

	// Test that GetEnv uses the fallback when the variable is not set
	if value := GetEnv("NON_EXISTENT_ENV", "default"); value != "default" {
		t.Errorf("Expected 'default', got '%s'", value)
	}
}
