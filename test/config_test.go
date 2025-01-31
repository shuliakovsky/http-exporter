package test

import (
	"reflect"
	"testing"

	"http-exporter/internal/config"
)

func TestReadConfig(t *testing.T) {
	expectedConfig := config.Config{
		Port:      8080,
		Interface: "0.0.0.0",
		Interval:  "30s",
		URLs:      []string{"http://example.com", "http://anotherexample.com"},
	}

	cfg, err := config.ReadConfig("../config.yaml")
	if err != nil {
		t.Fatalf("Error reading config: %v", err)
	}

	if !reflect.DeepEqual(cfg, expectedConfig) {
		t.Errorf("Expected config %+v, got %+v", expectedConfig, cfg)
	}
}
