package main

import (
	"flag"
	"github.com/rfaulhaber/rexpr/pkg/expr"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	stdout = log.New(os.Stdout, "", 0)
	stderr = log.New(os.Stderr, "rexpr: ", 0)
)

func main() {
	convertFlag := flag.Bool("c", false, "set if converting expression to infix")

	flag.Parse()

	var exprStr string

	if isFromPipe() {
		str, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			stderr.Fatalln("could not read expression: ", err.Error())
		}

		exprStr = trimQuotes(string(str))
	} else {
		if len(flag.Args()) > 1 {
			exprStr = strings.Join(flag.Args(), " ")
		} else {
			exprStr = flag.Args()[0]
		}
	}

	exprNode, err := expr.ParseString(strings.TrimSuffix(exprStr, "\n"))

	if err != nil {
		stderr.Fatalln("could not parse expression: ", err.Error())
	}

	if *convertFlag {
		stdout.Println(exprNode.String())
	} else {

		result, err := exprNode.Evaluate()

		if err != nil {
			stderr.Fatalln("could not evaluate expression: ", err.Error())
		}

		stdout.Println(result)
	}
}

func isFromPipe() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Mode() & os.ModeCharDevice == 0
}

func trimQuotes(s string) string {
	return strings.Replace(s, "\"", "", -1)
}

