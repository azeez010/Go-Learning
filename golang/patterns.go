package main
import "fmt"

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}


type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

func main (){
	m := new(Multiplication)
	add := Operation{Addition{}}
	// 8
	
	fmt.Println(add.Operate(3, 5))
	
	mult := Operation{Multiplication{}}
	
	
	fmt.Println(mult.Operate(3, 5), m.Apply(34, 2)) // 15

	t := Dec(Add)
	t(1, 4)
	t(29, 11)
}

type Ops func(a, b int) int;

func Dec(af Ops) Ops {
	return func(a, b int) int {
		fmt.Println("This is December,")
		re := af(a, b)
		fmt.Println("Na december be this. ", re)
		return re
	}
}

func Add(a, b int) int {
	return a + b
}