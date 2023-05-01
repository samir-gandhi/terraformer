package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// os.Args = []string{"", "import", "pingonedavinci", "--help"}
	os.Args = []string{"", "import", "pingonedavinci", "-r", "davinci_flow", "--path-pattern", "generated/davinci/", "-C=true", "-a=true"}
	main()
}
