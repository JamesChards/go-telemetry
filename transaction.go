package telemetry

import (
	"fmt"
)

// Transaction represents a log transaction
type Transaction struct {
	ID        string
	Tags      map[string]string
	telemetry *Logger
}

// NewTransaction creates a new transaction.
func NewTransaction(id string, telemetry *Logger) *Transaction {
	t := &Transaction{
		ID: id,
		Tags: map[string]string{
			"transaction_id": id,
		},
		telemetry: telemetry,
	}
	t.telemetry.SetTags(t.Tags)
	return t
}

// AddTag adds a new tag or updates an existing tag for the transaction.
func (t *Transaction) AddTag(key, value string) {
	t.Tags[key] = value
	t.telemetry.SetTags(t.Tags) // Update telemetry with the new tags
}

// RemoveTag removes a tag from the transaction.
func (t *Transaction) RemoveTag(key string) {
	delete(t.Tags, key)
	t.telemetry.SetTags(t.Tags) // Update telemetry after tag removal
}

// SetTags update tags for the transaction.
func (t *Transaction) SetTags(tags map[string]string) {
	t.Tags = map[string]string{}
	for k, v := range tags {
		t.Tags[k] = v
	}
	t.telemetry.SetTags(t.Tags)
}

// ResetTags remove all tags for the transaction.
func (t *Transaction) ResetTags() {
	t.Tags = map[string]string{}
	t.telemetry.SetTags(t.Tags)
}

// Debug logs a message with debug logging level.
func (t *Transaction) Debug(message string) {
	t.telemetry.Debug(fmt.Sprintf("[Transaction %s] %s", t.ID, message))
}

// Debug logs a message with debug logging level.
func (t *Transaction) Info(message string) {
	t.telemetry.Info(fmt.Sprintf("[Transaction %s] %s", t.ID, message))
}

// Debug logs a message with debug logging level.
func (t *Transaction) Warning(message string) {
	t.telemetry.Warning(fmt.Sprintf("[Transaction %s] %s", t.ID, message))
}

// Debug logs a message with debug logging level.
func (t *Transaction) Error(message string) {
	t.telemetry.Error(fmt.Sprintf("[Transaction %s] %s", t.ID, message))
}

func (t *Transaction) Start() {
	t.Info("Transaction started")
}

func (t *Transaction) End() {
	t.Info("Transaction ended")
}

func (t *Transaction) StartSubTransaction(subID string) *Transaction {
	subTransaction := NewTransaction(subID, t.telemetry)
	subTransaction.AddTag("parent_transaction_id", t.ID)
	return subTransaction
}
