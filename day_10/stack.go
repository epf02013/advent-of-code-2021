package main

type Stack struct {
	items []string
}

func (stack Stack) pop() (Stack, string, error) {
	return Stack{
		items: stack.items[0 : len(stack.items)-1],
	}, stack.items[len(stack.items)-1], nil
}

func (stack Stack) peek() string {
	return stack.items[len(stack.items)-1]
}

func (stack Stack) push(s string) Stack {
	return Stack{
		items: append(stack.items, s),
	}
}

func (stack Stack) size() int {
	return len(stack.items)
}
