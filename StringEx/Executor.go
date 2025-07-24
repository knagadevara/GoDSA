package stringex

import "fmt"

func Exec() {
	fmt.Printf("%s\n", MakeCaps("hello"))
	fmt.Println(CountVouls("Vai   SamSung"))
	fmt.Println(ValidChar("VaiSamSung"))
	fmt.Println(RevStr2("Apple"))
	fmt.Println(Palindrome("SaaS"))
	CountDuplicate("Aalmas")
}
