package main

func main() {
	m := make(map[int]int)
	go func() {
		m[1] = 1
	}()
	m[2] = 2
}
