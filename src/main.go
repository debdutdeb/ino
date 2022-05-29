package main

import (
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func main() {
	// topics := make(map[string][]string)
	topics := &[]string{}

	path := flag.String("ino-path", "", "Path to ino.yml file")

	flag.Parse()

	if *path == "" {
		panic("ino.yml path is mandatory")
	}

	file, err := ioutil.ReadFile(*path)
	if err != nil {
		panic("failed to read file " + *path + err.Error())
	}

	err = yaml.Unmarshal(file, topics)
	if err != nil {
		panic("failed to read yaml content; " + *path + err.Error())
	}
}
