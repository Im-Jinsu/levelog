package example_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/im-jinsu/easy-logging-level/logger"
	"gopkg.in/natefinch/lumberjack.v2"
)

// TestLogger : Test logger package
func TestLogger(t *testing.T) {
	// Set stdlogger with lumberjack
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     5, //days
	})

	logger.Info("1. logger.Info")
	logger.Error("2. logger.Error")
	logger.ErrorAndSend("3. logger.ErrorAndSend")
	logger.Debug("4. logger.Debug")
	// logger.Fatal("Cannot test fatal")

	if bytes, err := ioutil.ReadFile("./test.log"); err != nil {
		panic(err)
	} else {
		fmt.Println(string(bytes))
	}
	if err := os.Remove("./test.log"); err != nil {
		panic(err)
	}
}
