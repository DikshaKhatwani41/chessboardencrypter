package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
//	Q24()
//	Q25()
//	Q26()
//	Q27()
}
func Q24(){
	fmt.Println("Enter result of Coin 1: (H/T) ")
	var c1 string
	fmt.Scanln(&c1)

	fmt.Println("Enter result of Coin 2: (H/T) ")
	var c2 string
	fmt.Scanln(&c2)
	if(strings.ToUpper(c1)=="H"){
		if(strings.ToUpper(c2)=="H"){
			fmt.Println("Player wins 10$")
		}else{
			fmt.Println("Player wins 5$")
		}
	}else if(strings.ToUpper(c1)=="T"){
		fmt.Println("player looses")
	}else{
		fmt.Println("Invalid Input")
	}

}
func Q25(){
	fmt.Println("Enter number a")
	var a int
	fmt.Scanln(&a)

	fmt.Println("Enter number b")
	var b int
	fmt.Scanln(&b)

	fmt.Println("Enter number c")
	var c int
	fmt.Scanln(&c)

	if(a>b){
		if(a>c){
			fmt.Println("Greatest: "+ strconv.Itoa(a))
		}else{
			fmt.Println("Greatest: "+ strconv.Itoa(c))
		}
	}else{
		if(b>c){
			fmt.Println("Greatest: "+ strconv.Itoa(b))
		}else{
			fmt.Println("Greatest: "+ strconv.Itoa(c))
		}

	}
}
func Q26(){
	fmt.Println("Enter number a")
	var a int
	fmt.Scanln(&a)

	fmt.Println("Enter number b")
	var b int
	fmt.Scanln(&b)

	fmt.Println("Enter number c")
	var c int
	fmt.Scanln(&c)

	if(a<b && a<c){
		fmt.Print(strconv.Itoa(a)+" ")
		if(b<c){
			fmt.Print(strconv.Itoa(b)+" ")
			fmt.Print(strconv.Itoa(c)+" ")
		}else{
			fmt.Print(strconv.Itoa(c)+" ")
			fmt.Print(strconv.Itoa(b)+" ")
		}
	}else 	if(b<a && b<c){
		fmt.Print(strconv.Itoa(b)+" ")
		if(a<c){
			fmt.Print(strconv.Itoa(a)+" ")
			fmt.Print(strconv.Itoa(c)+" ")
		}else{
			fmt.Print(strconv.Itoa(c)+" ")
			fmt.Print(strconv.Itoa(a)+" ")
		}
	}else 	if(c<b && c<a){
		fmt.Print(strconv.Itoa(c)+" ")
		if(a<b){
			fmt.Print(strconv.Itoa(a)+" ")
			fmt.Print(strconv.Itoa(b)+" ")
		}else{
			fmt.Print(strconv.Itoa(b)+" ")
			fmt.Print(strconv.Itoa(a)+" ")
		}
	}



}
func Q27(){
	fmt.Println("Enter Salary:")
	var salary float64
	fmt.Scanln(&salary)

	fmt.Println("Are you married?(y/n)")
	var isMarried string
	fmt.Scanln(&isMarried)

	fmt.Println("No of children:")
	var noOfChildren int
	fmt.Scanln(&noOfChildren)

	fmt.Println("Have you arrived recently?(y/n)")
	var arrivedRecently string
	fmt.Scanln(&arrivedRecently)

	taxCut:=float64(noOfChildren)*0.5
	if(strings.ToLower(isMarried)=="y"){
		taxCut+=2
	}
	if(strings.ToLower(arrivedRecently)=="y"){
		taxCut+=8
	}
	taxAmount:=0.0
	if(salary>=60000.01){
		taxAmount = 1800+6400+18000+(salary-60000.00)*.4
	}else if(salary>=32000.01){
		taxAmount  =1800+6400+(salary-32000.00)*.3
	}else if(salary>=18000.01){
		taxAmount  = 1800+ (salary-18000.00)*.2
	}else{
		taxAmount = salary*.1
	}
	taxAmount = taxAmount *(100.00-taxCut)/100
	fmt.Println("Tax: "+strconv.FormatFloat(taxAmount,'f',2,32))
}
