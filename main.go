package main

import (
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	lcd := NewHD44780Driver(r, "7", "12", "11", [8]string{"13", "16", "15", "18", "22", "29", "32", "31"})

	lcd.Initialize(true)
	lcd.Clear()
	lcd.Println("Hello")
	lcd.ShiftRight()
	lcd.Left()
	lcd.Print("World")
}
