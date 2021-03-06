// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	fn    = "prelude.go"
	minFn = "prelude_min.go"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	bpkg, err := build.Import("github.com/gopherjs/gopherjs", ".", build.FindOnly)
	if err != nil {
		return fmt.Errorf("failed to locate path for github.com/gopherjs/gopherjs: %v", err)
	}

	preludeDir := filepath.Join(bpkg.Dir, "compiler", "prelude")

	files := []string{
		"prelude.js",
		"numeric.js",
		"types.js",
		"goroutines.js",
		"jsmapping.js",
	}

	prelude := new(bytes.Buffer)

	for _, f := range files {
		p := filepath.Join(preludeDir, f)
		out, err := ioutil.ReadFile(p)
		if err != nil {
			return fmt.Errorf("failed to read from %v: %v", p, err)
		}
		if _, err := prelude.Write(out); err != nil {
			return fmt.Errorf("failed to append prelude: %v", err)
		}
	}

	args := append([]string{
		filepath.Join(bpkg.Dir, "node_modules", ".bin", "uglifyjs"),
		"--config-file",
		filepath.Join(preludeDir, "uglifyjs_options.json"),
	}, files...)

	stderr := new(bytes.Buffer)

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = stderr

	minout, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to run %v: %v\n%s", strings.Join(args, " "), err, stderr.Bytes())
	}

	err = writeOutput("prelude.go", ""+
		"// Prelude is the GopherJS JavaScript interop layer.\n"+
		"const Prelude = %q", prelude.Bytes())

	if err != nil {
		return err
	}

	err = writeOutput("prelude_min.go", ""+
		"// Minified is an uglifyjs-minified version of Prelude.\n"+
		"const Minified = %q", minout)

	return err
}

func writeOutput(fn string, format string, byts []byte) error {
	content := fmt.Sprintf("// Code generated by genprelude; DO NOT EDIT.\n\n"+
		"package prelude\n\n"+format+"\n", byts)

	if err := ioutil.WriteFile(fn, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write to %v: %v", fn, err)
	}

	return nil
}
