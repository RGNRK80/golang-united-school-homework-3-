package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	t := make([]int64, len(input), cap(input))
	copy(t, input)

	for i := 0; i < len(t); i++ {
		t[i] = input[len(t)-i-1]
	}

	return t
}
