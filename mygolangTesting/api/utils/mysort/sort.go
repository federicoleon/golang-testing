package mysort

import "sort"

// BubbleSort takes a slice if int and sorts the elements
func BubbleSort(elements []int) {
	keepWorking := true
	for keepWorking {
		//Comment out the following to test infinit loop
		keepWorking = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] < elements[i+1] {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

// BubbleSortDesc changes to Decending order
func BubbleSortDesc(elements []int) {
	keepWorking := true
	for keepWorking {
		//keepWorking = false
		for i := 0; i < len(elements)-1; i++ {
			keepWorking = false
			if elements[i] < elements[i+1] {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

// Sort in increasing order takes a var elements as a slice of ints
func Sort(elements []int) {
	sort.Ints(elements)
}
