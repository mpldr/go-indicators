package spinner

import (
	"fmt"
	"time"
)

// This function pretty much shows all you need for using a spinner:
//   1. create a spinner
//   2. set a style
//   3. use .Next()
func Example_SimpleSpinner() {
	var s Spinner
	s.SetStyle("dots")

	for i := 0; i < 100; i++ {
		fmt.Printf("\rworking%v", s.Next())
		time.Sleep(80 * time.Millisecond)
	}
}

// This shows the same spinner 3 times and cleans up afterwards
func Example_Multiuse() {
	var s1 Spinner
	s1.SetStyle("bouncing-bar")

	for i := 0; i < 100; i++ {
		fmt.Printf(" %v\tdo stuff\n %v\texecuting something else\n %v\twaiting for some seconds\r\033[F\033[F", s1.Next(), s1.Current(), s1.Current())
		time.Sleep(80 * time.Millisecond)
	}

	// removes the spinner-characters but leaves the text untouched
	fmt.Printf(" %v\n %v\n %v\n", s1.Clear(), s1.Clear(), s1.Clear())
}
