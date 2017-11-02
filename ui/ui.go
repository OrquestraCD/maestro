package ui

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

// Modes
const (
	Standard = iota
	Debug
	Silent
)

var (
	UI       *DefaultUI
	stdoutMu *sync.Mutex = &sync.Mutex{}
	stderrMu *sync.Mutex = &sync.Mutex{}
)

type DefaultUI struct {
	StdOut          io.Writer
	StdErr          io.Writer
	DbgOut          io.Writer
	StdIn           io.Reader
	InteractiveMode bool
}

// Provides more context for errors when possible
func ErrorWithContext(err error) error {
	return err
}

func InitUI(mode int) {
	if mode == Silent {
		UI = &DefaultUI{
			StdOut:          ioutil.Discard,
			StdErr:          ioutil.Discard,
			DbgOut:          ioutil.Discard,
			StdIn:           os.Stdin,
			InteractiveMode: false,
		}
		return
	}

	UI = &DefaultUI{
		StdOut:          os.Stdout,
		StdErr:          os.Stderr,
		StdIn:           os.Stdin,
		InteractiveMode: true,
	}

	if mode == Debug {
		UI.DbgOut = os.Stdout
	} else {
		// Since debug mode is not on discard
		UI.DbgOut = ioutil.Discard
	}
}

func (u DefaultUI) Print(text string) {
	stdoutMu.Lock()
	defer stdoutMu.Unlock()
	fmt.Fprintln(u.StdOut, text)
}

func (u DefaultUI) Printf(format string, a ...interface{}) {
	stdoutMu.Lock()
	defer stdoutMu.Unlock()
	fmt.Fprintf(u.StdOut, format, a...)
}

func (u DefaultUI) Error(text string) {
	stderrMu.Lock()
	defer stderrMu.Unlock()
	fmt.Fprintln(u.StdErr, "ERROR:", text)
}

func (u DefaultUI) Errorf(format string, a ...interface{}) {
	stderrMu.Lock()
	defer stderrMu.Unlock()
	fmt.Fprint(u.StdErr, "ERROR: ", fmt.Sprintf(format, a...))
}

func (u DefaultUI) Debug(text string) {
	stdoutMu.Lock()
	defer stdoutMu.Unlock()
	fmt.Fprintln(u.DbgOut, "DEBUG:", text)
}

func (u DefaultUI) Debugf(format string, a ...interface{}) {
	stdoutMu.Lock()
	defer stdoutMu.Unlock()
	fmt.Fprint(u.DbgOut, "DEBUG: ", fmt.Sprintf(format, a...))
}

func (u DefaultUI) Scanf(name, format string, result ...interface{}) error {
	// If InteractiveMode is not set then error out that information is missing
	if !u.InteractiveMode {
		return InteractiveError()
	}

	_, err := fmt.Fscanf(u.StdIn, format, result...)

	return err
}

// Scans a line up to a newline and stores it in a string
func (u DefaultUI) Scanln(name string, result *string) error {
	// If InteractiveMode is not set then error out that information is missing
	if !u.InteractiveMode {
		return InteractiveError()
	}
	reader := bufio.NewReader(u.StdIn)

	rawStr, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	*result = strings.Trim(strings.Trim(rawStr, "\r"), "\n")
	return nil
}

// Prompt user for string input
func (u DefaultUI) PromptString(name, text string, result *string) error {
	if !u.InteractiveMode {
		return InteractiveError()
	}

	u.Printf(text)
	if err := u.Scanf(name, "%s\n", result); err != nil {
		return err
	}
	// Print extra newline for spacing
	u.Print("")

	return nil
}

// Prompt user for int input
func (u DefaultUI) PromptInt(name, text string, result *int) error {
	if !u.InteractiveMode {
		return InteractiveError()
	}
	u.Printf(text)
	if err := u.Scanf(name, "%d\n", result); err != nil {
		return err
	}
	// Print extra newline for spacing
	u.Print("")

	return nil
}

func InteractiveError() error {
	return errors.New("stdin ignored because --no-interactive flag is set\n")
}
