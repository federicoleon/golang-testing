package services

import (
	"github.com/federicoleon/golang-testing/src/api2/utils/sort"
	"testing"
)

func TestSort(t *testing.T)  {
	elements := sort.GetElements(10)

	Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("last elements should be 9")
	}
}

func TestSortMoreThan10000(t *testing.T)  {
	elements := sort.GetElements(10001)

	Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 10000 {
		t.Error("last elements should be 10000")
	}
}

func BenchmarkSort10K(b *testing.B) {
	elements := sort.GetElements(20000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkSort100K(b *testing.B) {
	elements := sort.GetElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
