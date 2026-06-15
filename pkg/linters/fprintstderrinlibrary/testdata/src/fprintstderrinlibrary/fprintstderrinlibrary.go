package fprintstderrinlibrary

import (
	"fmt"
	"os"
)

func bad() {
	fmt.Fprintf(os.Stderr, "error: %s", "something") // want `fmt\.Fprintf.*os\.Stderr.*called in library package`
	fmt.Fprintln(os.Stderr, "something")              // want `fmt\.Fprintln.*os\.Stderr.*called in library package`
}

func good() {
	fmt.Fprintf(os.Stdout, "output: %s", "val")
}

func suppressed() {
	//nolint:fprintstderrinlibrary
	fmt.Fprintf(os.Stderr, "suppressed previous line")
	fmt.Fprint(os.Stderr, "suppressed same line") //nolint:fprintstderrinlibrary
}
