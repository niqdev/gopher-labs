package myconcurrency

import (
	"fmt"
	"math/rand"
	"time"
)

// CHANNELS https://medium.com/codex/learn-how-golang-channels-work-by-building-them-72f49ed30f2c

type MyColor struct {
	name    string
	primary bool
}

func ColorFilter() {

	colorChan := make(chan MyColor)

	go findPrimary(colorChan)

	for color := range colorChan {
		color.printLabel()
	}
}

func findPrimary(colorChan chan MyColor) {
	for {
		color := randomColor()
		colorChan <- color
		if color.primary {
			break
		} else {
			time.Sleep(time.Second * 2)
		}
	}
	close(colorChan)
}

func randomColor() MyColor {
	colors := []MyColor{
		{name: "red", primary: true},
		{name: "green", primary: true},
		{name: "blue", primary: true},
		{name: "yellow", primary: false},
		{name: "brown", primary: false},
		{name: "purple", primary: false},
		{name: "white", primary: false},
		{name: "black", primary: false},
		{name: "orange", primary: false},
	}
	return colors[rand.Intn(len(colors))]
}

func (c *MyColor) label() string {
	if c.primary {
		return fmt.Sprintf("%s is primary", c.name)
	} else {
		return fmt.Sprintf("%s is not primary", c.name)
	}
}
func (c *MyColor) printLabel() {
	fmt.Println(c.label())
}
