package esi

import (
	"fmt"
	"strings"
)

func Print(vs ...any) {
	for _, v := range vs {
		fmt.Printf("%v", v)
	}
}

func Println(vs ...any) {
	for _, v := range vs {
		fmt.Printf("%v", v)
	}
	fmt.Println()
}

func Sprint(vs ...any) string {

	var sb strings.Builder

	for _, v := range vs {
		sb.WriteString(fmt.Sprintf("%v", v))
	}
	return sb.String()
}
