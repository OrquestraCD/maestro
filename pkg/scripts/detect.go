package scripts

import "strings"

const (
	PowerShell = iota
	Bash
	Unknown
)

func ScriptTypeByName(name string) int {
	switch name {
	case "bash", "Bash":
		return Bash
	case "powershell", "PowerShell":
		return PowerShell
	}

	return Unknown
}

// Given a script name return the resulting extension
func DetectScriptByExtension(name string) int {
	spltName := strings.Split(name, ".")
	switch spltName[len(spltName)-1] {
	case "bash":
		fallthrough
	case "sh":
		return Bash
	case "psd1":
		fallthrough
	case "psm1":
		fallthrough
	case "ps1":
		return PowerShell
	}

	return Unknown
}
