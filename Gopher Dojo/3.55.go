package main

func main() {
	ns := [...]string{"-偶数", "-奇数"}
	for i := 1; i <= 100; i++ {
		print(i)
		if i%2 == 0 {
			println(ns[0])
		} else {
			println(ns[1])
		}
	}
}
