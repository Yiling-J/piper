package main

import (
	"os"
)

func main() {
	configPath := os.Args[1]
	if configPath == "" {
		panic("empty config path")
	}
	dir, err := viperLoad(configPath)
	if err != nil {
		panic(err)
	}

	err = generate(dir)
	if err != nil {
		panic(err)
	}

}
