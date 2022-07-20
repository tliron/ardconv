package main

import (
	"github.com/tliron/kutil/util"

	_ "github.com/tliron/kutil/logging/simple"
)

func main() {
	util.ExitOnSIGTERM()
	Execute()
	util.Exit(0)
}
