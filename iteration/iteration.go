package iteration

// := syntax is shorthand for declaring and initializing a variable,
// e.g. for var f string = "apple"
// we can do f := "apple"
func Repeat(character string) string {
	var repeat string
	for i := 0; i < 5; i++ {
		repeat += character
	}
	return repeat
}
