package main

import (
	"example.com/library"
)

func main() {
	a := library.App{}
	a.Port = ":9003"
	a.Initialize()
	a.Run()
}