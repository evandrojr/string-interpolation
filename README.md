# String interpolation in Go finally made easy...

Simple string interpolation for golang. Interpolates anything in an easy way.

No need to pass the format parameters %d, %s, %t... anymore!

## Installation

```
go get github.com/evandrojr/string-interpolation
```

## Usage

```
package main

import (
	"fmt"

	"github.com/evandrojr/string-interpolation/esi"
)

func main() {
	esi.Print("Print ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)
	fmt.Print(" no line bread")
	fmt.Println()
	esi.Println("Println ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)
	f := esi.Sprint("Sprint ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)
	fmt.Print(f)
}

```
### Output

```
Print 10 7 interpolates anything true 3.4e+10 no line break
Println 10 7 interpolates anything true 3.4e+10
Sprint 10 7 interpolates anything true 3.4e+10% 
```
