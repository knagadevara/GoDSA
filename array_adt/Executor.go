package arrayadt

import (
	"fmt"
	"slices"
)

func Exec() {
	IntArr := MakeArrAdt[int](11)
	fmt.Println(IntArr.RmPop(0))
	IntArr.Addpend(10)
	IntArr.Addpend(20)
	IntArr.Addpend(30)
	IntArr.Addpend(40)
	IntArr.Insert(10, 100)
	IntArr.Insert(0, 99)
	IntArr.Insert(3, 56)
	IntArr.Insert(4, 86)
	IntArr.Insert(9, 16)
	IntArr.Insert(9, 65)
	slices.Sort(IntArr.adt)
	IntArr.SortedInsert(39)
	fmt.Println(IntArr.IsSorted())
	fmt.Println(IntArr.ArContainsBS(0))
	fmt.Println(IntArr.ArContainsBS(31))
	fmt.Println("Max: ", IntArr.Max())
	fmt.Println("Min: ", IntArr.Min())
	fmt.Println("Sum: ", Sum(IntArr.adt))
	IntArr.Display()
	fmt.Println("Riversed")
	IntArr.Reverse2()
	IntArr.Display()
	var a1 = make([]int, 4)
	a1 = []int{6, 7, 8, 9}
	var a2 = make([]int, 5)
	a2 = []int{1, 2, 3, 4, 5}
	mergedarr := MergeSortedArray[int](a1, a2)
	for i, v := range mergedarr {
		fmt.Println(i, v)
	}
}
