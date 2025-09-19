package sikh

import (
	"log"
	"os"
	"sync"

	"github.com/themilkman311/sikh/keymaps"
	"golang.org/x/term"
)

type InputHook struct {
	representation [4]byte
	running        bool
	runningMutex   sync.Mutex // don't touch
}

func (hook *InputHook) Start(handler func(string)) <-chan struct{} {
	done := make(chan struct{})

	if hook.running {
		return nil
	}

	hook.runningMutex.Lock() // I know this data race can never happen.
	hook.running = true      // The future may change that.
	hook.runningMutex.Unlock()

	go func() {
		defer close(done)
		for {
			if !hook.running {
				break
			}

			if err := hook.getKeystroke(); err != nil {
				log.Fatal(err)
			}

			job, err := hook.string()
			if err != nil {
				log.Fatal(err)
			}
			handler(job)
		}
	}()
	return done
}

func (hook *InputHook) Stop() {
	hook.runningMutex.Lock()
	hook.running = false
	hook.runningMutex.Unlock()
}

func (hook *InputHook) getKeystroke() error {
	// For whatever reason, putting and leaving the term in raw mode during operation
	// causes issues with keypresses; so much so that at the cost of very little
	// performance (thank goodness), I've chosen to just go in and out with every keypress.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 4)
	_, err = os.Stdin.Read(b)
	if err != nil {
		return err
	}

	copy(hook.representation[:], b[:4])
	return nil
}

func (hook *InputHook) string() (string, error) {
	return keymaps.StandardMap[hook.representation], nil
}
