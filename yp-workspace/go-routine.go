package main

import "time"

func main(){
	go countdown()
	go count()
	time.Sleep(100 * time.Millisecond)
}

func countdown(){
	for i := 10; i>0; i--{
		println(i)
	}
}

func count(){
	for i := byte('a'); i<byte('z'); i++ {
		println(string(i))
	}
}