# Sikh

A small **s**td**i**n **k**ey-**h**ooking library. It reads the raw input from a keypress, compares it against a map, and returns a string coresponding with the key pressed.

Module: `github.com/kyleraywed/sikh`

```
(sikh *Sikh) Start(handler func(string)) <-chan struct{}
(sikh *Sikh) Halt()
```

Example

```go
package main

import (
    "fmt"
    "github.com/kyleraywed/sikh"
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

Notes and design

- If you don't include some logic to call Halt(), you will have to kill the process yourself.
