package scripts

import (
	"strings"
)

// Format a string by escaping quotes
func EscapeString(content string) string {
	return strings.Replace(content, `"`, `\"`, -1)
}

// Return a script formatted for a command argument
func ScriptToSSMCommands(script string) []string {
	// Strip all carriage returns b/c Windows
	cleanScript := strings.Replace(script, "\r", "", -1)
	return strings.Split(cleanScript, "\n")
}
