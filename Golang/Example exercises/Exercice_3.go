package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
//	Q11()
//	Q12()
//	Q13()
//	Q14()
//	Q15()
//	Q16()
//	Q17()
//	Q18()
	Q19()
}

func Q11() {
	fmt.Println("Enter number 1: ")
	var num1 int
	fmt.Scanln(&num1)

	fmt.Println("Enter number 2: ")
	var num2 int
	fmt.Scanln(&num2)

	fmt.Print("Greatest: ")
	if(num1>num2){
		fmt.Println(strconv.Itoa(num1))
	}else {
		fmt.Println(strconv.Itoa(num2))
	}

}
func Q12(){
	tip:=0.0
	fmt.Println("Enter Amount of Bill: ")
	var amount float64
	fmt.Scanln(&amount)
	if(amount>=1){
		tip = amount*.15
	}
	fmt.Println("Tip: "+strconv.FormatFloat(tip,'f',2,32))
}
func Q13(){
	fmt.Println("No of diskettes: ")
	var noOfDsk int
	fmt.Scanln(&noOfDsk)
	var totalPrice float64
	if(noOfDsk>25){
		totalPrice = 25.0 + (float64(noOfDsk)-25.0)*0.7
	}else{
		totalPrice = float64(noOfDsk)*1.0
	}
	fmt.Println("Are you member of Club Z ? (y/n)");
	var isMember string
	fmt.Scanln(&isMember)
	if(strings.ToLower(isMember)=="y"){
		totalPrice = totalPrice*.98
	}
	fmt.Println("Total Price: "+strconv.FormatFloat(totalPrice,'f',2,32))
}
func Q14(){
	fmt.Println("No of prints: ")
	var noOfPrint int
	fmt.Scanln(&noOfPrint)
	var totalPrice float64
	if(noOfPrint>100){
		totalPrice = 5 + (float64(noOfPrint)-100.0)*0.03
	}else{
		totalPrice = float64(noOfPrint)*0.05
	}
	fmt.Println("Total Price: "+strconv.FormatFloat(totalPrice,'f',2,32))
}
func Q15(){
	fmt.Println("Enter current Balance: ")
	var currentBal float64
	fmt.Scanln(&currentBal)

	fmt.Println("Enter Amount of Withdrawal: ")
	var withdrawalAmt float64
	fmt.Scanln(&withdrawalAmt)

	if(withdrawalAmt>currentBal){
		fmt.Println("Insufficient Balance")
	}else{
		fmt.Println("New Balance: "+strconv.FormatFloat(currentBal-withdrawalAmt,'f',2,32))
	}
}
func Q16(){
	fmt.Println("Enter no of Days: ")
	var noOfDays int
	fmt.Scanln(&noOfDays)

	fmt.Println("Enter no of kWh Consumed: ")
	var noOfkWh int
	fmt.Scanln(&noOfkWh)

	totalAmount:=float64(noOfDays)*0.5
	if(noOfkWh>200){
		totalAmount = totalAmount + 200*0.3  + float64(noOfkWh-200)*0.2
	}else{
		totalAmount = totalAmount + float64(noOfkWh) *0.3
	}
	fmt.Println("Total Amount of Bill : "+ strconv.FormatFloat(totalAmount,'f',2,32))


}
func Q17(){
	fmt.Println("Enter number: ")
	var num int
	fmt.Scanln(&num)
	if(num%2==0){
		fmt.Println("Even")
	}else{
		fmt.Println("Odd")
	}
}
func Q18(){
	fmt.Println("Enter number m")
	var m int
	fmt.Scanln(&m)

	fmt.Println("Enter number n")
	var n int
	fmt.Scanln(&n)
	if(m%n==0){
		fmt.Println("m is multiple of n")
	}else{
		fmt.Println("m is not a multiple of n")
	}
}
func Q19(){
	fmt.Println("Enter number a")
	var a int
	fmt.Scanln(&a)

	fmt.Println("Enter number b")
	var b int
	fmt.Scanln(&b)

	fmt.Println("Enter number c")
	var c int
	fmt.Scanln(&c)

	if(a == b+c){
		fmt.Println(a)
	}else if(b == a+c){
		fmt.Println(b)
	}else if(c == b+a){
		fmt.Println(c)
	}else{
		fmt.Println("No Solution")
	}

}