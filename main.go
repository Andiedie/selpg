package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	start    = flag.Int("s", -1, "start page number")
	end      = flag.Int("e", -1, "end page number")
	line     = flag.Int("l", -1, "line number per page")
	formFeed = flag.Bool("f", false, "use '\\f' to paging instead of line")
	file     *os.File
)

func main() {
	parse()
	reader := getReader()
	print(reader)
	clean()
}

func report(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}

func parse() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: selpg [flags] path\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()
	var err error
	if *start == -1 {
		err = errors.New("flag 's' required")
	} else if *end == -1 {
		err = errors.New("flag 'e' required")
	} else if *start > *end {
		err = errors.New("The start page number can not be greater than the end page number")
	} else if *line != -1 && *formFeed {
		err = errors.New("flag 'l' and 'f' are mutually exclusive")
	}
	report(err)
	if *line == -1 && !*formFeed {
		*line = 72
	}
}

func getReader() *bufio.Reader {
	if paths := flag.Args(); len(paths) > 0 {
		var err error
		file, err = os.Open(paths[0])
		report(err)
		return bufio.NewReader(file)
	}
	return bufio.NewReader(os.Stdin)
}

func print(reader *bufio.Reader) {
	var (
		delim byte
		err   error
		bytes []byte
	)
	if *formFeed {
		delim = '\f'
		for i := 0; i < *end && err == nil; i++ {
			bytes, err = reader.ReadBytes(delim)
			if i >= (*start - 1) {
				fmt.Printf("%s\n", bytes[:len(bytes)-1])
			}
		}
	} else {
		delim = '\n'
		startLine := (*start - 1) * *line
		endLine := *end * *line
		for i := 0; i < endLine && err == nil; i++ {
			bytes, err = reader.ReadBytes(delim)
			if i >= startLine {
				fmt.Printf("%s", bytes)
			}
		}
	}
	if err != io.EOF {
		report(err)
	}
}

func clean() {
	file.Close()
}
