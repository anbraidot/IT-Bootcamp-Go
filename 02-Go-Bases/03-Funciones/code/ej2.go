package main

func main() {
	qualifications := [...]int{10, 10, 7, 4, 9, 6}

	println(qualificationAvg(qualifications[:]))

}

func qualificationAvg(qualifications []int) (result float32) {
	for _, value := range qualifications {
		result += float32(value)
	}
	result /= float32(len(qualifications))
	return
}