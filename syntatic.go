package main

import (
	"fmt"
	"os"
)

/*
*	submachines
 */
const (
	FSM_PROG = iota
	FSM_EXPR
)

const ERROR = -1

/*
*	State analysis struct
 */
type analysisState struct {
	currSubmachineState int
	submachine          int
	currToken           *Token
	lastToken           *Token
	getTokenFlag        bool
}

/*
 *	Stack node has the sub-machine and state
 */
type submachineStackNode struct {
	submachine int
	state      int
	next       *submachineStackNode
}

/*
 * 	Stack has just a top node and size
 */
type submachineStack struct {
	top  *submachineStackNode
	size int
}

/*
* push on stack
 */
func (s *submachineStack) push(submachine, state int) {
	sn := &submachineStackNode{submachine: submachine, state: state}
	sn.next = s.top
	s.top = sn
	s.size++

	globalState.getTokenFlag = false
}

/*
*	pop on stack
 */
func (s *submachineStack) pop() {
	if s.size > 0 {
		globalState.submachine = s.top.submachine
		globalState.currSubmachineState = s.top.state
		s.size--

		s.top = s.top.next
		globalState.getTokenFlag = false
	} else {
		globalState.currSubmachineState = ERROR
	}
}

/*
* Globals
 */
var submachines []func(t *Token) int
var globalState analysisState
var submachineCallsStack submachineStack = submachineStack{}

/*
*	Will initialize global state with its correct values
 */
func init() {
	submachines = []func(t *Token) int{
		fsmProg,
		fsmExpr,
	}
	globalState = analysisState{0, 0, nil, nil, true}
}

/*
*	Analyze is the function that truly starts the parsing
 */
func Analyze(file *os.File) int {
	for {
		if globalState.getTokenFlag {
			globalState.currToken = GetToken(file)
			if globalState.currToken == nil {
				break
			}
		}

		globalState.getTokenFlag = true
		globalState.currSubmachineState = submachines[globalState.submachine](globalState.currToken)
		if globalState.currSubmachineState == ERROR {
			fmt.Println("[ERROR] Compilation error")
			return 1
		}
	}

	fmt.Println("[INFO] Compilation Successful")
	semanticFlushCode()
	return 0
}

/*
*	Call submachine
 */
func callSm(submachine, retSt int) int {
	submachineCallsStack.push(globalState.submachine, retSt)
	globalState.submachine = submachine
	return 0
}

/////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////
/////////////////////////// SUB MACHINES FUNCTIONS //////////////////////////////
/////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

func fsmProg(t *Token) int {
	switch globalState.currSubmachineState {
	case 0:
		return callSm(FSM_EXPR, 1)

	case 1:
		return callSm(FSM_EXPR, 1)
	}

	return ERROR
}

func fsmExpr(t *Token) int {
	switch globalState.currSubmachineState {
	case 0:
		if t.Class == TOKEN_CLASS_S || t.Class == TOKEN_CLASS_K || t.Class == TOKEN_CLASS_I {
			semanticEnterToken(t)
			return 1
		} else if t.Class == TOKEN_CLASS_OPEN_SCOPE {
			semanticNewScope()
			return 2
		}

		return ERROR

	case 1:
		submachineCallsStack.pop()
		return globalState.currSubmachineState

	case 2:
		if t.Class == TOKEN_CLASS_CLOSE_SCOPE {
			semanticCloseScope()
			return 1
		}
		return callSm(FSM_EXPR, 2)
	}
	return ERROR
}
