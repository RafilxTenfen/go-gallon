package main

import (
	"os"
	"runtime"

	"github.com/RafilxTenfen/go-gallon/cli/gallon"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	os.Exit(gallon.Main(os.Args[1:]))
}
