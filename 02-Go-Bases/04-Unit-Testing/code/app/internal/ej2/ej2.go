package ej2

func QualificationAvg(qualifications []int) (result float32) {
	if len(qualifications) > 0 {
		for _, value := range qualifications {

			if value < 0 {
				return 0.00
			}

			result += float32(value)
		}
		result /= float32(len(qualifications))
	}
	return
}