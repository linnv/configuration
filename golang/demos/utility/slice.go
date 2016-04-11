package utility

// 找出a中不存在b中元素的数组
func IntArraySelectDiff(a, b []int) []int {
	r := make([]int, 0, len(a))
	for _, _a := range a {
		if !InIntArray(_a, b) {
			r = append(r, _a)
		}
	}
	return r
}

func UniqSlice(a []int) []int {
	var res = make([]int, len(a))
	var mp = make(map[int]bool)
	var index = 0
	for _, i := range a {
		if !mp[i] {
			res[index] = i
			index++
		}
		mp[i] = true
	}
	return res[:index]
}

// UniqStr returns a copy if the passed slice with only unique string results.
func UniqStr(col []string) []string {
	m := map[string]struct{}{}
	for _, v := range col {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}
	list := make([]string, len(m))

	i := 0
	for v := range m {
		list[i] = v
		i++
	}
	return list
}

func InIntArray(i int, ints []int) bool {
	for _, v := range ints {
		if i == v {
			return true
		}
	}
	return false
}

func AppendIfMissing(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func UniqueSlice(s []int) []int {
	l := len(s)
	for i := 0; i < l; i++ {
		for j := l; j > i; j-- {
			if s[i] == s[j] {
				s = append(s[:i], s[j:]...)
			}
		}
	}
	return s
}

// func IntArrayCommon(a, b []int) []int {
// 	diff := IntArraySubtract(a, b)
// 	return IntArraySubtract(a, diff)
//
// func TestIntArrayCommon(t *testing.T) {
// 	a := []int{1, 2, 3}
// 	b := []int{2, 3, 4}
// 	c := IntArrayCommon(a, b)
// 	rc := []int{2, 3}
//
// 	if !IntArrayEqual(c, rc) {
// 		//@todoDelelte
// 		fmt.Printf("c,cap(c):%d,len(c):%d  : %+v\n", c, cap(c), len(c))
// 		//@todoDelelte
// 		fmt.Printf("rc,cap(rc):%d,len(rc):%d  : %+v\n", rc, cap(rc), len(rc))
// 		t.Error(" TestIntArrayCommon: failed")
// 	}
// }
// }
