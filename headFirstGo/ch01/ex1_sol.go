/*
 * @Author: geekli
 * @Date: 2020-12-24 08:56:17
 * @LastEditTime: 2020-12-24 08:56:17
 * @LastEditors: your name
 * @Description: 
 * @FilePath: /goRoad/headFirstGo/ch01/ex1_sol.go
 */

 package main

 import "fmt"
 
 func main() {
	 var count int = 20
	 unitWeight := 0.4
	 totalWeight := float64(count) * unitWeight
	 fmt.Println(count, "cans weigh", totalWeight, "kilograms")
 }