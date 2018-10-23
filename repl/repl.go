/*
package repl implements "Read Eval Print Loop".
The REPL reads input, sends it to the interpreter for evaluation,
prints the result/output of the interpreter and start again
*/
package repl

import (
	"io"
	"bufio"
	"fmt"
	"myinterpreter/lexer"
	"myinterpreter/token"
)

const PROMPT = ">>>>>>>>>>>>"

func Start(in io.Reader, out io.Writer){
	sc := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		isScanned := sc.Scan()

		if !isScanned {
			return
		}

		line := sc.Text()
		l := lexer.New(line)

		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken(){
			fmt.Printf("%+v\n", t)
		}
	}
}