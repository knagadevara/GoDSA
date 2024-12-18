package arrayadt

import (
	"fmt"
	"slices"
)

func Exec() {
	IntArr := MakeArrAdt[int](10)
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
	IntArr.Display()
	IntArr.Display()
	IntArr.Reverse()
	IntArr.ArContains(16)
	slices.Sort(IntArr.adt)
	fmt.Println(IntArr.ArContainsBS(40))
	IntArr.Display()
}
