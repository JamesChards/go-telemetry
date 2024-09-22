package telemetry

import (
	"example/telemetry/drivers"
	"strconv"
	"sync"
	"testing"
)

// Test concurrency with TextFileDriver
func TestConcurrentLogging(t *testing.T) {
	driver, err := drivers.NewTextFileDriver("test.txt")
	if err != nil {
		t.Fatalf("Error creating TextFileDriver: %v", err)
	}
	defer driver.Close()

	logger := NewLogger()
	logger.SetDriver(driver)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			logger.SetTags(map[string]string{"id": strconv.Itoa(id)})
			logger.Info("Logging from goroutine")
		}(i)
	}

	wg.Wait()
}
