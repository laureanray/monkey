package vm

import (
	"monkey/code"
	"monkey/object"
)

const StackSize = 2048

type VM struct {
	constants    []object.Object
	instructions code.Instructions

	stack []object.Object

	sp int // Always point to the next value. Top of the stack is stack [sp-1]
}
