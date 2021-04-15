package example_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/im-jinsu/levelog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// TestLogger : Test levelog package
func TestLogger(t *testing.T) {
	// Set stdlogger with lumberjack
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     5, //days
	})

	levelog.Info("1. levelog.Info")
	levelog.Error("2. levelog.Error")
	levelog.ErrorAndSend("3. levelog.ErrorAndSend")
	levelog.Debug("4. levelog.Debug")
	// levelog.Fatal("Cannot test fatal")

	if bytes, err := ioutil.ReadFile("./test.log"); err != nil {
		panic(err)
	} else {
		fmt.Println(string(bytes))
	}
	if err := os.Remove("./test.log"); err != nil {
		panic(err)
	}
}
