package main

//MyInt is myint
type MyInt int

// Inc is func
func (nn MyInt) Inc() MyInt { return nn + 1 }

func main() {
	var n MyInt
	n = 8
	println(n)
	n = n.Inc()
	println(n)
}
