package puppy

import (
	"github.com/danjelhysenaj-dev/go-dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woff Woof!"
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
