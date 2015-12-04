package main

import (
	"fmt"
	"os"
)

/*
*	TOKEN CLASS
 */
const (
	TOKEN_CLASS_S = iota
	TOKEN_CLASS_K
	TOKEN_CLASS_I
	TOKEN_CLASS_OPEN_SCOPE
	TOKEN_CLASS_CLOSE_SCOPE
)

/*
*	Token struct
 */
type Token struct {
	Class int
	Value string
}

/*
*	Return a new token built from the byte read
 */
func newToken(b []byte) *Token {
	bString := string(b)
	var class int

	if bString == "S" || bString == "s" {
		class = TOKEN_CLASS_S
	} else if bString == "K" || bString == "k" {
		class = TOKEN_CLASS_K
	} else if bString == "I" || bString == "i" {
		class = TOKEN_CLASS_I
	} else if bString == "(" {
		class = TOKEN_CLASS_OPEN_SCOPE
	} else if bString == ")" {
		class = TOKEN_CLASS_CLOSE_SCOPE
	} else {
		return nil
	}

	return &Token{class, bString}
}

/*
*	String method for Token
 */
func (t *Token) String() {
	fmt.Println("Token: ")
	fmt.Println("Class = ", t.Class, "Value = ", t.Value)
}

/*
*	Print formatted token on stdout
 */
func (t *Token) Print() {
	t.String()
}

/*
*	GetToken return next token on file
 */
func GetToken(file *os.File) *Token {
	b := make([]byte, 1)
	n, err := file.Read(b)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if n == 0 {
		fmt.Println("[DEBUG] EOF")
		return nil
	}

	return newToken(b)
}
