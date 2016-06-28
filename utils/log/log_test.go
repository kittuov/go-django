package log

import (
	"log"
	"strings"
	"testing"
)

type testByteBuffer string

func (l *testByteBuffer) Write(p []byte) (int, error) {
	*l = testByteBuffer(string(*l) + string(p))
	return len(p), nil
}

func TestLog(t *testing.T) {
	b := testByteBuffer("")
	log.SetOutput(&b)
	baseLog(VERB, "Verbose")
	lev := levels.getLevel(VERB)
	if !strings.Contains(string(b), lev.Prefix) && !strings.Contains(string(b), "Verbose") {
		t.Errorf("The printed log doesnot contain expected prefix %s", lev.Prefix)
	}
}

func TestSetMin(t *testing.T) {
	SetMin(1)
	b := testByteBuffer("")
	log.SetOutput(&b)
	baseLog(VERB, "Verbose")
	if string(b) != "" {
		t.Error("minimum set to Debug but Verbose logs are printing")
	}
	baseLog(DEBU, "Debug")
	baseLog(INFO, "Info")
	if !strings.Contains(string(b), "DEBU") && !strings.Contains(string(b), "INFO") {
		t.Error("minimum is set to Debug but debug or info logs are not printing")
	}
}

func TestSetMin2(t *testing.T) {
	err := SetMin(-1)
	if err == nil {
		t.Error("given wrong level but no error is returned")
	}
}
