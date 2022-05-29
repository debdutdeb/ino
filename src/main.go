package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	path, output string
)

func init() {
	flag.StringVar(&path, "ino-path", "", "Path to ino.yml file")
	flag.StringVar(&output, "output", "/tmp/output.md", "Output file location")
}

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	// topics := make(map[string][]string)
	topics := &[]string{}

	if path == "" {
		panic("ino.yml path is mandatory")
	}

	if output == "" {
		fmt.Fprintln(os.Stderr, "output file location is mandatory")
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read file "+path+err.Error())
		os.Exit(1)
	}

	err = yaml.Unmarshal(file, topics)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read yaml content; "+path+err.Error())
		os.Exit(1)
	}

	outFile, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer outFile.Close()

	for _, topic := range *topics {
		outFile.WriteString(fmt.Sprintf("- %s\n", topic))
	}

}
