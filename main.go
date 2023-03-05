package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type stateData struct {
	SPtr string `json:"s_ptr"`
	OPtr string `json:"o_ptr"`
}

func loadStateFromJSON(path string) (*stateData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var state stateData
	err = json.NewDecoder(file).Decode(&state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func main() {
	// Check if state file exists
	if _, err := os.Stat("state.json"); err == nil {
		// State file exists, load previous state
		state, err := loadStateFromJSON("state.json")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Previous state loaded:")
		fmt.Printf("-s=%s\n", state.SPtr)
		fmt.Printf("-o=%s\n", state.OPtr)

	}
