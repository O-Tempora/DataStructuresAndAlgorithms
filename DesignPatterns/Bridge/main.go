package main

import "fmt"

func main() {
	var (
		bwf = BWFormatter{Fmt: Backwarder{}}
		tbf = TBFormatter{Fmt: Tabber{trimSpaces: true}}
		str = "Lorem ipsum dolor sit amet"
	)

	PrintFormatted(str, bwf)
	PrintFormatted(str, tbf)
}

func PrintFormatted(s string, fm Formatter) {
	fmt.Println(fm.Format(s))
}
