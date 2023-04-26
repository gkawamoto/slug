package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gosimple/slug"
)

func main() {
	args := os.Args[1:]

	var reader io.Reader

	if len(args) > 0 {
		reader = bytes.NewBufferString(strings.Join(args, "\n"))
	} else {
		reader = os.Stdin
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Fprintln(os.Stdout, slug.Make(scanner.Text()))
	}
}
