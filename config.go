package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Config represents the structure of our configuration JSON file
type Config struct {
	APIKey   string `json:"api_key"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func main() {
	// The path to your JSON configuration file
	configFilePath := "./.config.json"

	// Read the file
	fileContents, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Error reading the config file: %s", err)
	}

	// Unmarshal the JSON into a Config struct
	var config Config
	if err := json.Unmarshal(fileContents, &config); err != nil {
		log.Fatalf("Error decoding config file: %s", err)
	}

	// Now you can use the config in your program
	fmt.Printf("API Key: %s\n", config.APIKey)
	fmt.Printf("User: %s\n", config.User)
	fmt.Printf("Password: %s\n", config.Password)
}
