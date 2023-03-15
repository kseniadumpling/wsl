package pkg

func Fn() {
	switch {
	case true, false:
		fmt.Println("ok")
	case true || false:
		fmt.Println("ok")
	case true, false:
		fmt.Println("ok")
	case true || false:
		fmt.Println("ok")
	case true,
		false: // want "block should not start with a whitespace"

		fmt.Println("starting whitespace multiline case")
	case true ||
		false: // want "block should not start with a whitespace"

		fmt.Println("starting whitespace multiline case")
	case true,
		false:
		fmt.Println("ending whitespace multiline case")

	case true ||
		false:
		fmt.Println("ending whitespace multiline case")

	case true, false:
		fmt.Println("all")
	}
}
