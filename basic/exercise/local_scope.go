package exercise

var a = "G"

func LocalScope() {
	n()
	m()
	n()
}

func n() { print(a) }

func m() {
	a := "O"
	print(a)
}
