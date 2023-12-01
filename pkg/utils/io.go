package utils

import (
	"fmt"
	"os"
	"strings"
)

func ImportAsArray(path string) (input []string, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return []string{}, err
	}

	input = strings.Split(string(data), "\n")
	return
}

func ImportAsString(path string) (input string, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	input = string(data)
	return
}

func Output(data []byte, path string) {
	var err error
	_, err = os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			// Some unknown error
			fmt.Printf("Failed to stat file: %s", err.Error())
			os.Exit(1)
		}
	}

	f, err2 := os.Create(path)
	err = err2
	if err != nil {
		// Some error during creating
		fmt.Printf("Failed to create file: %s", err.Error())
		os.Exit(1)
	}

	_, err = f.Write(data)
	if err != nil {
		fmt.Printf("Failed to write file: %s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Written data to: %s\n", path)
	fmt.Printf("Output is %s\n", string(data))
	os.Exit(0)
}
