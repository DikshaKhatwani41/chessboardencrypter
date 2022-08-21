package main

import (
	"fmt"
	"math"
	"strconv"
)

func main(){
//	Q44()
//	Q45()
//	Q46()
	Q47()
}
func Q44(){
	var amount float64
	var duration float64
	var interestRate float64
	amount = 500000
	duration = 5
	interestRate = 10

	fmt.Println("Amount after 5 years: "+strconv.FormatFloat(compoundInterest(amount,duration,interestRate),'f',2,32))

	fmt.Println("Enter Amount: ")
	fmt.Scanln(&amount)
	fmt.Println("Enter duration: ")
	fmt.Scanln(&duration)
	fmt.Println("Enter interest Rate: ")
	fmt.Scanln(&interestRate)

	fmt.Println("Amount after "+strconv.FormatFloat(duration,'f',2,32)+" years: "+strconv.FormatFloat(compoundInterest(amount,duration,interestRate),'f',2,32))
}
func compoundInterest(amount float64, duration float64, interestRate float64) float64{
	return amount*math.Pow((1+interestRate/100),duration)
}
func Q45(){
	value1 := uint64(0)
	value2 := uint64(1)
	for i:=0 ;i<100;i++{
		newVal := value1 + value2
		fmt.Println(strconv.FormatUint(newVal,10) + " ")
		value1 = value2
		value2  = newVal
	}
}
func Q46(){
	for i:=2;i<50000;i++{
		if(isPrime(i)){
			fmt.Println(strconv.Itoa(i))
		}
	}
}
func isPrime(num int) bool{
	for i:=2;i<=num/2;i++{
		if(num%i==0){
			return false;
		}
	}
	return true
}
func Q47(){
	var noOfStudents int
	fmt.Println("Enter no of Students: ")
	fmt.Scanln(&noOfStudents)
	sumOfMarks:=0.0
	var w1 float64
	var w2 float64
	var w3 float64
	var w4 float64
	for {
		fmt.Println("Enter Weight 1: ")
		fmt.Scanln(&w1)
		fmt.Println("Enter Weight 2: ")
		fmt.Scanln(&w2)
		fmt.Println("Enter Weight 3: ")
		fmt.Scanln(&w3)
		fmt.Println("Enter Weight 4: ")
		fmt.Scanln(&w4)
		if (w1+w2+w3+w4 == 100) {
			break
		} else {
			fmt.Println("Sum of weights should be 100\n\n\n\n")
		}
	}
	for i:=0;i<noOfStudents;i++{
		var g1 float64
		var g2 float64
		var g3 float64
		var g4 float64
		for{
			fmt.Println("Enter Grade 1: ")
			fmt.Scanln(&g1)
			if(g1<=100 && g1>=0){
				break
			}else{
				fmt.Println("Grade Should be bw 0 & 100\n\n")
			}
		}
		for{
			fmt.Println("Enter Grade 2: ")
			fmt.Scanln(&g2)
			if(g2<=100 && g2>=0){
				break
			}else{
				fmt.Println("Grade Should be bw 0 & 100\n\n")
			}
		}
		for{
			fmt.Println("Enter Grade 3: ")
			fmt.Scanln(&g3)
			if(g3<=100 && g3>=0){
				break
			}else{
				fmt.Println("Grade Should be bw 0 & 100\n\n")
			}
		}
		for{
			fmt.Println("Enter Grade 4: ")
			fmt.Scanln(&g4)
			if(g4<=100 && g4>=0){
				break
			}else{
				fmt.Println("Grade Should be bw 0 & 100\n\n")
			}
		}
		finalMarks := g1*w1/100+g2*w2/100+g3*w3/100+g4*w4/100
		sumOfMarks+=finalMarks
		if(finalMarks>=60){
			fmt.Println("Pass")
		}else{
			fmt.Println("Fail")
		}
	}
	fmt.Println("\nAverage Marks: "+strconv.FormatFloat(sumOfMarks/float64(noOfStudents),'f',2,32))
}
