package main

import (
	"os"
	"time"

	"pault.ag/go/luxafor"
)

func vws(flag *luxafor.Flag) {
	for {
		for _, color := range []luxafor.Color{
			luxafor.Red,
			luxafor.Red,
			luxafor.Green,
		} {
			flag.SetColor(luxafor.Off)
			time.Sleep(time.Second / 6)
			flag.SetColor(color)
			time.Sleep(time.Second / 6)
		}
	}
}

func main() {
	flag, err := luxafor.OpenFlag()
	if err != nil {
		panic(err)
	}

	var color luxafor.Color
	switch os.Args[1] {
	case "red":
		color = luxafor.Red
	case "blue":
		color = luxafor.Blue
	case "green":
		color = luxafor.Green
	case "magenta":
		color = luxafor.Magenta
	case "yellow":
		color = luxafor.Yellow
	case "off":
		color = luxafor.Off
	case "vws":
		vws(flag)
		return
	default:
		panic("Unknown color")
	}

	if err := flag.SetColor(color); err != nil {
		panic(err)
	}
}
