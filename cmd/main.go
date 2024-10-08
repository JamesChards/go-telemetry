package main

import (
	"fmt"

	telemetry "github.com/JamesChards/go-telemetry"
)

func main() {
	err := telemetry.CreateDefaultConfig()
	if err != nil {
		fmt.Println(err)
	}

	myLog := telemetry.NewLogger()

	if closer, ok := myLog.IsCloser(); ok {
		defer func() {
			if cerr := closer.Close(); cerr != nil {
				fmt.Println("Error closing text log file:", cerr)
			}
		}()
	}

	myTrans := telemetry.NewTransaction("a3124", myLog)
	myTrans.AddTag("a", "Hello")
	myTrans.AddTag("b", "Hi")

	myTrans.Start()
	myTrans.Debug("This log uses the default driver")
	myTrans.SetTags(map[string]string{
		"first":  "hello",
		"second": "hi",
	})
	myTrans.Debug("This is the second log")
	myTrans.End()

	subTrans := myTrans.SubTransaction("subtrans")
	subTrans.Start()
	subTrans.AddTag("abc", "Hello")
	subTrans.End()
}
