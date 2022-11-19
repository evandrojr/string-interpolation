package main

import "github.com/evandrojr/string-interpolation/esi"

func main() {
	esi.Println("First ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)
	esi.Println("First ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)

	f := esi.Sprint("Second ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)

	println(f)
}
