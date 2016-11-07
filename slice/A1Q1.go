package main

import (
"fmt"
)

func frequency_count(userInput int) []int{
digits :=make([]int,10)
i:=userInput

for i>0{

digits[i%10]++

i=i/10
}

return digits
}

func main(){

digits :=make([]int,10)
var userInput int

fmt.Println("Enter the series of digits")
fmt.Scanln(&userInput)

frequency:=frequency_count(userInput)
fmt.Println("Frequency of each digit in the user Input is :")
fmt.Println(frequency)


}


