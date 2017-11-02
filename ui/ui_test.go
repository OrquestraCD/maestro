package ui

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestInitUI_standard(t *testing.T) {
	want := DefaultUI{
		StdOut:          os.Stdout,
		StdErr:          os.Stderr,
		DbgOut:          ioutil.Discard,
		StdIn:           os.Stdin,
		InteractiveMode: true,
	}

	InitUI(Standard)

	if !reflect.DeepEqual(*UI, want) {
		t.Errorf("Expected: %+v\nGot: %+v\n", want, *UI)
	}
}

func TestInitUI_debug(t *testing.T) {
	want := DefaultUI{
		StdOut:          os.Stdout,
		StdErr:          os.Stderr,
		DbgOut:          os.Stdout,
		StdIn:           os.Stdin,
		InteractiveMode: true,
	}

	InitUI(Debug)

	if !reflect.DeepEqual(*UI, want) {
		t.Errorf("Expected: %+v\nGot: %+v\n", want, *UI)
	}
}

func TestInitUI_silent(t *testing.T) {
	want := DefaultUI{
		StdOut:          ioutil.Discard,
		StdErr:          ioutil.Discard,
		DbgOut:          ioutil.Discard,
		StdIn:           os.Stdin,
		InteractiveMode: false,
	}

	InitUI(Silent)

	if !reflect.DeepEqual(*UI, want) {
		t.Errorf("Expected: %+v\nGot: %+v\n", want, *UI)
	}
}

func TestPrint(t *testing.T) {
	want := "test string\n"
	var buffer bytes.Buffer
	ui := DefaultUI{StdOut: &buffer}

	ui.Print("test string")

	got := buffer.String()
	if got != want {
		t.Errorf("Expected \"%s\", got \"%s\"", want, got)
	}
}

func TestPrintf(t *testing.T) {
	want := "test string"
	var buffer bytes.Buffer
	ui := DefaultUI{StdOut: &buffer}

	ui.Printf("test string")

	got := buffer.String()
	if got != want {
		t.Errorf("Expected \"%s\", got \"%s\"", want, got)
	}
}

func TestError(t *testing.T) {
	want := "ERROR: test error\n"
	var buffer bytes.Buffer
	ui := DefaultUI{StdErr: &buffer}

	ui.Error("test error")
	got := buffer.String()
	if got != want {
		t.Errorf("Expected \"%s\", got \"%s\"", want, got)
	}
}

func TestErrorf(t *testing.T) {
	want := "ERROR: test error"
	var buffer bytes.Buffer
	ui := DefaultUI{StdErr: &buffer}

	ui.Errorf("test error")
	got := buffer.String()
	if got != want {
		t.Errorf("Expected \"%s\", got \"%s\"", want, got)
	}
}

func TestDebug(t *testing.T) {
	want := "DEBUG: test debug\n"
	var buffer bytes.Buffer
	ui := DefaultUI{DbgOut: &buffer}

	ui.Debug("test debug")
	got := buffer.String()
	if got != want {
		t.Errorf("Expected \"%s\", got \"%s\"", want, got)
	}
}

func TestDebugf(t *testing.T) {
	want := "DEBUG: test debug"
	var buffer bytes.Buffer
	ui := DefaultUI{DbgOut: &buffer}

	ui.Debugf("test debug")
	got := buffer.String()
	if got != want {
		t.Errorf("Expected \"%s\", got \"%s\"", want, got)
	}
}

func TestScanf(t *testing.T) {
	want := "test"
	var got string

	reader := bytes.NewBufferString("test")
	ui := DefaultUI{InteractiveMode: true, StdIn: reader}

	err := ui.Scanf("test", "%s", &got)
	if err != nil {
		t.Errorf(err.Error())
	}

	if got != want {
		t.Errorf("Expected \"%s\", got \"%s\"", want, got)
	}
}

func TestScanf_interactive_off(t *testing.T) {
	reader := bytes.NewBufferString("test")
	ui := DefaultUI{InteractiveMode: false, StdIn: reader}

	want := fmt.Errorf("stdin ignored because --no-interactive flag is set\n")

	got := ui.Scanf("test", "%s", &bytes.Buffer{})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected error \"%v\", got \"%v\"", want, got)
	}
}
