package progress

// The easiest way to get a progressbar is as follows
func Example() {
	var p Progress
	p.GetBar(275, 346)
}

// Changing the Style is not hard, in fact it's quite easy
func Example_styled() {
	var p Progress
	p.SetStyle("parallelogram")
	p.GetBar(275, 346)
}

// Using a custom style is as simple as adding it and changing to it
func Example_customStyle() {
	var p Progress

	DefineStyle("my-style", []string{"○", "◔", "◑", "◕", "●"})
	p.SetStyle("my-style")
	p.GetBar(275, 346)
}
