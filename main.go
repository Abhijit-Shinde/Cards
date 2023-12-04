package main

import (
	// "fmt"
	"io"
	"os"
)

func main() {
	arguments := os.Args

	file,_ := os.Open(arguments[1])

	io.Copy(os.Stdout,file)

}
