package telemetry

import (
	"example/telemetry/drivers"
	"testing"
)

// Test transaction logging.
func TestTransaction_LogTransaction(t *testing.T) {
	mockDriver := NewMockDriver()
	tel := NewLogger()
	tel.SetDriver(mockDriver)

	// Create a transaction
	transaction := NewTransaction("tx-123", tel)

	// Log within the transaction
	transaction.Warning("Processing transaction")

	if len(mockDriver.Logs) != 1 {
		t.Errorf("expected 1 log entry, got %d", len(mockDriver.Logs))
	}

	// Validate the transaction log content
	log := mockDriver.Logs[0]
	if log.Message != "[Transaction tx-123] Processing transaction" {
		t.Errorf("unexpected log message: got '%s'", log.Message)
	}
	if log.Level != drivers.Warning {
		t.Errorf("unexpected log level: got %v", log.Level)
	}
	if log.Tags["transaction_id"] != "tx-123" {
		t.Errorf("expected transaction_id 'tx-123', got '%s'", log.Tags["transaction_id"])
	}
	if _, ok := log.Tags["timestamp"]; !ok {
		t.Errorf("expected log to have timestamp, but it doesn't")
	}
}
