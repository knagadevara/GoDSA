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
	slices.Sort(IntArr.adt)
	fmt.Println(IntArr.ArContainsBS(0))
	fmt.Println(IntArr.ArContainsBS(31))
	fmt.Println("Max: ", IntArr.Max())
	fmt.Println("Min: ", IntArr.Min())
	fmt.Println("Sum: ", Sum(IntArr.adt))
	IntArr.Display()
	fmt.Println("Riversed")
	IntArr.Reverse2()
	IntArr.Display()

}
