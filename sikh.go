package main

import (
	"fmt"
	"log"
	"os"
	"sync/atomic"

	"github.com/themilkman311/sikh/keymaps"
	"golang.org/x/term"
)

type Sikh struct {
	isRunning atomic.Bool
}

func (sikh *Sikh) Start(handler func(string)) <-chan struct{} {
	done := make(chan struct{})

	if sikh.isRunning.Load() {
		return nil
	}

	sikh.isRunning.Store(true)

	go func() {
		defer close(done)
		for {
			if !sikh.isRunning.Load() {
				return
			}

			rep, err := sikh.getKeystroke()
			if err != nil {
				log.Fatal(err) // i know
			}

			if job, ok := sikh.toString(rep); ok {
				handler(job)
			}
		}
	}()
	return done
}

func (sikh *Sikh) Halt() {
	sikh.isRunning.Store(false)
}

func (sikh *Sikh) getKeystroke() ([4]byte, error) {
	var rep [4]byte

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return rep, err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 4)
	_, err = os.Stdin.Read(b)

	if err != nil {
		return rep, err
	}

	copy(rep[:], b[:4])
	return rep, nil
}

func (sikh *Sikh) toString(rep [4]byte) (string, bool) {
	job, ok := keymaps.StandardMap[rep]
	return job, ok
}

func main() {
	var sikh Sikh

	done := sikh.Start(func(s string) {
		switch s {
		case "[Ctrl+c]":
			sikh.Halt()
		default:
			fmt.Println(s)
		}
	})
	<-done
}
