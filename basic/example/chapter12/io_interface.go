package chapter12

import (
	"bufio"
	"fmt"
	"os"
)

func Io_interface() {
	fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuufered")
	// buffered: os.Stdout implements io.Writer
	buf := bufio.NewWriter(os.Stdout)

	fmt.Fprintf(buf, "%s\n", "hello world - buffered")
	buf.Flush()
}
