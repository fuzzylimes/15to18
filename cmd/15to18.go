package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fuzzylimes/sf15to18"
)

var source string

func init() {
	flag.StringVar(&source, "file", "", "Specifies a file to be converted")
	flag.Parse()
}

func main() {
	args := sf15to18.HandleArgs(source, flag.Args())

	if len(args) < 1 {
		fmt.Println("Must provide SF Id or file")
		os.Exit(1)
	}

	input := sf15to18.GetValidIds(args)

	for _, i := range input {
		output := sf15to18.Convert(i)
		fmt.Printf("%s -> %s\n", i, output)
	}
}
