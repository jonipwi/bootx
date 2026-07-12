package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/jonipwi/bootx/prototype/personal-companion/internal/engine"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/model"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/tui"
)

const maxJSONInput = 1 << 20

func main() {
	inputPath := flag.String("input", "", "strict JSON request path, or - for stdin; omit for TUI")
	compact := flag.Bool("compact", false, "emit compact JSON in backend mode")
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println(model.Version)
		return
	}
	decisionEngine, err := engine.New()
	if err != nil {
		fatal(err)
	}
	if *inputPath == "" {
		if err := tui.Run(os.Stdin, os.Stdout, decisionEngine); err != nil {
			fatal(err)
		}
		return
	}

	request, err := readRequest(*inputPath)
	if err != nil {
		fatal(err)
	}
	packet, err := decisionEngine.Process(request)
	if err != nil {
		fatal(err)
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(true)
	if !*compact {
		encoder.SetIndent("", "  ")
	}
	if err := encoder.Encode(packet); err != nil {
		fatal(err)
	}
}

func readRequest(path string) (model.Request, error) {
	var source io.Reader
	var file *os.File
	if path == "-" {
		source = os.Stdin
	} else {
		opened, err := os.Open(path)
		if err != nil {
			return model.Request{}, err
		}
		file = opened
		defer file.Close()
		source = file
	}
	decoder := json.NewDecoder(io.LimitReader(source, maxJSONInput))
	decoder.DisallowUnknownFields()
	var request model.Request
	if err := decoder.Decode(&request); err != nil {
		return model.Request{}, fmt.Errorf("decode strict request JSON: %w", err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return model.Request{}, fmt.Errorf("request must contain exactly one JSON object")
	}
	return request, nil
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "bootx-companion:", err)
	os.Exit(1)
}
