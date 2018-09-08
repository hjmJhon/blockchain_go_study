package main

import "fmt"

//定义接口
type Car interface {
	Run()
}

type Phone interface {
	Call()
}

//定义 struct
type SmartCar struct {
	name string
}

//定义 struct
type Truck struct {

}

// struct 实现接口方法
func (smartCar *SmartCar) Run(){
	fmt.Println("smartCar Run method is running")
}

func (truck *Truck) Run()  {
	fmt.Println("truct Run method is running")
}

func (smartCar *SmartCar) Call() {
	fmt.Println("smartCar Call method is running")
}

func main() {
	var car Car
	car =  new(SmartCar)
	car.Run()

	car = new(Truck)
	car.Run()

	var phone Phone
	phone = new(SmartCar)
	phone.Call()
}
