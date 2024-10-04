package main

import (
	"flag"
	"fmt"
)

type CmdFlags struct {
	Add bool
	Sub bool
	Mul bool
	Div bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.BoolVar(&cf.Add, "add", false, "Add two variables")
	flag.BoolVar(&cf.Sub, "sub", false, "Subtract two variables")
	flag.BoolVar(&cf.Mul, "mul", false, "Multiply two variables")
	flag.BoolVar(&cf.Div, "div", false, "Divide two variables")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(operator *Variables) {
	switch {
	case cf.Add:
		operator.add()
	case cf.Sub:
		operator.subtract()
	case cf.Mul:
		operator.multiply()
	case cf.Div:
		operator.division()
	default:
		fmt.Println("invalid command")
	}
}
