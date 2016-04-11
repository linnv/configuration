// Package main provides ...
package newDir

import "fmt"

type MS struct {
	Exit    bool
	Index   int
	Content interface{}
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	// if ms[111].Exit {
	// 	os.Stdout.Write(append([]byte("good"), '\n'))
	// }
	// mpSlice := make([]*MS, len)
	// for i := 0; i < len; i++ {
	// 	mpSlice[i] = &MS{Index: i}
	// }

	// tmp := ms[0]
	// tmp.Index = 111
	// ms[0].Index = 3333
	// ms[0] = *mpSlice[0]

	// fmt.Printf("m: %+v\n", m)
	// ms[0] = m
	// fmt.Printf("ms[0]: %+v\n", ms[0])

	len := 100
	ms := make(map[int]MS, len)
	mpSlice := make(map[int]MS, len)
	for i := 0; i < len; i++ {
		mpSlice[i] = MS{Index: i * 9}
	}

	// var m MS
	m := mpSlice[0]
	m.Index = 333
	ms[0] = m

	// ms[0] = mpSlice[0]
	// ms[0].Index = 333

	fmt.Printf("ms[0]: %+v\n", ms[0])

	fmt.Printf("mpSlice[0]: %+v\n", mpSlice[0])

	// m := make(map[int][]int, 3)
	// m[3] = []int{1, 34}
	// //@toDelete
	// //@toDelete
	// fmt.Printf("  1len(m[1]: %+v\n", len(m[1]))
	// fmt.Printf("  len(m): %+v\n", len(m))
	// if len(m[1]) > 0 {
	// 	//@toDelete
	// 	fmt.Printf("  m[1]: %+v\n", m[1])
	// }

	// fmt.Printf(" 2 len(m[1]: %+v\n", len(m[1]))
	// var mp = make(map[int]bool)
	// mp[4] = true
	// for i := 0; i < 10; i++ {
	// 	//@toDelete
	// 	fmt.Printf("  mp[i]: %+v\n", mp[i])
	// }

	// raw := `
	// 	{ID:[48 57 55 101 99 55 101 52 99 97 56 50 57 48 98 54] Body:[123 34 84 121 112 101 34 58 45 49 48 44 34 65 99 116 105 111 110 34 58 49 44 34 73 100 115 34 58 91 49 50 51 52 44 52 51 52 51 44 52 51 52 51 93 44 34 82 101 108 73 100 115 34 58 110 117 108 108 125] Timestamp:1452246666687973765 Attempts:2 NSQDAddress:sf41:4150 Delegate:0xc8200c0360 autoResponseDisabled:0 responded:0}
	// `
	// // raw := `[48 57 55 101 99 55 101 52 99 97 56 50 57 48 98 54] Body:[123 34 84 121 112 101 34 58 45 49 48 44 34 65 99 116 105 111 110 34 58 49 44 34 73 100 115 34 58 91 49 50 51 52 44 52 51 52 51 44 52 51 52 51 93 44 34 82 101 108 73 100 115 34 58 110 117 108 108 125`
	// // raw := `48 57 55 101 99 55 101 52 99 97 56 50 57 48 98 54`
	// s := []byte(raw)
	// //@toDelete
	// fmt.Printf("  s: %s\n", string(s))
	println("-----------------------------JustDemo end>>>")
	return
}
