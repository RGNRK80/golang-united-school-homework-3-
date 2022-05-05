package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	arKey := make([]int, 0, 0)

	for k, _ := range input {
		arKey = append(arKey, k)
		//fmt.Println(k)

	}
	sort.Ints(arKey)
	//fmt.Println(arKey)

	arStr := make([]string, 0, 0)

	for i := 0; i < len(input); i++ {
		arStr = append(arStr, input[arKey[i]])
	}

	//Place your code here
	return arStr
}
