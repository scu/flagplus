# flagplus Go Library

flagplus is a wrapper and extension to the Go core flag package.
The flagplus.FlagSet type emphasizes simple declarative flag definitions,
easily divisible by main() or init()

The Usage output is readable and informative, including a summary option list.

## Getting Started
### Installing the package
```
go get github.com/scu/flagplus
```

## Usage

```go
import "github.com/scu/flagplus"
```

## Basic Example
See godocs for additional code examples.
```go
package main

import (
	"fmt"

	"github.com/scu/util/flagplus"
)

var flags *flagplus.FlagSet

func main() {
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

	// Print program usage
	fmt.Println(flags.Usage())
}

func init() {
	// Create the FlagSet object
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
}
```
### Usage Output:
```
Utility that does stuff
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
```

## Running the tests
```
go test -v
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Versioning

[SemVer](http://semver.org/) is used for versioning. For the versions available, see the [tags on this repository](https://github.com/scu/flagplus/tags). 

## Authors

* **Scott Underwood** - *Initial work* - [Scott Underwood](https://github.com/scu)

## License
[MIT](https://choosealicense.com/licenses/mit/)