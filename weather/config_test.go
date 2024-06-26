package weather

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveLoadConfig(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "")
	fp := filepath.Join(tempDir, "test.json")
	if err != nil {
		t.Fatal("Failed to create temporary directory")
	}
	defer os.RemoveAll(tempDir)

	input := Config{
		Key:      "testKey",
		Location: "testLocation",
	}
	output := &Config{}

	err = Save[Config](input, fp)
	output, err = Load[Config](fp)
	if err != nil {
		t.Fatal(err)
	}
	if input != *output {
		t.Fatalf("Expected %v, got %v", input, output)
	}
}
