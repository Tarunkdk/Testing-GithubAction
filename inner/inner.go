package inner

type name struct {
	name string
}

func hello() {
	n := name{
		name: "jack",
	}
	println("hello", n.name)
}
