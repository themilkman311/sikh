package sikh

import (
	"log"
	"os"
	"sync/atomic"

	"github.com/kyleraywed/sikh/keymaps"
	"golang.org/x/term"
)

type Sikh struct {
	isRunning atomic.Bool
}

func (sikh *Sikh) Start(handler func(string)) {
	if sikh.isRunning.Load() {
		return
	}

	sikh.isRunning.Store(true)
	defer sikh.isRunning.Store(false)

	for {
		if !sikh.isRunning.Load() {
			return
		}

		rep, err := sikh.getKeystroke()
		if err != nil {
			log.Println(err)
		}

		if job, ok := sikh.toString(rep); ok {
			handler(job)
		}
	}
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
