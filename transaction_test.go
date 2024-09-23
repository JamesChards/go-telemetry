package telemetry

import (
	"testing"
)

func TestTransaction_StartEnd(t *testing.T) {
	logger := NewLogger("cli")
	trans := NewTransaction("trans123", logger)

	trans.Start()
	trans.End()
	// There's no direct output or return, but we ensure it doesn't panic or error.
}

func TestTransaction_AddAndRemoveTags(t *testing.T) {
	logger := NewLogger("cli")
	trans := NewTransaction("trans123", logger)

	trans.AddTag("user", "admin")
	if trans.Tags["user"] != "admin" {
		t.Errorf("Expected tag 'user' to be 'admin', got %s", trans.Tags["user"])
	}

	trans.RemoveTag("user")
	if _, ok := trans.Tags["user"]; ok {
		t.Error("Expected 'user' tag to be removed, but it's still present")
	}
}

func TestTransaction_SubTransaction(t *testing.T) {
	logger := NewLogger("cli")
	parentTrans := NewTransaction("parent123", logger)

	subTrans := parentTrans.SubTransaction("sub456")
	if subTrans.ParentID != parentTrans.ID {
		t.Errorf("Expected sub-transaction ParentID to be '%s', got '%s'", parentTrans.ID, subTrans.ParentID)
	}
}
