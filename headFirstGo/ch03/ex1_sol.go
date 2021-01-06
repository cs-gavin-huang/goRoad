/*
 * @Author: geekli
 * @Date: 2021-01-06 19:23:52
 * @LastEditTime: 2021-01-06 19:23:52
 * @LastEditors: your name
 * @Description: 
 * @FilePath: /goRoad/headFirstGo/ch03/ex1_sol.go
 */
 package main

 import "fmt"
 
 // Declare a scoreSummary function with four parameters.
 // "name" has a type of "string". "score1", "score2",
 // and "score3" all have a type of "float64".
 func scoreSummary(name string, score1 float64, score2 float64, score3 float64) {
	 // The parentheses ensure that the addition operations
	 // take place before dividing.
	 average := (score1 + score2 + score3) / 3.0
	 // %s formats a string value.
	 // %10s pads a string with spaces to a width of 10.
	 // %f formats a floating-point value.
	 // %8.2f pads a floating-point value with spaces to a
	 // width of 8, ensuring that a minimum of two decimal
	 // places are displayed.
	 // Don't forget the newline character at the end!
	 fmt.Printf("%10s | %8.2f | %8.2f | %8.2f | %8.2f\n",
		 name, score1, score2, score3, average)
 }
 
 func main() {
	 fmt.Printf("%10s | %8s | %8s | %8s | %8s\n",
		 "Name", "Score 1", "Score 2", "Score 3", "Average")
	 fmt.Println("------------------------------------------------------")
	 scoreSummary("Jermaine", 95.4, 82.3, 74.6)
	 scoreSummary("Jessie", 79.3, 99.1, 82.5)
	 scoreSummary("Lamar", 82.2, 95.4, 77.6)
 }