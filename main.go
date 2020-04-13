package main

import "github.com/pantheonproject/gaia/internal/cmd"

var startFunc = cmd.Execute

func main() {
	_ = startFunc()
}
