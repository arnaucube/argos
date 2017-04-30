package main

import "fmt"

//Color struct, defines the color
type Color struct{}

var c Color

//DarkGray color
func (c Color) DarkGray(t string) {
	fmt.Print("\x1b[30;1m") //dark gray
	fmt.Println(t)
	fmt.Print("\x1b[0m") //defaultColor
}

//Red color
func (c Color) Red(t string) {
	fmt.Print("\x1b[31;1m") //red
	fmt.Println(t)
	fmt.Print("\x1b[0m") //defaultColor
}

//Green color
func (c Color) Green(t string) {
	fmt.Print("\x1b[32;1m") //green
	fmt.Println(t)
	fmt.Print("\x1b[0m") //defaultColor
}

//Yellow color
func (c Color) Yellow(t string) {
	fmt.Print("\x1b[33;1m") //yellow
	fmt.Println(t)
	fmt.Print("\x1b[0m") //defaultColor
}

//Blue color
func (c Color) Blue(t string) {
	fmt.Print("\x1b[34;1m") //blue
	fmt.Println(t)
	fmt.Print("\x1b[0m") //defaultColor
}

//Purple color
func (c Color) Purple(t string) {
	fmt.Print("\x1b[35;1m") //purple
	fmt.Println(t)
	fmt.Print("\x1b[0m") //defaultColor
}

//Cyan color
func (c Color) Cyan(t string) {
	fmt.Print("\x1b[36;1m") //cyan
	fmt.Println(t)
	fmt.Print("\x1b[0m") //defaultColor
}
