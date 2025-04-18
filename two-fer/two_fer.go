// Package to greet a person in line when giving them a cookie.
package twofer

// Returns a greeting based on the name. If name is 'zero', uses "you" instead.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
