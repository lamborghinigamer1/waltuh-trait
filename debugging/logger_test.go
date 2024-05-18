package debugging

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestLogger_Info(t *testing.T) {
	old := os.Stdout // Keep original stdout
	r, w, _ := os.Pipe()
	os.Stdout = w // Replace stdout with pipe writer

	logger := NewLogger("TestLogger")

	logger.Info("This is an info message")
	expectedPrefix := "[INFO - "

	w.Close() // Close the pipe writer
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r) // Read from pipe reader

	os.Stdout = old // Restore original stdout

	if !strings.Contains(buf.String(), expectedPrefix) {
		t.Errorf("Info message doesn't have expected prefix. Got: %s, Expected Prefix: %s", buf.String(), expectedPrefix)
	}
}

func TestLogger_Debug(t *testing.T) {
	old := os.Stdout // Keep original stdout
	r, w, _ := os.Pipe()
	os.Stdout = w // Replace stdout with pipe writer

	logger := NewLogger("TestLogger")

	logger.Debug("This is an debug message")
	expectedPrefix := "[DEBUG - "

	w.Close() // Close the pipe writer
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r) // Read from pipe reader

	os.Stdout = old // Restore original stdout

	if !strings.Contains(buf.String(), expectedPrefix) {
		t.Errorf("Info message doesn't have expected prefix. Got: %s, Expected Prefix: %s", buf.String(), expectedPrefix)
	}
}

func TestLogger_Error(t *testing.T) {
	old := os.Stdout // Keep original stdout
	r, w, _ := os.Pipe()
	os.Stdout = w // Replace stdout with pipe writer

	logger := NewLogger("TestLogger")

	logger.Error("This is an error message")
	expectedPrefix := "[ERROR - "

	w.Close() // Close the pipe writer
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r) // Read from pipe reader

	os.Stdout = old // Restore original stdout

	if !strings.Contains(buf.String(), expectedPrefix) {
		t.Errorf("Info message doesn't have expected prefix. Got: %s, Expected Prefix: %s", buf.String(), expectedPrefix)
	}
}

func TestLogger_Warning(t *testing.T) {
	old := os.Stdout // Keep original stdout
	r, w, _ := os.Pipe()
	os.Stdout = w // Replace stdout with pipe writer

	logger := NewLogger("TestLogger")

	logger.Warning("This is an error message")
	expectedPrefix := "[WARNING - "

	w.Close() // Close the pipe writer
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r) // Read from pipe reader

	os.Stdout = old // Restore original stdout

	if !strings.Contains(buf.String(), expectedPrefix) {
		t.Errorf("Info message doesn't have expected prefix. Got: %s, Expected Prefix: %s", buf.String(), expectedPrefix)
	}
}

func TestLogger_Fatal(t *testing.T) {
	old := os.Stdout // Keep original stdout
	r, w, _ := os.Pipe()
	os.Stdout = w // Replace stdout with pipe writer

	logger := NewLogger("TestLogger")

	logger.Fatal("This is an error message")
	expectedPrefix := "[FATAL - "

	w.Close() // Close the pipe writer
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r) // Read from pipe reader

	os.Stdout = old // Restore original stdout

	if !strings.Contains(buf.String(), expectedPrefix) {
		t.Errorf("Info message doesn't have expected prefix. Got: %s, Expected Prefix: %s", buf.String(), expectedPrefix)
	}
}
