package sort

import "sort"

func BubbleSort(elements []int)  {
	keepWorking := true
	for keepWorking {
		keepWorking = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

func Sort(elements []int) {
	sort.Ints(elements)
}

func GetElements(n int) []int {
	results := make([]int, n)
	j := 0;
	for i := n-1; i > 0; i-- {
		results[j] = i
		j++
	}
	return results
}
