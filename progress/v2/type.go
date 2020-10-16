package progress

// Style acts as a container for the progressbar design
type Style struct {
	Bar            []string // The characters used to construct the bar
	AlternatingBar bool     // The bar is not incremental like `block` but like `doublesinglebar`

	Empty []string // The characters used to construct the unfilled part of the progressbar

	Head []string // The character replacing the last one of the bar, if multiple are given they will be alternated

	FrameStart string // A character to print before the bar
	FrameEnd   string // A character to print after the unfilled part

	BarStyle   func(char string) string // A function to apply to Bar characters
	HeadStyle  func(char string) string // A function to apply to Head characters
	EmptyStyle func(char string) string // A function to apply to Empty-ish(?) characters
}
