package ej5

func AnimalOrquestator(animal string) (func(q int) float64, string) {
	switch animal {
	case "dog":
		return AnimalDog, ""
	case "cat":
		return AnimalCat, ""
	case "tarantula":
		return AnimalTarantula, ""
	case "hamster":
		return AnimalHamster, ""
	default:
		return nil, "Invalid animal"
	}
}

func AnimalDog(q int) float64 {
	if q <= 0 {
		return 0.00
	}
	return float64(q) * 10.00
}

func AnimalCat(q int) float64 {
	if q <= 0 {
		return 0.00
	}
	return float64(q) * 5.00
}

func AnimalTarantula(q int) float64 {
	if q <= 0 {
		return 0.00
	}
	return float64(q) * 0.15
}

func AnimalHamster(q int) float64 {
	if q <= 0 {
		return 0.00
	}
	return float64(q) * 0.25
}