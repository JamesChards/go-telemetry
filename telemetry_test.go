package telemetry

import (
	"testing"

	"github.com/JamesChards/go-telemetry/drivers"
)

func TestLogManager_SetDriverWithName(t *testing.T) {
	lm := NewLogger("cli")

	lm.SetDriverWithName("text")
	if _, ok := lm.driver.(*drivers.TextFileDriver); !ok {
		t.Errorf("Expected TextFileDriver, got %T", lm.driver)
	}

	lm.SetDriverWithName("json")
	if _, ok := lm.driver.(*drivers.JSONDriver); !ok {
		t.Errorf("Expected JSONDriver, got %T", lm.driver)
	}

	lm.SetDriverWithName("cli")
	if _, ok := lm.driver.(*drivers.CLIDriver); !ok {
		t.Errorf("Expected CLIDriver, got %T", lm.driver)
	}
}

func TestLogManager_Log(t *testing.T) {
	lm := NewLogger("cli")

	tags := map[string]string{"user": "admin"}
	lm.AddTag("env", "test")

	lm.Log(drivers.Info, "Test log message", "", "1234", tags)

	// Since the actual Log method in CLIDriver uses stdout, it's hard to capture its output,
	// but this ensures that the method runs without error.
}

func TestLogManager_AddAndRemoveTags(t *testing.T) {
	lm := NewLogger("cli")

	lm.AddTag("env", "production")
	if lm.tags["env"] != "production" {
		t.Errorf("Expected tag 'env' to be 'production', got %s", lm.tags["env"])
	}

	lm.RemoveTag("env")
	if _, ok := lm.tags["env"]; ok {
		t.Error("Expected 'env' tag to be removed, but it's still present")
	}
}

func TestLogManager_SetTags(t *testing.T) {
	lm := NewLogger("cli")

	tags := map[string]string{"env": "production", "user": "admin"}
	lm.SetTags(tags)

	if len(lm.tags) != 2 {
		t.Errorf("Expected 2 tags, got %d", len(lm.tags))
	}

	if lm.tags["env"] != "production" || lm.tags["user"] != "admin" {
		t.Error("Tags were not set correctly")
	}

	lm.ResetTags()

	if len(lm.tags) != 0 {
		t.Errorf("Expected 0 tags after reset, got %d", len(lm.tags))
	}
}
