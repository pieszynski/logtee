package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

    sti, stiError := os.Stdin.Stat()
    if stiError != nil || sti.Mode() & os.ModeNamedPipe == 0 {
        log.Fatal("No data redirected from unix pipe or error: ", stiError)
        os.Exit(11)
    }

    if len(os.Args) < 2 {
        log.Fatal("App requires one parameter: filename to log to")
        os.Exit(7)
    }

    fileName := os.Args[1]

    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
        appendLog(fileName, scanner.Bytes())
    }
}

func appendLog(fileName string, data []byte) {

    data = append(data, []byte("\n")...)

    fLogFileName := getLogFileName(fileName, time.Now())

    f, err := os.OpenFile(fLogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

    if err != nil {
        log.Fatal(err)
    }

    if _, err := f.Write([]byte(data)); err != nil {
        f.Close() // ignore error; Write error takes precedence
        log.Fatal(err)
    }

    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}

func getLogFileName(fileName string, now time.Time) (string) {
    fExt := filepath.Ext(fileName) 
    fBase := strings.TrimSuffix(filepath.Base(fileName), fExt)
    fLogFileName := fmt.Sprintf("%s.%s%s", fBase, now.Format("20060102"), fExt)
    return fLogFileName
}