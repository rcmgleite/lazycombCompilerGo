package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	generated  bytes.Buffer
	outputfile string = "out.c"
)

var template = `
#include<stdio.h>
#include<vm.h>

int main() {
	novo_escopo();

@PLACEHOLDER

	imprime();

	return 0;
}
`

func semanticEnterToken(t *Token) {
	generated.WriteString("	entra(" + t.Value + ");\n")
	generated.WriteString("	tenta_reduzir();\n")
}

func semanticNewScope() {
	generated.WriteString("	novo_escopo();\n")
}

func semanticCloseScope() {
	generated.WriteString("	fecha_escopo();\n")
	generated.WriteString("	tenta_reduzir();\n")
}

func semanticFlushCode() {
	template = strings.Replace(template, "@PLACEHOLDER", generated.String(), -1)
	err := ioutil.WriteFile(outputfile, []byte(template), 0644)
	if err != nil {
		fmt.Println("[ERROR] Unbale to write output to file.", err.Error())
	}
}
