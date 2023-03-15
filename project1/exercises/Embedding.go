package main

import "fmt"

func main() {
	var c IDraw = Circle{Name: "Circle1", privateValue: "private circle 1"}
	c.Draw()

	c.(ITransform).Transform()
	//var e IEat = c.(IEat)  // main.Circle is not main.IEat: missing method Eat
	//e.(ITransform).Transform()

	t1, ok := c.(IDraw)
	fmt.Println("-- IDraw(Circle) is IDraw?", ok, t1)
	t2, ok := c.(ITransform)
	fmt.Println("-- IDraw(Circle) is ITransform?", ok, t2)
	t3, ok := c.(IEat)
	fmt.Println("-- IDraw(Circle) is IEat?", ok, t3)

	t4, ok := c.(Shape)
	fmt.Println("-- IDraw(Circle) is Shape?", ok, t4)
	t5, ok := c.(Circle)
	fmt.Println("-- IDraw(Circle) is Circle?", ok, t5)
}

type IDraw interface {
	Draw()
}

type ITransform interface {
	Transform()
}

type IEat interface {
	Eat()
}

type Shape struct {
	Name         string
	privateValue string
}

func (s Shape) Draw() {
	fmt.Println("Draw a shape.")
}

func (s Shape) Transform() {
	fmt.Println("Transform a shape.")
}

type Circle struct {
	Name         string
	privateValue string
	Shape
}

func (c Circle) Draw() {
	fmt.Println("Draw a circle.")
}
