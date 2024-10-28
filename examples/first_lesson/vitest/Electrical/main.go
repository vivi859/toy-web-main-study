package main

import (
	"fmt"
)

func main() {
	carmelaa := Camera{}
	pann := Pan{}
	usbb := UsbWorker{}
	usbb.UsbWorker(&carmelaa)
	usbb.UsbWorker(&pann)
}

type Usb interface {
	start()
	stop()
}
type UsbWorker struct{}

// 多态
func (u *UsbWorker) UsbWorker(usb Usb) {
	usb.start()
	usb.stop()
}

type Camera struct {
}

func (c *Camera) start() {
	fmt.Println("Camera start")
}
func (c *Camera) stop() {
	fmt.Println("Camera stop")
}

type Pan struct {
}

func (p *Pan) start() {
	fmt.Println("Pan start")
}
func (p *Pan) stop() {
	fmt.Println("Pan stop")
}
