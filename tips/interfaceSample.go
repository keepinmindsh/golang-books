package tips

import "fmt"

type Pet struct {
	Name string
	Age  int
}

func (p Pet) getName() {
	fmt.Printf("pet name is %s.\n", p.Name)
}

func (p Pet) getAge() {
	fmt.Printf("pet age is %d.\n", p.Age)
}

type GetName interface {
	getName()
}

type GetAge interface {
	getAge()
}

type PetInfo interface {
	GetName
	GetAge
}

func InterfaceSample() {
	var interfaceSample PetInfo

	interfaceSample.getName()
}
