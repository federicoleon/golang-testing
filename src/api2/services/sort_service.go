package services

import "github.com/federicoleon/golang-testing/src/api2/utils/sort"

const (
	PublicConst = "public"
)

func Sort(elements []int) {
	if len(elements) <= 20000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
