package resistorcolor

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

// Colors returns the list of all colors.
func Colors() []string {
	return colours
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	return colourMap[color]
}
