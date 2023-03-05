package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

type State struct {
	SPtr string `json:"s"`
	OPtr string `json:"o"`
}

func main() {
	// Parse command line arguments
	sPtr := flag.String("s", "", "source file path")
	oPtr := flag.String("o", "", "output file path")
	flag.Parse()

	// Check if both -s and -o are provided
	if *sPtr == "" || *oPtr == "" {
		state, err := readState()
		if err != nil {
			log.Fatalln("Both -s and -o options are required")
		}
		sPtr = &state.SPtr
		oPtr = &state.OPtr
	} else {
		state := State{*sPtr, *oPtr}
		err := writeState(state)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Print the values of -s and -o
	fmt.Printf("-s flag value: %s\n", *sPtr)
	fmt.Printf("-o flag value: %s\n", *oPtr)
}

func readState() (State, error) {
	data, err := ioutil.ReadFile("state.json")
	if err != nil {
		return State{}, err
	}
	var state State
	err = json.Unmarshal(data, &state)
	if err != nil {
		return State{}, err
	}
	return state, nil
}

func writeState(state State) error {
	data, err := json.MarshalIndent(state, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("state.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}
