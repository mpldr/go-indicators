package progress

// The easiest way to get a progressbar is as follows
func Example() {
	var p Progress
	p.GetBar(275, 346)
}

// The easiest way to get a progressbar is as follows
func Example_styled() {
	var p Progress
	p.SetStyle("parallelogram")
	p.GetBar(275, 346)
}

// The easiest way to get a progressbar is as follows
func Example_customStyle() {
	var p Progress

	DefineStyle("my-style", []string{"○", "◔", "◑", "◕", "●"})
	p.SetStyle("my-style")
	p.GetBar(275, 346)
}
