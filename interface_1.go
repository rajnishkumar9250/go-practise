package main

import (
	"fmt"
)

type Shape interface{
	Area() float64
	Perimeter() float64
}

type Rectangle struct{
	height float64
	width float64

}

func (rec Rectangle) Area() float64{
	return	rec.height * rec.width	

}

func (rec Rectangle) Perimeter() float64{

	return  2*(rec.height + rec.width)
}

func main(){
	var s Shape
	s = Rectangle{5.0, 4.0}
	r := Rectangle{6.0, 5.0}
	fmt.Println("area of rectange s is ", s.Area())
	fmt.Println("area of rectange r is ", r.Area())


}
