package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)


/*
file attributes values
t - file type
n - file name
s - [x] - size [x] human readable etc..
o - owner
g - group
w - world (other)

 */
const DEFAULT_OUTPUT_TEMPLATE = "%t.%s.%n"

//var offsetPrefix string
//var itemGlyph string
//var lastItemGlyph string
//var lineContGlyph string
var debugMode bool
//var format string

func main() {
	// setup the flags
	flag.BoolVar(&debugMode, "debug", false, "turn on debug mode")
	flag.Parse()

	if debugMode {
		fmt.Printf("debug mode is enabled")
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fileInfo := dumpFileAttributes(scanner.Text(), DEFAULT_OUTPUT_TEMPLATE)
		fmt.Printf("> %s\n", fileInfo)
	}
}

func dumpFileAttributes(filepath string, format string) string {
	fileInfo, err := os.Lstat(filepath)

	if err != nil {
		fmt.Printf("\nerror!!\n")
		panic(err)
	}

	fileType := "f"
	if fileInfo.IsDir() {
		fileType = "d"
	}

	output := strings.ReplaceAll(format, "%t", fileType)
	output = strings.ReplaceAll(output, "%n", fileInfo.Name())

	if fileInfo.IsDir(){
		output = strings.ReplaceAll(output, "%s", strconv.FormatInt(fileInfo.Size(), 10))
	}

	return output
}
