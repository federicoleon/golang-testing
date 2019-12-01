package services

import (
	"fmt"
	"github.com/golang-testing/mygolangTesting/api/utils/mysort"
)

const (
	privateConstant = "private"
)

// BSort takes a list of its and sorts them
func BSort(elements []int) {
	mysort.BubbleSort(elements)
}

// Sort takes a list of its and sorts them
func Sort(elements []int) {
	mysort.Sort(elements)
}

// BSortLimit Limits by using native sort for elements > 10000
func BSortLimit(elements []int) {
	if len(elements) <= 10000 {
		mysort.BubbleSort(elements)
		return
	}
	mysort.Sort(elements)
}

func init() {
	fmt.Println("Init sort Service")
	LocationsService = &locationsService{}
}
