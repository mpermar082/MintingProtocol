// internal/mintingprotocol/mintingprotocol_test.go
package mintingprotocol

import (
    "testing"
)

// TestNewApp verifies the NewApp function returns a non-nil app with correct properties
func TestNewApp(t *testing.T) {
    app := NewApp(true)
    if app == nil {
        t.Fatal("NewApp returned nil")
    }
    if !app.Verbose {
        t.Error("Expected verbose to be true")
    }
    if app.ProcessedCount != 0 {
        t.Errorf("Expected ProcessedCount to be 0, got %d", app.ProcessedCount)
    }
}

// TestProcess verifies the Process function returns a successful result and updates ProcessedCount
func TestProcess(t *testing.T) {
    app := NewApp(false)
    result, err := app.Process("test data")
    if err != nil {
        t.Fatalf("Process returned error: %v", err)
    }
    if !result.Success {
        t.Error("Expected result.Success to be true")
    }
    if app.ProcessedCount != 1 {
        t.Errorf("Expected ProcessedCount to be 1, got %d", app.ProcessedCount)
    }
}

// TestRun verifies the Run function returns no error
func TestRun(t *testing.T) {
    app := NewApp(false)
    err := app.Run("", "")
    if err != nil {
        t.Fatalf("Run returned error: %v", err)
    }
}