package animal

type Cat struct {
	Animal
}

func NewCat(name string) *Cat {
	return &Cat{
		Animal: *NewAnimal(name),
	}
}