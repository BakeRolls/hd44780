package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
)

// HD44780Driver is a gobot driver for modules based on the HD44780.
// Datasheet: https://www.beta-estore.com/download/rk/RK-10290_410.pdf
type HD44780Driver struct {
	pinRS    *gpio.DirectPinDriver
	pinRW    *gpio.DirectPinDriver
	pinE     *gpio.DirectPinDriver
	pinsData [8]*gpio.DirectPinDriver
	execTime time.Duration

	name       string
	connection gobot.Connection
	gobot.Commander
}

// NewHD44780Driver return a new HD44780Driver given a gobot.Connection, the
// RS, RW, E and data pins.
func NewHD44780Driver(a gobot.Connection, rs, rw, e string, data [8]string) *HD44780Driver {
	pinsData := [8]*gpio.DirectPinDriver{}
	for i, pin := range data {
		pinsData[i] = gpio.NewDirectPinDriver(a, pin)
	}

	return &HD44780Driver{
		name:       gobot.DefaultName("HD44780"),
		connection: a,
		Commander:  gobot.NewCommander(),
		pinRS:      gpio.NewDirectPinDriver(a, rs),
		pinRW:      gpio.NewDirectPinDriver(a, rw),
		pinE:       gpio.NewDirectPinDriver(a, e),
		pinsData:   pinsData,
		execTime:   2160 * time.Microsecond,
	}
}

// Halt implements the Driver interface.
func (h *HD44780Driver) Halt() (err error) { return }

// Name returns the HD44780Drivers name.
func (h *HD44780Driver) Name() string { return h.name }

// SetName sets the HD44780Drivers name.
func (h *HD44780Driver) SetName(n string) { h.name = n }

// Connection returns the HD44780Driver Connection.
func (h *HD44780Driver) Connection() gobot.Connection { return h.connection }

// Off turns all pins low.
func (h *HD44780Driver) Off() error {
	h.pinRS.Off()
	h.pinRW.Off()
	h.pinE.Off()
	for _, pin := range h.pinsData {
		pin.Off()
	}
	return nil
}

// Initialize turns on the display and enables the cursor.
func (h *HD44780Driver) Initialize(displayCursor bool) error {
	h.Off()
	h.pinE.On()
	if displayCursor {
		h.pinsData[0].On()
		h.pinsData[1].On()
	}
	h.pinsData[2].On()
	h.pinsData[3].On()
	h.pinE.Off()
	return nil
}

// Clear clears the display.
func (h *HD44780Driver) Clear() error {
	h.Off()
	h.pinE.On()
	h.pinsData[0].On()
	h.pinE.Off()
	return nil
}

// Home resets the cursor to the position of the first charactor on the first
// row.
func (h *HD44780Driver) Home() error {
	h.pinE.On()
	h.SendData(byte(0x02))
	h.pinE.Off()
	return nil
}

// RightToLeft enables writing from right to left.
func (h *HD44780Driver) RightToLeft() error {
	h.pinE.On()
	h.SendData(byte(0x04))
	h.pinE.Off()
	return nil
}

// LeftToRight enables writing from left to right.
func (h *HD44780Driver) LeftToRight() error {
	h.pinE.On()
	h.SendData(byte(0x06))
	h.pinE.Off()
	return nil
}

// Left moves the cursor left.
func (h *HD44780Driver) Left() error {
	h.SendData(byte(16))
	h.pinE.On()
	time.Sleep(h.execTime)
	h.pinE.Off()
	return nil
}

// Right moves the cursor right.
func (h *HD44780Driver) Right() error {
	h.pinE.On()
	h.SendData(byte(20))
	h.pinE.Off()
	return nil
}

// ShiftLeft moves the screen left.
func (h *HD44780Driver) ShiftLeft() error {
	h.SendData(byte(24))
	h.pinE.On()
	time.Sleep(h.execTime)
	h.pinE.Off()
	return nil
}

// ShiftRight moves the screen right.
func (h *HD44780Driver) ShiftRight() error {
	h.pinE.On()
	h.SendData(byte(28))
	h.pinE.Off()
	return nil
}

// SendData turns on the register select and sends a byte which might be an
// ASCII character.
func (h *HD44780Driver) SendData(data byte) error {
	for _, pin := range h.pinsData {
		if data&1 == 1 {
			pin.On()
		} else {
			pin.Off()
		}
		data >>= 1
	}
	return nil
}

// Print splits an ASCII string into bytes and sends them to the display.
func (h *HD44780Driver) Print(strs ...string) error {
	h.Off()
	h.pinRS.On()
	for _, str := range strs {
		for _, data := range str {
			h.pinE.On()
			h.SendData(byte(data))
			time.Sleep(h.execTime)
			h.pinE.Off()
		}
	}
	h.pinRS.Off()
	return nil
}

// Println fills the strings to the displays width and prints them.
func (h *HD44780Driver) Println(strs ...string) error {
	for i := range strs {
		strs[i] = fmt.Sprintf("%-20s", strs[i])
	}
	return h.Print(strs...)
}
