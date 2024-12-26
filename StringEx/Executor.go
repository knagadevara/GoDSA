package stringex

import "fmt"

func Exec() {
	fmt.Printf("%s\n", MakeCaps("hello"))
	fmt.Println(CountVouls("Sai   Karthik"))
	fmt.Println(ValidChar("SaiKarthik"))
}
