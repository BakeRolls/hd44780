## GoBot HD44780 Driver

This is a driver for Character LCDs based on the HD44780 using GoBot.

```go
func main() {
	r := raspi.NewAdaptor()
	lcd := NewHD44780Driver(r, "7", "12", "11", [8]string{"13", "16", "15", "18", "22", "29", "32", "31"})

	lcd.Display(true)
	lcd.Clear()
	lcd.Println("Hello")
	lcd.ShiftRight()
	lcd.Left()
	lcd.Print("World")
	lcd.Right()
	lcd.Print("---")
}
```
