package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to
	fileName := os.Args[1]
	var ioWriter io.Writer
	writer := bufio.NewWriter(ioWriter)
	inputFile, err := os.Open(fileName)
	check(err)
	fileContents := make([]byte, 100)

	bytesRead, err := inputFile.Read(fileContents)
	check(err)

	fmt.Printf("%d bytes: %s\n", bytesRead, string(fileContents[:bytesRead]))
	fmt.Fprintf(writer, "From Fprint: %d bytes: %s\n", bytesRead, string(fileContents[:bytesRead]))
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("scanner.Text():", scanner.Text())
	}
}
