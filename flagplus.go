// Copyright 2020 Scott Underwood.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package flagplus provides a wrapper for the standard flags package.
// The FlagSet type is more declarative, more easily encapsulated, and
// automatically provides for long and short options.
// In addition, the Usage output is much more readable and informative.
package flagplus

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// FlagType holds the type of the flag
type FlagType int

const (
	// BASE is a binary flag - either set or unset
	BASE FlagType = iota
	// BOOL is a boolean flag, true or false
	BOOL
	// INT is an integer flag
	INT
	// FLOAT is a float flag
	FLOAT
	// STRING is a string flag
	STRING
)

// Flag represents the state of a flag
type Flag struct {
	key          string      // Key to the map index, also the long name
	shortName    string      // Short name as it appears on command line
	flagType     FlagType    // The type of the flag
	value        interface{} // The value as set
	defaultValue interface{} // Holds the dynamic value of the flag (for usage)
	usage        string      // Usage statement
}

// FlagSet represents a set of defined flags
type FlagSet struct {
	isParsed    bool             // Has the FlagSet been parsed using the Parse() func?
	coreFlagSet flag.FlagSet     // Core FlagSet
	flag        map[string]*Flag // Flags in the FlagSet
	name        string           // Optional name of the flag set
	description string           // Optional description of command line
}

// String implements fmt.string interface for Flag
func (f *Flag) String() string {
	var s, typeStr, defStr string
	switch f.flagType {
	case BASE:
		typeStr = "BASE"
		defStr = "n/a"
	case BOOL:
		typeStr = "BOOL"
		defStr = strconv.FormatBool(f.defaultValue.(bool))
	case INT:
		typeStr = "INT"
		defStr = fmt.Sprintf("%d", f.defaultValue.(int64))
	case FLOAT:
		typeStr = "FLOAT"
		defStr = fmt.Sprintf("%f", f.defaultValue.(float64))
	case STRING:
		typeStr = "STRING"
		defStr = f.defaultValue.(string)
	}
	s += fmt.Sprintf("TYPE=%s shortName=%q usage=%q default=%q\n",
		typeStr, f.shortName, f.usage, defStr)

	return s
}

// String implements the fmt.string interface for FlagSet
func (fs *FlagSet) String() string {
	var s string
	s += fmt.Sprintf("Flagset name=%q description=%q\n{\n", fs.name, fs.description)
	for k, v := range fs.flag {
		s += fmt.Sprintf("\tflag[%q]: %v", k, v)
	}
	s += "}\n"

	return s
}

// AddStringFlag adds a string flag to a FlagSet
func (fs *FlagSet) AddStringFlag(key, shortName, usage string, defaultValue string) {
	fs.addFlag(
		STRING,
		key,
		shortName,
		usage,
		defaultValue,
	)
}

// AddFloatFlag adds a float flag to a FlagSet
func (fs *FlagSet) AddFloatFlag(key, shortName, usage string, defaultValue float64) {
	fs.addFlag(
		FLOAT,
		key,
		shortName,
		usage,
		defaultValue,
	)
}

// AddFlag adds a base flag to a FlagSet
func (fs *FlagSet) AddFlag(key, shortName, usage string) {
	fs.addFlag(
		BASE,
		key,
		shortName,
		usage,
		nil,
	)
}

// AddIntFlag adds an integer flag to a FlagSet
func (fs *FlagSet) AddIntFlag(key, shortName, usage string, defaultValue int64) {
	fs.addFlag(
		INT,
		key,
		shortName,
		usage,
		defaultValue,
	)
}

// AddBoolFlag adds a boolean flag to a FlagSet
func (fs *FlagSet) AddBoolFlag(key, shortName, usage string, defaultValue bool) {
	fs.addFlag(
		BOOL,
		key,
		shortName,
		usage,
		defaultValue,
	)
}

// addFlag adds a new flag to a FlagSet
func (fs *FlagSet) addFlag(
	flagType FlagType,
	key, shortName, usage string,
	defaultValue interface{}) {

	newFlag := new(Flag)
	newFlag.key = key
	newFlag.flagType = flagType
	newFlag.shortName = shortName
	newFlag.defaultValue = defaultValue
	newFlag.usage = usage

	// Initialize values in core.flag
	switch flagType {
	case BASE:
		newFlag.value = fs.coreFlagSet.Bool(key, false, usage)
		fs.coreFlagSet.BoolVar(newFlag.value.(*bool), shortName, false, usage)
	case BOOL:
		newFlag.value = fs.coreFlagSet.Bool(key, defaultValue.(bool), usage)
		fs.coreFlagSet.BoolVar(newFlag.value.(*bool), shortName, defaultValue.(bool), usage)
	case INT:
		newFlag.value = fs.coreFlagSet.Int64(key, defaultValue.(int64), usage)
		fs.coreFlagSet.Int64Var(newFlag.value.(*int64), shortName, defaultValue.(int64), usage)
	case FLOAT:
		newFlag.value = fs.coreFlagSet.Float64(key, defaultValue.(float64), usage)
		fs.coreFlagSet.Float64Var(newFlag.value.(*float64), shortName, defaultValue.(float64), usage)
	case STRING:
		newFlag.value = fs.coreFlagSet.String(key, defaultValue.(string), usage)
		fs.coreFlagSet.StringVar(newFlag.value.(*string), shortName, defaultValue.(string), usage)
	}

	// Assign flag to FlagSet map
	fs.flag[key] = newFlag
}

// GetArgs returns the arguments after flags
func (fs *FlagSet) GetArgs() []string {
	return fs.coreFlagSet.Args()
}

// Get returns a basic flag value
func (fs *FlagSet) Get(key string) (bool, error) {
	if err := fs.flagCheck(key, BASE); err != nil {
		return false, err
	}

	return *fs.flag[key].value.(*bool), nil
}

// GetBool returns a boolean flag value
func (fs *FlagSet) GetBool(key string) (bool, error) {
	if err := fs.flagCheck(key, BOOL); err != nil {
		return false, err
	}

	return *fs.flag[key].value.(*bool), nil
}

// GetInt returns an integer flag value
func (fs *FlagSet) GetInt(key string) (int64, error) {
	if err := fs.flagCheck(key, INT); err != nil {
		return 0, err
	}

	return *fs.flag[key].value.(*int64), nil
}

// GetFloat returns a float flag value
func (fs *FlagSet) GetFloat(key string) (float64, error) {
	if err := fs.flagCheck(key, FLOAT); err != nil {
		return 0.00, err
	}

	return *fs.flag[key].value.(*float64), nil
}

// GetString returns a string flag value
func (fs *FlagSet) GetString(key string) (string, error) {
	if err := fs.flagCheck(key, STRING); err != nil {
		return "", err
	}

	return *fs.flag[key].value.(*string), nil
}

// flagCheck inspects the flag map by key for presence, type and
// if being requested prior to parse
func (fs *FlagSet) flagCheck(key string, flagType FlagType) error {
	// Has the flag set been parsed?
	if !fs.isParsed {
		return fmt.Errorf("FlagSet %q has not been parsed", fs.name)
	}

	// Check if key exists
	if _, ok := fs.flag[key]; !ok {
		return fmt.Errorf("%q: flag does not exist", key)
	}

	// Check if the flag type matches expectation
	if fs.flag[key].flagType != flagType {
		return fmt.Errorf("%q: incorrect flag type", key)
	}

	return nil
}

// FlagSetDescription sets the optional description of the FlagSet
func (fs *FlagSet) FlagSetDescription(description string) {
	fs.description = description
}

// SimulateArg allows the test suite to simulate command-line arguments
func (fs *FlagSet) SimulateArg(name string, value string) error {
	return fs.coreFlagSet.Set(name, value)
}

// Parse parses flag definitions
func (fs *FlagSet) Parse(args ...string) error {
	if len(args) > 0 {
		os.Args = args
	}
	err := fs.coreFlagSet.Parse(os.Args[1:])
	if err != nil {
		return fmt.Errorf("Could not parse FlagSet %q", fs.name)
	}
	fs.isParsed = true
	return nil
}

func unquoteUsage(flag *Flag) (name string, usage string) {
	usage = flag.usage

	for i := 0; i < len(usage); i++ {
		if usage[i] == '`' {
			for j := i + 1; j < len(usage); j++ {
				if usage[j] == '`' {
					name = usage[i+1 : j]
					usage = usage[:i] + name + usage[j+1:]
					return name, usage
				}
			}
			break // Only one back quote; use type name.
		}
	}

	// If not explicit in usage `backquotes`, use type
	switch flag.flagType {
	case BOOL:
		name = "bool"
	case INT:
		name = "int"
	case FLOAT:
		name = "float"
	case STRING:
		name = "string"
	}
	return
}

// sortFlags returns the flags as a slice in lexicographical sorted order.
func sortFlags(flags map[string]*Flag) []*Flag {
	result := make([]*Flag, len(flags))
	i := 0
	for _, f := range flags {
		result[i] = f
		i++
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].key < result[j].key
	})
	return result
}

// flagDefaultValue creates output if there is a default
// value on all flag types except for BASE
func flagDefaultValue(flag *Flag) string {
	s := ""

	switch flag.flagType {
	case BOOL:
		s = fmt.Sprintf(" (default=%v)", flag.defaultValue.(bool))
	case INT:
		s = fmt.Sprintf(" (default=%v)", flag.defaultValue.(int64))
	case FLOAT:
		s = fmt.Sprintf(" (default=%v)", flag.defaultValue.(float64))
	case STRING:
		if flag.defaultValue.(string) != "" {
			s = fmt.Sprintf(" (default=%v)", flag.defaultValue.(string))
		}
	}

	return s
}

// flagUsage builds the usage string for each command line option.
func flagUsage(flag *Flag) string {
	// Get optional unquote usage
	name, usage := unquoteUsage(flag)

	s := fmt.Sprintf("\n  -%s, --%s %s\n     %s",
		flag.shortName, flag.key, name, usage)

	if flag.defaultValue != nil {
		s += flagDefaultValue(flag)
	}

	return s
}

// Usage prints program usage information
func (fs *FlagSet) Usage() string {
	var s string

	// Optional description
	if fs.description != "" {
		s += fmt.Sprintf("%s\n", fs.description)
	}

	// Summary "Usage: ..." statement
	s += fmt.Sprintf("Usage:\n  %s", fs.name)
	if len(fs.flag) > 0 {
		s += " [-"
		for _, f := range sortFlags(fs.flag) {
			s += f.shortName
			// Get optional unquote usage
			if n, _ := unquoteUsage(f); n != "" {
				s += fmt.Sprintf(" %s", n)
			}

			s += "|"
		}
		s = strings.TrimRight(s, "|")
		s += "]"
	}

	// Full option description
	if len(fs.flag) > 0 {
		s += "\nOptions:"
		for _, f := range sortFlags(fs.flag) {
			s += flagUsage(f)
		}
	}

	return s
}

// NewFlagSet returns a new, empty flag set
func NewFlagSet(name ...string) *FlagSet {
	// Allow multiple names (or no name) to be the set name
	f := &FlagSet{name: strings.Join(name, "|")}

	// Create the flag map, preallocate space for 64 flags
	f.flag = make(map[string]*Flag, 64)

	return f
}
