# Sikh

A stdin key-hooking utility in less than 100 LOC. It reads the raw input from a keypress, compares it against a map, and returns a string coresponding with the key pressed.

Module: `github.com/themilkman311/sikh`

```
(hook *InputHook) Start(handler func(string)) <-chan struct{}
(hook *InputHook) Stop()
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
	var hook InputHook

    // switch over the string and implement logic
	done := hook.Start(func(s string) {
		switch s {
		case "[Ctrl+c]":
			hook.Stop() // do not forget this
		default:
			fmt.Println(s)
		}
	})

	<-done
}
```