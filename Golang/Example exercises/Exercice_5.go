package main

import "fmt"

func main (){
	var selection int
	fmt.Print("1: Question 1\n2: Question 2\n3: Question 3\n4: Question 4\n5: Question 5\n6: Question 6\n7: Question 7\n")
	fmt.Print("Please select the question number : ")

	fmt.Scanln(&selection)

	switch selection {
	case 1:{
		var percentage float64
		fmt.Print("Please enter student percentage (0 to 100) : ")
		fmt.Scanln(&percentage)

		i:=5
		for i==5{
			if percentage < 0 || percentage > 100 {
				fmt.Print("Please enter student percentage (0 to 100) : ")
				fmt.Scanln(&percentage)
				i = 5
			}else{
				i = 1
			}

		}

		if percentage >= 90 && percentage <= 100{
			fmt.Println("Student grade : A")
		}else if percentage >= 80 && percentage <= 89{
			fmt.Println("Student grade : B")
		}else if percentage >= 70 && percentage <= 79{
			fmt.Println("Student grade : C")
		}else if percentage >= 60 && percentage <= 69{
			fmt.Println("Student grade : D")
		}else{
			fmt.Println("Student grade : F")
		}
	}
	case 2:{
		var year1 int
		var year2 int
		var month1 int
		var month2 int
		var day1 int
		var day2 int
		fmt.Println("Please enter first date")
		fmt.Print("Year : ")
		fmt.Scanln(&year1)
		fmt.Print("Month : ")
		fmt.Scanln(&month1)
		fmt.Print("Day : ")
		fmt.Scanln(&day1)
		fmt.Println("Please enter second date")
		fmt.Print("Year : ")
		fmt.Scanln(&year2)
		fmt.Print("Month : ")
		fmt.Scanln(&month2)
		fmt.Print("Day : ")
		fmt.Scanln(&day2)

		if year1 == year2{
			if month1 == month2{
				if day1 == day2{
					fmt.Println("Both dates are same")
				}else if day1 > day2{
					fmt.Println("The first date does not come before second")
				}else if day1 < day2{
					fmt.Println("The first date comes before second")
				}
			}else if month1 > month2{
				fmt.Println("The first date does not come before second")
			}else if month1 < month2{
				fmt.Println("The first date comes before second")
			}

		}else if year1 > year2{
			fmt.Println("The first date does not come before second")
		}else if year1 < year2{
			fmt.Println("The first date comes before second")
		}
	}
	case 3:{
		var side1 float64
		var side2 float64
		var side3 float64
		fmt.Println("Please enter lengths of the sides")
		fmt.Print("Side1 : ")
		fmt.Scanln(&side1)
		fmt.Print("Side2 : ")
		fmt.Scanln(&side2)
		fmt.Print("Side3 : ")
		fmt.Scanln(&side3)

		if side1 == side2 && side1 == side3 {
			fmt.Println("Given sides are of an equilateral triangle (three equal sides)")
		}else if side1 == side2 || side2 == side3 || side1 == side3{
			fmt.Println("Given sides are of an isosceles triangle (two equal sides)")
		}else{
			fmt.Println("Given sides are of an scalene triangle (three different sides)")
		}

	}

	case 4:{
		var oldPremium float64
		var newPremium float64
		var noOfAccidents int
		fmt.Print("Please enter old premium value : ")
		fmt.Scanln(&oldPremium)
		fmt.Print("Please enter number of accidents occurred : ")
		fmt.Scanln(&noOfAccidents)

		if noOfAccidents == 0{
			newPremium = oldPremium * 1.02
		}else if noOfAccidents == 1 || noOfAccidents == 2{
			newPremium = oldPremium * 1.05
		}else if noOfAccidents == 3{
			newPremium = oldPremium * 1.10
		}else if noOfAccidents >= 4{
			newPremium = oldPremium * 1.30
		}
		fmt.Printf("The new premium value : %f\n",newPremium)
	}
	case 5:{
		var a [6]int
		fmt.Println("Please enter values into the array")
		highest := -9999
		lowest := 99999
		finalScore := 0.0
		for i:= 0; i < 6; i++{
			fmt.Scanln(&a[i])
		}
		for i:= 0; i < 6; i++{
			finalScore += float64(a[i])
			if highest < a[i] {
				highest = a[i]
			}
			if lowest>a[i] {
				lowest = a[i]
			}
		}
		finalScore = (finalScore - float64(highest + lowest))/4

		fmt.Printf("The final score : %f\n",finalScore)
	}
	case 6:{
		var year int
		var month int
		var day int
		var newYear int
		var newMonth int
		var newDay int
		fmt.Println("Please enter date")
		fmt.Print("Year : ")
		fmt.Scanln(&year)
		fmt.Print("Month : ")
		fmt.Scanln(&month)
		fmt.Print("Day : ")
		fmt.Scanln(&day)
		noOfDays := getNoDays(month)

		if day == noOfDays {
			newDay = 1
			if month == 12 {
				newMonth = 1
				newYear = year + 1

			}else{
				newMonth = month + 1
				newYear = year
			}
		}else{
			newDay = day + 1
			newMonth = month
			newYear = year
		}

		fmt.Printf("The next date : %d - %d - %d\n", newYear,newMonth,newDay)
	}
	case 7:{
		var year int
		fmt.Println("Please enter year to know whether it is a leap year or not")
		fmt.Print("Year : ")
		fmt.Scanln(&year)

		if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0{
			fmt.Printf("The year %d is a leap year", year)
		} else{
			fmt.Printf("The year %d is not a leap year", year)
		}

	}



	default:
		fmt.Println("Invalid entry!!!")
	}

}

func getNoDays(month int) int{
	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		return 31
	}else if month == 2{
		return 28
	}else{
		return 30
	}
}
