package main

import (
	"fmt"
	"math"
	"errors"
	"log"
)

//finding area of a circle using function 
//checking if what was passed is equal to 0 or less
//then it will produce an error since negative number or zero
//are not allowed.
//returns an answer and error if any
func circle_Area(radius float64) (float64, error) {
	if radius <= 0 {
		err := errors.New("cannot not have a negative number or zero")
		return -1, err
	}
	var y float64 = 2
	result := math.Pi * math.Pow(radius,y)
	return result, nil

}

func main() {
	radius := 1

	//pass radius to get area 
	result,err := circle_Area(float64(radius))
	//if err, is not equal to nil then the error message will be shown
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

}