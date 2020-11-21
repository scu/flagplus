package flagplus

import (
	"fmt"
	"strconv"
	"testing"
)

// initializeFlagSet creates a new FlagSet for test suite
func initalizeFlagSet() *FlagSet {
	var flags *FlagSet
	flags = NewFlagSet("util")

	return flags
}

func TestFlagSet_AddStringFlag_DefaultValue(t *testing.T) {
	expect := "/var/log/output"

	flags := initalizeFlagSet()
	flags.AddStringFlag("output", "-o", "Output `directory`", expect)

	// Simulate arguments
	flags.Parse("util")

	got, err := flags.GetString("output")
	if err != nil {
		t.Fatalf("Could not get flag output: %v", err)
	}

	if got != expect {
		t.Errorf("Expected %q, got %q", expect, got)
	}
}

func TestFlagSet_AddStringFlag_ArgValue(t *testing.T) {
	expect := "foo"

	flags := initalizeFlagSet()
	flags.AddStringFlag("output", "-o", "Output `directory`", "/var/log/out")

	// Simulate arguments
	flags.SimulateArg("output", expect)
	flags.Parse("util")

	got, err := flags.GetString("output")
	if err != nil {
		t.Fatalf("Could not get flag output: %v", err)
	}

	if got != expect {
		t.Errorf("Expected %q, got %q", expect, got)
	}
}

func getAndSetStringFlag(t *testing.T) {
	TestFlagSet_AddStringFlag_DefaultValue(t)
	TestFlagSet_AddStringFlag_ArgValue(t)
}

func TestFlagSet_AddStringFlag(t *testing.T) {
	getAndSetStringFlag(t)
}

func getAndSetFloatFlag(t *testing.T) {
	expect := 3.55

	flags := initalizeFlagSet()
	flags.AddFloatFlag("skew", "-s", "Percentage to skew", 2.24)

	// Simulate arguments
	flags.SimulateArg("skew", strconv.FormatFloat(expect, 'f', 6, 64))
	flags.Parse("util")

	got, err := flags.GetFloat("skew")
	if err != nil {
		t.Fatalf("Could not get flag output: %v", err)
	}

	if got != expect {
		t.Errorf("Expected %f, got %f", expect, got)
	}
}

func TestFlagSet_AddFloatFlag(t *testing.T) {
	getAndSetFloatFlag(t)
}

func getAndSetAddFlag(t *testing.T) {
	expect := true

	flags := initalizeFlagSet()
	flags.AddFlag("flag", "-f", "Only true if set")

	// Simulate arguments
	flags.SimulateArg("flag", "true")
	flags.Parse("util")

	got, err := flags.Get("flag")
	if err != nil {
		t.Fatalf("Could not get flag output: %v", err)
	}

	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}
}

func TestFlagSet_AddFlag(t *testing.T) {
	getAndSetAddFlag(t)
}

func getAndSetIntFlag(t *testing.T) {
	var expect int64 = 4

	flags := initalizeFlagSet()
	flags.AddIntFlag("line", "-l", "Line Number", 1)

	// Simulate arguments
	flags.SimulateArg("line", fmt.Sprintf("%v", expect))
	flags.Parse("util")

	got, err := flags.GetInt("line")
	if err != nil {
		t.Fatalf("Could not get flag output: %v", err)
	}

	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}
}

func TestFlagSet_AddIntFlag(t *testing.T) {
	getAndSetIntFlag(t)
}

func getAndSetBoolFlag(t *testing.T) {
	expect := true

	flags := initalizeFlagSet()
	flags.AddBoolFlag("boolflag", "-b", "Only true if set", false)

	// Simulate arguments
	flags.SimulateArg("boolflag", fmt.Sprintf("%v", expect))
	flags.Parse("util")

	got, err := flags.GetBool("boolflag")
	if err != nil {
		t.Fatalf("Could not get flag output: %v", err)
	}

	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}
}

func TestFlagSet_AddBoolFlag(t *testing.T) {
	getAndSetBoolFlag(t)
}

func TestFlagSet_GetArgs(t *testing.T) {
	expect := []string{"one", "two", "three"}

	flags := initalizeFlagSet()
	flags.AddBoolFlag("boolflag", "-b", "Only true if set", false)

	// Simulate arguments
	flags.Parse("util", "one", "two", "three")
	got := flags.GetArgs()

	for i, item := range got {
		if item != expect[i] {
			t.Errorf("Arguments %q does not match %q", item, expect[i])
		}
	}
}

func TestFlagSet_Get(t *testing.T) {
	getAndSetBoolFlag(t)
}

func TestFlagSet_GetBool(t *testing.T) {
	getAndSetBoolFlag(t)
}

func TestFlagSet_GetInt(t *testing.T) {
	getAndSetIntFlag(t)
}

func TestFlagSet_GetFloat(t *testing.T) {
	getAndSetFloatFlag(t)
}

func TestFlagSet_GetString(t *testing.T) {
	getAndSetStringFlag(t)
}

func TestNewFlagSet(t *testing.T) {
	var got *FlagSet
	var expect string

	// Empty argument
	got = NewFlagSet()
	if got.name != "" {
		t.Errorf("Expected empty name but got %q", got.name)
	}

	// Single name argument
	got = NewFlagSet("package")
	expect = "package"
	if got.name != expect {
		t.Errorf("Expected name %q but got %q", expect, got.name)
	}

	// Multiple name arguments
	got = NewFlagSet("package", "pack", "pkg")
	expect = "package|pack|pkg"
	if got.name != expect {
		t.Errorf("Expected name %q but got %q", expect, got.name)
	}

}

func TestFlagSetDescription(t *testing.T) {
	got := NewFlagSet()
	if got == nil {
		t.Fatal("Could not create a new FlagSet")
	}

	got.FlagSetDescription("Does some stuff")
	expect := "Does some stuff"
	if got.description != expect {
		t.Errorf("Expected description %q but got %q", expect, got.description)
	}
}
