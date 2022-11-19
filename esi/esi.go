package esi

import (
	"fmt"
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
	s := ""
	for _, v := range vs {
		s += fmt.Sprintf("%v", v)
	}
	return s
}
