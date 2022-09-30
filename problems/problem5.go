package problems

import "fmt"

type AnimalAction interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

func (bird Bird) Eat() {
	fmt.Println("grass")
}

func (bird Bird) Move() {
	fmt.Println("walk")
}

func (bird Bird) Speak() {
	fmt.Println("moo")
}

func (snake Snake) Eat() {
	fmt.Println("grass")
}

func (snake Snake) Move() {
	fmt.Println("walk")
}

func (snake Snake) Speak() {
	fmt.Println("moo")
}

func AnimalAct(action string, animal AnimalAction) {
	switch action {
	case "MOVE":
		animal.Move()
	case "EAT":
		animal.Eat()
	case "SPEAK":
		animal.Speak()
	}
}

func main() {
	AnimalAct("MOVE", Cow{})
	AnimalAct("EAT", Cow{})
	AnimalAct("SPEAK", Cow{})

	AnimalAct("MOVE", Snake{})
	AnimalAct("EAT", Snake{})
	AnimalAct("SPEAK", Snake{})

	AnimalAct("MOVE", Bird{})
	AnimalAct("EAT", Bird{})
	AnimalAct("SPEAK", Bird{})
}
