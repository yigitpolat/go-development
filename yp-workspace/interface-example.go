package main

import "fmt"

func main(){
	//constructorla yapılabilirdi
	ferr := new(Ferrari)
	ferr.Brand = "Ferrari"
	ferr.Color = "Red"
	ferr.Special = true
	fmt.Println(ferr.Information("test"))

}

type CarFace interface {
	Run() bool
	Stop() bool
	Information() string
}

type Car struct {
	Brand string
	Color string
	Speed int
}

type SpecialProduction struct {
	Special bool
}

// Composition
type Ferrari struct {
	Car
	SpecialProduction
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
 	return false
}

func (x *Ferrari) Information(a string) string {
	return "information" + a
}