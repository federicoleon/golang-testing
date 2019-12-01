package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Constants(t *testing.T) {
	if privateConstant != "private" {
		t.Error("The constant is not private")
	}
}
func Test_BSort(t *testing.T) {
	elements := []int{9, 8, 7, 5, 4, 6, 1, 0, 3, 2}
	// fmt.Println(elements)
	BSort(elements)
	lelement := elements[len(elements)-1]
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])
	if elements[len(elements)-1] != 0 {
		t.Error("The lastElement should be 9. it is ", lelement)
	}
	assert.EqualValues(t, 0, lelement)
	felement := elements[0]
	assert.EqualValues(t, 9, felement)
	if elements[0] != 9 {
		t.Error("The first element should be 9, it is ", felement)
	}
	// fmt.Println(elements)
}
func Test_BSort10X3(t *testing.T) {
	elements := getElements(10001)
	// fmt.Println(elements)
	assert.NotNil(t, elements)
	BSort(elements)
	lelement := elements[len(elements)-1]
	if elements[len(elements)-1] != 0 {
		t.Error("The lastElement should be 9. it is ", lelement)
	}
	felement := elements[0]
	if elements[0] != 10000 {
		t.Error("The first element should be 10000, it is ", felement)
	}
	// fmt.Println(elements)
}

func Test_Sort(t *testing.T) {
	elements := []int{9, 8, 7, 5, 4, 6, 1, 0, 3, 2}
	fmt.Println(elements)

	Sort(elements)
	lelement := elements[len(elements)-1]
	if elements[len(elements)-1] != 9 {
		t.Error("The lastElement should be 9. it is ", lelement)
	}
	felement := elements[0]
	if elements[0] != 0 {
		t.Error("The first element should be 0, it is ", felement)
	}
	fmt.Println(elements)
}
func Test_SortAssert(t *testing.T) {
	elements := []int{9, 8, 7, 5, 4, 6, 1, 0, 3, 2}
	fmt.Println(elements)
	assert.NotNil(t, elements)
	Sort(elements)
	lelement := elements[len(elements)-1]
	assert.EqualValues(t, 9, lelement)
	felement := elements[0]
	assert.EqualValues(t, 0, felement)
	fmt.Println(elements)
}
func Test_Sort10X3(t *testing.T) {
	elements := getElements(10001)
	// fmt.Println(elements)

	Sort(elements)
	lelement := elements[len(elements)-1]
	if elements[len(elements)-1] != 10000 {
		t.Error("The lastElement should be 10000. it is ", lelement)
	}
	felement := elements[0]
	if elements[0] != 0 {
		t.Error("The first element should be 0, it is ", felement)
	}
	// fmt.Println(elements)
}

func BenchmarkBubbleSort(b *testing.B) {

	elements := []int{9, 8, 7, 5, 4, 6, 1, 0, 3, 2}
	for i := 0; i < b.N; i++ {
		BSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := []int{9, 8, 7, 5, 4, 6, 1, 0, 3, 2}
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func getElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}
func BenchmarkBubbleSort10X3(b *testing.B) {
	elements := getElements(10000)
	for i := 0; i < b.N; i++ {
		BSort(elements)
	}
}

func BenchmarkBSortLimitGT10X3(b *testing.B) {
	elements := getElements(10000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkSort10X3(b *testing.B) {
	elements := getElements(10000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
