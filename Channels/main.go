package main

import(
	"log"
)

const num_pool = 1000

func Calculate_Value(intChan chan in t){
	Random_Number := helpers.Random_Number(num_pool)
	intChan <- Random_Number
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go Calculate_Value(intChan)

	num := <-intChan
	log.Println(num)
}