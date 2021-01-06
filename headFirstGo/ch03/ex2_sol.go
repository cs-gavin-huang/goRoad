/*
 * @Author: geekli
 * @Date: 2021-01-06 19:24:33
 * @LastEditTime: 2021-01-06 19:24:33
 * @LastEditors: your name
 * @Description: 
 * @FilePath: /goRoad/headFirstGo/ch03/ex2_sol.go
 */
 package main

 import "fmt"
 
 func perimeter(length float64, width float64) float64 {
	 return length*2 + width*2
 }
 
 func main() {
	 var total float64
	 total += perimeter(8.2, 10)
	 total += perimeter(5, 5.4)
	 total += perimeter(6.2, 4.5)
	 fmt.Println("You'll need", total, "meters of fencing")
 }