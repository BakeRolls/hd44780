package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	lcd := NewHD44780Driver(r, "7", "12", "11", [8]string{"13", "16", "15", "18", "22", "29", "32", "31"})

	work := func() {
		lcd.Initialize(false)
		lcd.Clear()
		lcd.Println("Hello World")
	}

	robot := gobot.NewRobot("LCD", []gobot.Connection{r}, lcd, work)

	robot.Start()
}
