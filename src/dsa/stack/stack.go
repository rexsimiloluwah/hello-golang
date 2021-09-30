package stack

import "fmt"

type Stack struct {
	data []interface{}
	size int
}

func (stack *Stack) IsEmpty() bool {
	return (*stack).size == 0
}

func (stack *Stack) Size() int {
	return (*stack).size
}

func (stack *Stack) Push(el interface{}) {
	(*stack).data = append((*stack).data, el)
	(*stack).size++
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		panic("Stack is empty")
	}
	len := stack.Size()
	poppedEl := (*stack).data[len-1]
	(*stack).data = (*stack).data[:(len - 1)]
	(*stack).size--
	return poppedEl
}

func (stack *Stack) Peek() interface{} {
	if stack.Size() == 0 {
		panic("Stack is empty")
	}
	len := (*stack).Size()
	return (*stack).data[len-1]
}

func (stack *Stack) Print() {
	fmt.Println((*stack).data)
}
