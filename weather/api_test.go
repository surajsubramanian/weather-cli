package weather

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSaveLoadWeather(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "")
	fp := filepath.Join(tempDir, "test.json")
	if err != nil {
		t.Fatal("Failed to create temporary directory")
	}
	defer os.RemoveAll(tempDir)

	input := testWeatherInput("2022-04-05")
	output := &Weather{}

	err = Save[Weather](input, fp)
	output, err = Load[Weather](fp)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(input, *output) {
		t.Fatalf("Expected %v, got %v", input, output)
	}
}
