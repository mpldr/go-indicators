package main

import (
	"fmt"
	"time"

	"./progress"
	"./spinner"
)

func main() {
	/*	basicexample()
		fmt.Println(strings.Repeat("=", 20))
		multispinner()
		fmt.Println(strings.Repeat("=", 20))
		repeatedspinner()
		fmt.Println(strings.Repeat("=", 20))
		preparationspinner()
	*/
	var p progress.Progress

	for i := 0; i < 350; i++ {
		fmt.Printf(" %v\r", p.GetBar(float64(i), 300))
		time.Sleep(50 * time.Millisecond)
	}
}

func basicexample() {
	var s spinner.Spinner
	s.SetStyle("dots-bounce")

	for i := 0; i < 100; i++ {
		fmt.Printf(" %v working\r", s.Next())
		time.Sleep(80 * time.Millisecond)
	}

	fmt.Printf(" %v\n", s.Clear())
}

func multispinner() {
	var s1, s2, s3 spinner.Spinner
	s1.SetStyle("line")
	s2.SetStyle("windows-10")
	s3.SetStyle("clock")

	for i := 0; i < 100; i++ {
		fmt.Printf(" %v\tloading protocols\n %v\tbooting bloated OS\n %v\tperforming checks\r\033[F\033[F", s1.Next(), s2.Next(), s3.Next())
		time.Sleep(80 * time.Millisecond)
	}

	fmt.Printf(" %v\n %v\n %v\n", s1.Clear(), s2.Clear(), s3.Clear())

}

func repeatedspinner() {
	var s1 spinner.Spinner
	s1.SetStyle("bouncing-bar")

	for i := 0; i < 100; i++ {
		fmt.Printf(" %v\tloading protocols\n %v\tbooting bloated OS\n %v\tperforming checks\r\033[F\033[F", s1.Next(), s1.Current(), s1.Current())
		time.Sleep(80 * time.Millisecond)
	}

	fmt.Printf(" %v\n %v\n %v\n", s1.Clear(), s1.Clear(), s1.Clear())

}

func preparationspinner() {
	var s spinner.Spinner
	s.SetStyle("bouncing-bar")

	for i := 0; i < 50; i++ {
		fmt.Printf(" %v preparing download...\r", s.Next())
		time.Sleep(80 * time.Millisecond)
	}

	for i := 0; i < 50; i++ {
		fmt.Printf(" [%3.f%%] downloading...\033[K\r", float64(i))
		time.Sleep(80 * time.Millisecond)
	}
	fmt.Print(" [done] completed\033[K\n")

}
