package base_type

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func TestInterface() {
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}