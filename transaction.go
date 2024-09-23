package telemetry

// Transaction represents a log transaction
type Transaction struct {
	ID       string
	ParentID string
	Tags     map[string]string
	logger   *LogManager
}

// NewTransaction creates a new transaction.
func NewTransaction(id string, logger *LogManager) *Transaction {
	return &Transaction{
		ID: id,
		Tags: map[string]string{
			"transaction_id": id,
		},
		logger: logger,
	}
}

// AddTag adds a new tag or updates an existing tag for the transaction.
func (t *Transaction) AddTag(key, value string) {
	t.Tags[key] = value
}

// RemoveTag removes a tag from the transaction.
func (t *Transaction) RemoveTag(key string) {
	delete(t.Tags, key)
}

// SetTags update tags for the transaction.
func (t *Transaction) SetTags(tags map[string]string) {
	t.Tags = copyTags(tags)
}

// ResetTags remove all tags for the transaction.
func (t *Transaction) ResetTags() {
	t.Tags = map[string]string{}
}

// Debug logs a message with debug logging level.
func (t *Transaction) Debug(message string) {
	t.logger.Debug(message, t.ParentID, t.ID, t.Tags)
}

// Info logs a message with info logging level.
func (t *Transaction) Info(message string) {
	t.logger.Info(message, t.ParentID, t.ID, t.Tags)
}

// Warning logs a message with warning logging level.
func (t *Transaction) Warning(message string) {
	t.logger.Warning(message, t.ParentID, t.ID, t.Tags)
}

// Error logs a message with error logging level.
func (t *Transaction) Error(message string) {
	t.logger.Error(message, t.ParentID, t.ID, t.Tags)
}

func (t *Transaction) Start() {
	t.Info("Transaction started")
}

func (t *Transaction) End() {
	t.Info("Transaction ended")
}

func (t *Transaction) SubTransaction(subID string) *Transaction {
	// Create a new sub-transaction with the same LogManager
	subTransaction := NewTransaction(subID, t.logger)
	subTransaction.ParentID = t.ID
	// Add parent transaction ID to the sub-transaction tags
	subTransaction.AddTag("parent_transaction_id", t.ID)

	return subTransaction
}

func copyTags(tags map[string]string) map[string]string {
	newTags := make(map[string]string)
	for k, v := range tags {
		newTags[k] = v
	}
	return newTags
}
