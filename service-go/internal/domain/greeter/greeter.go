package greeter

func Greet(name string) string {
	if name == "" {
		name = "stranger"
	}

	return "Hello, " + name
}
