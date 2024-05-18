package animal

// Animal struct with a private field.
type Animal struct {
	name string
}

// NewAnimal is a constructor function that initializes a new Animal instance.
func NewAnimal(name string) *Animal {
	return &Animal{
		name: name,
	}
}

// GetAnimalName is a getter method for the name field.
func (a *Animal) GetAnimalName() string {
	return a.name
}

// SetAnimalName is a setter method for the name field.
func (a *Animal) SetAnimalName(name string) {
	a.name = name
}