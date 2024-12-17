package arrayadt

import "fmt"

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
	IntArr.Display()
	fmt.Println(IntArr.RmPop(0))
	fmt.Println(IntArr.RmPop(0))
	fmt.Println(IntArr.RmPop(0))
	fmt.Println(IntArr.RmPop(0))
	fmt.Println("Removed First Element")
	IntArr.Display()
	IntArr.Reverse()
	fmt.Println(IntArr.Swap(0, 4))
	IntArr.Display()
}
