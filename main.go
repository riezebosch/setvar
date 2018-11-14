package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	isOutput := flag.Bool("isOutput", false, "output variable")
	isSecret := flag.Bool("isSecret", false, "secret variable")
	name := flag.String("name", "", "variable name")
	flag.Parse()

	if *name == "" {
		panic(fmt.Errorf("specify value for '-name'"))
	}

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		panic(fmt.Errorf("pipe output into this tool"))
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		fmt.Printf("##vso[task.setvariable variable=%v;isOutput=%t;isSecret=%t]%s", *name, *isOutput, *isSecret, input)
	}
}
