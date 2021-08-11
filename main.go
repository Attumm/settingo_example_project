package main

import (
	"fmt"

	"github.com/Attumm/settingo/settingo"
)

func main() {
	// Setting strings
	settingo.Set("FOOBAR", "default_value_for_foo_bar", "explain why something is foobar")

	// Setting integers
	settingo.SetInt("answer_to_the_universe", 42, "explain why 42 is the answer to the universe")

	// Setting Booleans
	settingo.SetBool("hamlet", true, "to be or not to be?")

	// Setting Maps
	defaultMap := make(map[string][]string)
	defaultMap["foo"] = []string{"bar"}
	settingo.SetMap("FOO", defaultMap, "help text")

	// Run parse after the setting set.
	settingo.Parse()

	foobar := settingo.Get("FOOBAR")

	fmt.Println("foobar =", foobar)
	fmt.Println("answer to the universe =", settingo.GetInt("answer_to_the_universe"))
	fmt.Println("to be or not to be? =", settingo.GetBool("hamlet"))

	fmt.Println("the result of the map", settingo.GetMap("FOO"))
}

//settingo.SetBool("starred", "y", "values (y, yes, true, t, '') have the property of truthiness")

// run parse to parse input from:
// priority order
// 1. command line input
// 2. environment
// 3. default values

// to access the values they are stored in SettingS
// settingo is a singleton and is accessable from each file where the import is done.
// according to best practices it's better to take the settings vars out in the main.

// yup that's it

// to compile
// go mod init
// go get
// go build

// to run
//$ ./example --help
//Usage of ./example:
//  -FOOBAR string
//    	explain why something is foobar (default "default_value_for_foo_bar")
//  -answer_to_the_universe int
//    	explain why 42 is the answer to the universe (default 42)

//$ ./example -answer_to_the_universe 44
//foobar = default_value_for_foo_bar
//answer to the universe = 44
