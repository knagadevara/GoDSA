package stack

func CreateStkArr() *[]int {
	var stkArr []int
	return &stkArr
}

func PushAr(s *[]int, val int) *[]int {
	*s = append(*s, val)
	return s
}

func PopAr(s []int) (*[]int, int) {
	arrLnth := len(s)
	if arrLnth != 0 {
		popped := (s[arrLnth-1])
		new_slice := s[0 : len(s)-2]
		s = append(s, new_slice...)
		return &s, popped
	} else {
		return nil, 0
	}
}
func PeekAr(s []int) int {
	arrLnth := len(s)
	if arrLnth != 0 {
		return (s[arrLnth-1])
	} else {
		return 0
	}
}
