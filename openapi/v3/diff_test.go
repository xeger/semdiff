package v3_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	v3 "github.com/xeger/semdiff/openapi/v3"
)

func loadSpec(t *testing.T, name string) *v3.OpenAPI {
	t.Helper()

	file, err := os.Open("testdata/" + name)
	if err != nil {
		t.Fatalf("failed to open spec: %v", err)
	}
	defer file.Close()

	var spec v3.OpenAPI
	if err := json.NewDecoder(file).Decode(&spec); err != nil {
		t.Fatalf("failed to decode spec: %v", err)
	}

	return &spec
}

func TestDiff(t *testing.T) {
	old, new := loadSpec(t, "2024-05-17.json"), loadSpec(t, "2025-05-17.json")
	diffs := v3.Diff(old, new)
	assert.True(t, diffs.Major, "expected major revision")
	assert.NotEmpty(t, diffs.Details, "expected some diffs")
}
