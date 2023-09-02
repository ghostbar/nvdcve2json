package main

import "strings"

// StringFlags represents a slice of strings that satisfies the interface of
// flag.Value. This allows for a flag to be passed several times and each time
// it'll create a single-entry string that can then be retrieved atomically.
type StringFlags []string

// String returns a concatenated output of the individual strings passed as
// input, separated by a space.
func (i StringFlags) String() string {
	out := ""
	for _, v := range i {
		out += v + " "
	}
	return strings.TrimSuffix(out, " ")
}

// Set appends the value as a single-entry for a slice flag. Error always
// returns nil.
func (i *StringFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// Slice returns the raw slice of strings with all the values set for the flag.
func (i StringFlags) Slice() []string {
	return i
}
