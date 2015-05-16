package main

import(
	"fmt"
	"errors"
)

func main() {
	var sm StackMachine
	sm.operators = map[rune]func()error{
		'+': sm.add,
		'*': sm.multiply,
	}
	fmt.Println(sm.calculate("13+62*7+*"))
}

type StackMachine struct {
	numStack []int
	operators   map[rune]func()error
}

func (sm *StackMachine) pop() (num int, err error) {
	if len(sm.numStack) < 1 {
		err = errors.New("Not enough numbers")
		return
	}
	num, sm.numStack = sm.numStack[len(sm.numStack)-1], sm.numStack[:len(sm.numStack)-1]
	return
}

func (sm *StackMachine) calculate(userInput string) int {
	for _, r := range userInput {
		if sm.operators[r] != nil {
			err := sm.operators[r]()
			if err != nil{
				return -1
			}
		} else {
			sm.numStack = append(sm.numStack, int(r - '0'))
		}
	}
	return sm.numStack[len(sm.numStack) - 1]
}

func (sm *StackMachine) add() (err error) {
	var num1, num2 int
	num1, err = sm.pop()
	num2, err = sm.pop()
	if err != nil {
		return
	}
	sm.numStack = append(sm.numStack, num1 + num2)
	return
}

func (sm *StackMachine) multiply() (err error) {
	var num1, num2 int
	num1, err = sm.pop()
	num2, err = sm.pop()
	if err != nil {
		return
	}
	sm.numStack = append(sm.numStack, num1 * num2)
	return
}