# Sikh

A stdin key-hooking utility in less than 100 LOC. It reads the raw input from a keypress, compares it against a map, and returns a string coresponding with the key pressed.

Module: `github.com/themilkman311/sikh`

```
(sikh *Sikh) Start(handler func(string)) <-chan struct{}
(sikh *Sikh) Halt()
```

Example

```go
package main

import (
    "fmt"
    "github.com/themilkman311/sikh"
)

func main() {
    // create the hook
	var sikh Sikh

    // switch over the string and implement logic
	done := sikh.Start(func(s string) {
		switch s {
		case "[Ctrl+c]":
			sikh.Halt() // do not forget this
		case "[Esc]":
			fmt.Println("Trying to escape? Try ctrl+c")
		default:
			fmt.Println(s)
		}
	})

	<-done
}
```