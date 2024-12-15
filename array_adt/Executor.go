package arrayadt

import "fmt"

func Exec() {
	IntArr := MakeArrAdt[int](10)
	IntArr.Addpend(10)
	IntArr.Addpend(20)
	IntArr.Addpend(30)
	IntArr.Addpend(40)
	IntArr.Insert(10, 20)
	IntArr.Display()
	IntArr.Reverse()
	fmt.Println("Display Reverse")
	IntArr.Display()
	fmt.Println(IntArr.ArContains(0))
	fmt.Println(IntArr.Get(0))
	IntArr.Display()
	fmt.Println(IntArr.Swap(0, 4))
	fmt.Println("After Swap")
	IntArr.Display()
	IntArr.Insert(0, 99)
	IntArr.Display()

}
