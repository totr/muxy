package log

import (
	"log"
	"testing"
)

type FakeWriter struct {
	Count int
}

func (w *FakeWriter) Write(p []byte) (n int, err error) {
	w.Count = w.Count + 1
	return 0, nil
}

func TestLogLevel(t *testing.T) {

	writer := &FakeWriter{}
	logger := &MuxyLogger{Level: DEBUG}
	log.SetOutput(writer)
	logger.Log(INFO, "Info")
	logger.Log(TRACE, "Trace") // Should not appear
	if writer.Count != 1 {
		t.Fatalf("Logger should have written once, but received %d calls", writer.Count)
	}
	logger.Debug("Debug")

	if writer.Count != 2 {
		t.Fatalf("Logger should have written twice, only received %d calls", writer.Count)
	}

	writer.Count = 0
	logger.Trace("yo")
	logger.Debug("yo")
	logger.Info("yo")
	logger.Error("yo")
	logger.Warn("yo")
	if writer.Count != 4 {
		t.Fatalf("Logger should have written 4 times, received %d calls", writer.Count)
	}

	logger.Level = TRACE
	writer.Count = 0
	logger.Trace("yo")
	if writer.Count != 1 {
		t.Fatalf("Logger should have written once, received %d calls", writer.Count)
	}
}

func TestLog_Level(t *testing.T) {
	levels := []Level{
		TRACE,
		INFO,
		WARN,
		ERROR,
		FATAL,
	}
	logger := MuxyLogger{Level: NONE}

	for _, level := range levels {
		logger.SetLevel(level)
		logger.Trace("Test_LogLevel - Trace")
		logger.Debug("Test_LogLevel - Debug")
		logger.Info("Test_LogLevel - Info")
		logger.Warn("Test_LogLevel - Warn")
		logger.Error("Test_LogLevel - Error")
		logger.Log(level, "TestLog_Level")
	}
}

func TestLog_Functions(t *testing.T) {
	SetLevel(NONE)
	Trace("TestLog_Functions - Trace")
	Debug("TestLog_Functions - Debug")
	Info("TestLog_Functions - Info")
	Warn("TestLog_Functions - Warn")
	Error("TestLog_Functions - Error")
	Log(ERROR, "TestLog_Functions - ERROR")
}

func TestLogColour(t *testing.T) {
	logger := &MuxyLogger{Level: DEBUG}
	logger.Log(INFO, "Info %s", Colorize(LIGHTRED, " some words "))
	logger.Log(INFO, "Info something else not in colour")
}
