package main

import (
	"runtime"

	"github.com/yuankunzhang/paradox/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	commands.Execute()
}