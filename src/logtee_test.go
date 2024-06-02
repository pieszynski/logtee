package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGetLogFileName(t *testing.T) {
    now := time.Now()
    shouldBe := fmt.Sprintf("abc.%s.log", now.Format("20060102"))
    logFile := getLogFileName("abc.log", now)
    if logFile != shouldBe {
        t.Fatalf(`getLogFileName("abc.log") = %s, want "%s", error`, logFile, shouldBe)
    }
}
