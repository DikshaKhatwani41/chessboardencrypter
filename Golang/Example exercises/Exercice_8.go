package main

import (
	"fmt"
	"strconv"
)

/* Exercise1
Determine what will be displayed when each of the following lines is executed. Suppose that x = 2 and y = 3.
       a) Write x
       b) Write x + x
       c) Write “x =”
       d) Write “x * x”
       e) Write “x * y”, y + x
       f) Write x * y + y + x
       g) Write “x * y + y + x”
*/

/* Exercise2
Write a program that asks the user to enter two numbers, and that displays the sum, the product, the difference,
and the quotient of these two numbers.
 */

/* Exercise3
Write a program that asks the user to enter two integers, and that displays the larger number, followed by the words “
is greater than ”, followed by the smaller number. However, if the numbers are equal, the program should display
“These numbers are equal.”
 */

/* Exercise4
Write a program that receives three integers entered via the keyboard, and that displays the sum, the average,
the product, the smallest, and the largest of these numbers.
 */

/* Exercise5
Write a program that reads the radius of a circle, and that displays the diameter, the circumference,
and the area of the circle. Use the constant value 3.14159 for π.
 */

/* Exercise6
Drivers are concerned with the mileage of their automobiles. A driver decides to record the number
of kilometers traveled and the number of litres of gasoline used, each time they refill their gas tank.
Develop a program with the goal of being able to enter the number of kilometers traveled and the number of litres
used upon each gas refill. The program should calculate and display the rate of gas consumption
(in litres per 100 kilometers) between each gas refill. After having processed all of the information entered,
the program should calculate and display the total rate of gas consumption (in litres per 100 kilometers)
for all of the gas refills.
 */

/* Exercise7
A large chemical products company compensates its commercial representatives by commission.
The representatives receive $200 per week plus 9% of their gross sales per week. For example, a representative
who sells $5000 of chemical products in one week receives a salary of $200 plus 9% of $5000, for a total of $650.
Develop a program that asks for the gross weekly sales of each representative and that calculates and displays their salary.
Process the information of one representative at a time.
*/

func main() {

	fmt.Println("\n\n**** Exercise 1  *****")
	x := 2
	y := 3
	fmt.Println(x) 						//2
	fmt.Println(x+x) 					//4
	fmt.Println("x =") 				//x =
	fmt.Println("x * x") 			//x * x
	fmt.Println("x * y", x + y ) 	//x * y 5
	fmt.Println(x * y + y + x ) 		//11
	fmt.Println("x * y + y + x") 	//x * y + y + x

	fmt.Println("\n\n**** Exercise 2  *****")
	var num1, num2 float32
	fmt.Print("Enter the first number: ")
	fmt.Scanf("%f", &num1)
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%f", &num2)
	println("The results are: ")
	Calculator(num1, num2)

	fmt.Println("\n\n**** Exercise 3  *****")
	var number1, number2 int
	fmt.Print("Enter the first number: ")
	fmt.Scanf("%d", &number1)
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%d", &number2)
	Comparative(number1, number2)

	fmt.Println("\n\n**** Exercise 4  *****")
	var nu1, nu2, nu3 int
	fmt.Print("Enter the first number: ")
	fmt.Scanf("%d", &nu1)
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%d", &nu2)
	fmt.Print("Enter the third number: ")
	fmt.Scanf("%d", &nu3)
	Comparative1(nu1, nu2, nu3)

	fmt.Println("\n\n**** Exercise 5  *****")
	var radius float32
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)
	Circumference(radius)

	fmt.Println("\n\n**** Exercise 6  *****")
	var litUsed, kmTraveled, sumlitUsed, sumKmTraveled float32
	sumlitUsed = 0
	sumKmTraveled = 0
	for {
		fmt.Print("Enter the number of litres used (-1 to terminate): ")
		fmt.Scanf("%f", &litUsed)
		if litUsed != -1 {
			fmt.Print("Enter the number of kilometers traveled: ")
			fmt.Scanf("%f", &kmTraveled)
			fmt.Println("The rate of gas consumption in litres per 100 kilometers for this gas refill is ", (litUsed / kmTraveled) *100)
			sumlitUsed = sumlitUsed + litUsed
			sumKmTraveled = sumKmTraveled + kmTraveled
		} else {
			fmt.Println("\nThe total rate of gas consumption in litres per 100 kilometers is ", (sumlitUsed / sumKmTraveled) * 100)
			return
		}
	}


	fmt.Println("\n\n**** Exercise 7  *****")
	const salary float64 = 200
	const commission float64 = 0.09
	var sell, valueCommission, totalSalary float64
	for {
		fmt.Print("Enter the representative’s sales in dollars (-1 to terminate): ")
		fmt.Scanf("%f", &sell)
		if sell != -1 {
			valueCommission = sell * commission
			totalSalary = salary + valueCommission
			fmt.Println("The representative’s salary is : ", fmt.Sprintf("%.4f", totalSalary))
		} else {
			return
		}
	}
}

func Circumference (rad float32) {
	const pi = 3.14159
	var diameter, area, circumferenc float32
	diameter = rad * 2
	circumferenc = diameter * pi
	area = rad * rad * pi
	fmt.Println("The diameter is: ", fmt.Sprintf("%.2f", diameter),
				", the circumference is: ",	fmt.Sprintf("%.2f", circumferenc),
				", the area is: ",	fmt.Sprintf("%.2f", area))
}

func Calculator (n1 float32, n2 float32) {
	fmt.Println("The sum is: ", fmt.Sprintf("%0.2f", (n1 + n2)))
	fmt.Println("The difference is: ", fmt.Sprintf("%0.2f", n1 - n2))
	fmt.Println("The product is: ", fmt.Sprintf("%0.2f",n1 * n2))
	fmt.Println("The quotient is: ", fmt.Sprintf("%0.2f",n1 / n2))
}

func Comparative (n1 int, n2 int){
	if n1 > n2 {
		println(strconv.Itoa(n1) + " is greater than " +  strconv.Itoa(n2))
	} else if n2 > n1 {
		println(strconv.Itoa(n2) + " is greater than " +  strconv.Itoa(n1))
	} else {
		println("These numbers are equal")
	}
}

func Comparative1 (n1 int, n2 int, n3 int){
	fmt.Println("The sum is: ", fmt.Sprintf("%d", (n1 + n2 + n3)))
	fmt.Println("The average is: ", fmt.Sprintf("%d", (n1 + n2 + n3)/3))
	fmt.Println("The product is: ", fmt.Sprintf("%d",n1 * n2 * n3))
	LargestNum(n1, n2, n3)
	SmallestNum(n1, n2, n3)
}

func LargestNum(n1 int, n2 int, n3 int) {
	if n1 > n2 {
		if n1 > n3 {
			fmt.Println("The largest number is: " + strconv.Itoa(n1))
		} else {
			fmt.Println("The largest number is: " + strconv.Itoa(n3))
		}
	} else if n2 > n3 {
		fmt.Println("The largest number is: " + strconv.Itoa(n2))
	} else {
		fmt.Println("The largest number is: " + strconv.Itoa(n3))
	}
}

func SmallestNum(n1 int, n2 int, n3 int) {
	if n1 < n2 {
		if n1 < n3 {
			fmt.Println("The smallest number is: " + strconv.Itoa(n1))
		} else {
			fmt.Println("The smallest number is: " + strconv.Itoa(n3))
		}
	} else if n2 < n3 {
		fmt.Println("The smallest number is: " + strconv.Itoa(n2))
	} else {
		fmt.Println("The smallest number is: " + strconv.Itoa(n3))
	}
}