package resistorcolorduo

var colours = []string{
	"black",
	"brown",
	"red",
	"orange",
	"yellow",
	"green",
	"blue",
	"violet",
	"grey",
	"white",
}

var colourMap map[string]int

func init() {
	colourMap = make(map[string]int)
	for i, c := range colours {
		colourMap[c] = i
	}
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	return 10*colourMap[colors[0]] + colourMap[colors[1]]
}
