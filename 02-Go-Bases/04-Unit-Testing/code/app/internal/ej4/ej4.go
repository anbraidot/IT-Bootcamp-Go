package ej4

func Operation(operation string) (func(...int) float32, string) {
	switch operation {
	case "minimum":
		return MinFunc, ""
	case "average":
		return AverageFunc, ""
	case "maximum":
		return MaxFunc, ""
	default:
		return nil, "Invalid operation"
	}
}

func MinFunc(qualifications ...int) float32 {

	if len(qualifications) == 0 {
		return 0.0
	}

	result := float32(qualifications[0])
	for _, value := range qualifications {

		if value < 0 {
			return 0.0
		}

		if value < int(result) {
			result = float32(value)
		}
	}
	return result
}

func AverageFunc(qualifications ...int) (result float32) {
	if len(qualifications) == 0 {
		return 0.0
	}
	
	for _, value := range qualifications {
		if value < 0 {
			return 0.0			
		}

		result += float32(value)
	}
	result /= float32(len(qualifications))
	return
}

func MaxFunc(qualifications ...int) float32 {
	if len(qualifications) == 0 {
		return 0.0
	}

	result := float32(qualifications[0])
	for _, value := range qualifications {
		if value < 0 {
			return 0.0
		}
		
		if value > int(result) {
			result = float32(value)
		}
	}
	return result
}