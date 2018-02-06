package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	Process(reader, writer)
}

// Process data from reader and write to writer
func Process(reader *bufio.Reader, writer *bufio.Writer) {
	defer writer.Flush()
	fmt.Fprint(writer, "Input text:\n")
	writer.Flush()

	text, _ := reader.ReadString('\n')
	fmt.Fprint(writer, text)
}
