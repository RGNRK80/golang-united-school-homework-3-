package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var s float32
	for i := 0; i < len(input); i++ {
		s += input[i]
	}
	s = s / (float32)(len(input))
	return s
}
