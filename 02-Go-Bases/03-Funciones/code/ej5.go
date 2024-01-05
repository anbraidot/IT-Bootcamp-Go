package main

func main() {

	const (
		dog    = "dog"
		cat    = "cat"
		tarantula = "tarantula"
		hamster = "hamster"
	)

	function, _ := animalOrquestator(dog)
	println(function(10))

}

func animalOrquestator(animal string) (func(q int) float32, string) {
	switch animal {
	case "dog":
		return animalDog, ""
	case "cat":
		return animalCat, ""
	case "tarantula":
		return animalTarantula, ""
	case "hamster":
		return animalHamster, ""
	default:
		return nil, "Invalid animal"
	}
}

func animalDog(q int) float32 {
	return float32(q) * 10.00
}

func animalCat(q int) float32 {
	return float32(q) * 5.00
}

func animalTarantula(q int) float32 {
	return float32(q) * 0.15
}

func animalHamster(q int) float32 {
	return float32(q) * 0.25
}