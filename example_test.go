package flagplus_test

import (
	"fmt"

	"github.com/scu/util/flagplus"
)

func Example() {
	// Create the FlagSet object
	var flags *flagplus.FlagSet
	flags = flagplus.NewFlagSet("util", "utility")

	// Set optional description
	flags.FlagSetDescription("Utility that does stuff")

	// Add flags of various types
	flags.AddBoolFlag("verbose", "v", "Print extra debugging information", false)
	flags.AddFlag("help", "h", "Help")
	flags.AddIntFlag("line", "l", "Start counting at `line_number`", 1)
	flags.AddStringFlag("output", "o", "Output `directory`", "/var/log/output")
	flags.AddFloatFlag("skew", "s", "Skew `percentage`", 2.33)

	// Parse the flags
	flags.Parse()

	// Get and use the flag options
	outOpt, err := flags.GetString("output")
	if err != nil {
		panic("Error retreiving output option")
	}
	fmt.Printf("output option is %q\n", outOpt)

	// Get and use command line arguments provided after flags
	args := flags.GetArgs()
	for _, arg := range args {
		fmt.Printf("arg is %q\n", arg)
	}

	/* Print program usage:
	fmt.Println(flags.Usage())

	Output:
	=====================================================================
	Usage:
	  util|utility [-h|l line_number|o directory|s percentage|v bool]
	Options:
	  -h, --help
	     Help
	  -l, --line line_number
	     Start counting at line_number (default=1)
	  -o, --output directory
	     Output directory (default=/var/log/output)
	  -s, --skew percentage
	     Skew percentage (default=2.33)
	  -v, --verbose bool
		 Print extra debugging information (default=false)
	=====================================================================
	*/
}

func ExampleFlagSet_AddStringFlag() {
	// Create the FlagSet object
	var flags *flagplus.FlagSet
	flags = flagplus.NewFlagSet("util", "utility")

	// Add the string flag
	flags.AddStringFlag("output", "o", "Output `directory`", "/var/log/output")

	// Parse
	flags.Parse()
}

func ExampleFlagSet_GetString() {
	// Create the FlagSet object
	var flags *flagplus.FlagSet
	flags = flagplus.NewFlagSet("util", "utility")

	// Add the string flag
	flags.AddStringFlag("output", "o", "Output `directory`", "/var/log/output")

	// Parse
	flags.Parse()

	// Get the string option
	outOpt, err := flags.GetString("output")
	if err != nil {
		panic("Error retreiving output option")
	}
	fmt.Printf("output option is %q\n", outOpt)
}

func ExampleFlagSet_Parse() {
	// Create the FlagSet object
	var flags *flagplus.FlagSet
	flags = flagplus.NewFlagSet("util", "utility")

	// Add flags here...

	// Parse
	flags.Parse()

	// Test flags here...
}
