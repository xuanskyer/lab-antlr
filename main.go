package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"lab-antlr/parser"
	"strconv"
)

func main() {
	str := "( 1 + 2 ) * 3 + 4"
	result := calc(str)
	fmt.Printf("%s = %d\n", str, result)
	str = "1 + 2 * 3 + 4"
	result = calc(str)
	fmt.Printf("%s = %d\n", str, result)
	str = "3 * (( 1 + 2 ) * 3 + 4)"
	result = calc(str)
	fmt.Printf("%s = %d\n", str, result)
	str = "(3 + ( 1 + 2 ) * 3)/2 + 4"
	result = calc(str)
	fmt.Printf("%s = %d\n", str, result)
	str = "(3 + ( 1 + 2 ) * 3)/2 + 4"
	result = calc(str)
	fmt.Printf("%s = %d\n", str, result)
}

// https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/
func calc(input string) int {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	return listener.pop()
}

type calcListener struct {
	*parser.BaseCalcListener

	stack []int
}

func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()
	fmt.Printf("ExitMulDiv: left: %d, right: %d\n", left, right)
	switch c.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left * right)
	case parser.CalcParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {

	right, left := l.pop(), l.pop()
	fmt.Printf("ExitAddSub, left: %d, right: %d\n", left, right)

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(left + right)
	case parser.CalcParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitNumber(c *parser.NumberContext) {

	i, err := strconv.Atoi(c.GetText())
	fmt.Printf("ExitNumber: %d\n", i)
	if err != nil {
		panic(err.Error())
	}

	l.push(i)
}
