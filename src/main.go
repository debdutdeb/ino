package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	// topics := make(map[string][]string)
	topics := &[]string{}

	path := flag.String("ino-path", "", "Path to ino.yml file")
	output := flag.String("output", "/tmp/output.md", "Output file location")

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

	outFile, err := os.OpenFile(*output, 0777, fs.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	for _, topic := range *topics {
		outFile.WriteString(fmt.Sprintf("- %s\n", topic))
	}

}
