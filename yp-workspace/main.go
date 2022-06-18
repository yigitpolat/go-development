package main

import (
	"fmt"
	"strconv"
)

type Human struct{
	FirstName string
	LastName string
	Age int
}
// Constructor
func newHuman(FirstName string) *Human {
	h := new(Human)
	h.FirstName = "yigit"
	return h
}


// Global Scope
var name string = "golang" //private
var Version string = "1.2.3" //public
const CONSTANT = "Constant"

var(
	firstVariable = "firstVariable"
	secondVariable = "secondVariable"
)
// -----------------
// -----------------


// Constant & ENUM
type Brand string
const (
	FACEBOOK Brand = "Facebook"
	GOOGLE Brand = "Google"
)

func printBrand(brand Brand){
	fmt.Println(brand)
}
// -----------------
// -----------------


func main() {
	//Variables
	var message1 string
	message1 = "message1"
	fmt.Println(message1)

	var message2 = "message2"
	fmt.Println(message2)
	
	var a,b,c int = 3,4,5
	var x,y,z = 1, true, "hello"
	fmt.Println(a,b,c)
	fmt.Println(x,y,z)

	k := 'A' // char !!
	l,_ := "abc", true
	fmt.Println(k,l)

	printBrand(GOOGLE)
	// -----------------
	// -----------------


	// Console Input - Output
	// Input
	// reader := bufio.NewReader(os.Stdin)
	stringLength, _ := fmt.Println("stringLength")
	fmt.Println(stringLength)
	fmt.Printf("This is how placeholders can be used %v, %T", stringLength, stringLength)
	// -----------------
	// -----------------


	// String conversion
	str := "1"
	number, _ := strconv.Atoi(str) // tersi Itoa - blank identifier
	fmt.Println(number)
	// -----------------
	// -----------------


	// Casting
	// küçükten büyüğe atama yapılabilir
	var i int = 5
	var j float64 =  float64(i)
	fmt.Println(j)
	// -----------------
	// -----------------


	// Conditions
	if 5 > 4 {
		fmt.Println("Greater")
	} else {
		fmt.Println("Smaller")
	}

	if foo := 1; foo == 1 {
		fmt.Println("equal")
	}
	// switch

	// for i := 1; i<10; i++ {}
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
	// continue - break

	// there is no while

	// -----------------
	// -----------------


	// Array - Range 
	myArray :=[5]string{"a", "b","c","d","e"}
	var pow = [...]int{1, 2, 3, 4, 5} 
	for i,v := range pow {
		fmt.Print(i, "-", v, "(", pow[i], "),")
	}

	// Slice
	// A slice does not store any data, it just describes a section of an underlying array.
	// more common
	var s []string = myArray[1:3] // capacity of s slice is 4
	fmt.Println(s)				  // works with pointers
	var stringSlice = []string{"1", "2", "3", "4", "5"}
	stringSlice = append(stringSlice, "6")
	fmt.Println(stringSlice)
	makeSlice := make([]int, 5, 5)
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))

	r := []struct {
		h int
		l bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(r)

	// Map
	// key - value pair
	myMap := make(map[int]string)
	myMap[34] = "Istanbul"
	myMap[35] = "Izmir"
	fmt.Println(myMap)


	// Functions
	fmt.Println(add(3,5))

	message := "Hello"
	changeMessage(&message)
	fmt.Println(message)

	surname, name := swap("Yigit", "Polat")
	fmt.Println(surname,name)

	fmt.Println(addd(1,2,3,4,5))

	// anonymous function
	addFunc := func(a,b int) (x int) {
		return a + b
	}
	fmt.Println(addFunc(10,10))

	// defer
	defer fmt.Println("En son çalışacak method bu")
	// -----------------
	// -----------------

	// Struct
	human := Human{FirstName: "yigit"}
	fmt.Println(human)
	// x := new(Human)




	// Exception Handling
	// myError := errors.New("Error")
	// fmt.Println(myError)
	// fmt.Errorf("Error")
	// file, err := os.Open("test.txt")
	// fmt.Println(err.Error)
	// os.Exit(1)
	// log.fatal("Error")
	// -----------------
	// -----------------
}

func add(x, y int) int {
	return x + y
}

func changeMessage(str *string){
	*str = "Hi"
}

func swap(x,y string) (string, string){
	return y,x
}

func addd(terms ...int) (int,int){
	result := 0
	for _,term := range terms {
		result += term
	}
	return len(terms), result
}

/*
func coolFunction(sum int) (x,y int) {
	x = 10
	y = 20
	return
}
 */



