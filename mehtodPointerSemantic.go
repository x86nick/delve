// Pointer Semantic
package main

import (
	"fmt"
	"time"
)

type car struct {
	year        int
	maker, model string
	color       string
}

func (c *car) isNew() bool {  // copy of pointer to car
	return c.year == time.Now().Year()
}
func (c *car) paint(color string) {
	c.color = color
}

func main() {
	fredsCar := car{
		year:  time.Now().Year(),
		maker:  "Ford",
		model: "Pinto",
		color: "Puke Green",
	}
	fmt.Println(fredsCar)
	fmt.Println("New?", fredsCar.isNew())
    fredsCar.paint("Navajo White")
    fmt.Println(fredsCar)
}
