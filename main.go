package main

import (
	"log"
)

func area(length float64, width float64) float64 {

	result := length * width
	return result
}

func main() {

	length := 5.5
	width := 10

	result := area(length, float64(width))

	log.Println(result)

}
