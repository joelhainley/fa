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
s - size
o - owner
g - group
w - world (other)

 */
const DEFAULT_OUTPUT_FORMAT = "%t.%s.%n"

var debugMode bool
var format string

/** TODO
* tokenizer for the format string
*
 */

func main() {
	// setup the flags
	flag.BoolVar(&debugMode, "debug", false, "turn on debug mode")
	flag.StringVar(&format, "format", DEFAULT_OUTPUT_FORMAT, "set the output format")
	flag.Parse()

	if debugMode {
		fmt.Printf("debug mode is enabled")
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fileInfo := dumpFileAttributes(scanner.Text(), format)
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

	if !fileInfo.IsDir(){
		output = strings.ReplaceAll(output, "%s", strconv.FormatInt(fileInfo.Size(), 10))
	} else {
		output = strings.ReplaceAll(output, "%s", "0")
	}

	return output
}
