package main

var toPrint = []string{
	"notPro",
}

func main() {
	a := 1
	b := 2
	for _, f := range toPrint {
		println(f)
	}
	if true {
		add(a, b)
	}

}
func add(a, b int) {
	println(a + b)
}
