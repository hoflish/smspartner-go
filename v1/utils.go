package smspartner

import "strings"

// FirstNonEmptyString iterates through its
// arguments trying to find the first string
// that is not blank or consists entirely  of spaces.
func FirstNonEmptyString(args ...string) string {
	for _, arg := range args {
		if arg == "" {
			continue
		}
		if strings.TrimSpace(arg) != "" {
			return arg
		}
	}
	return ""
}
