/*
 * @Author: geekli
 * @Date: 2020-12-24 08:56:31
 * @LastEditTime: 2020-12-24 08:57:01
 * @LastEditors: your name
 * @Description: 
 * @FilePath: /goRoad/headFirstGo/ch01/ex2_sol.go
 */
 package main

 // Since we're only importing one package, we can use the
 // single-line format for "import".
 import "fmt"
 
 func main() {
	 // We can use short variable declarations for everything.
	 // Later words in multi-word variable names should be
	 // capitalized.
	 pebbleWeight := 0.1
	 rockWeight := 1.2
	 boulderWeight := 502.4
	 // Underscores are legal in variable names, but are not
	 // conventional.
	 // Take the underscore out and capitalize the second word.
	 totalWeight := pebbleWeight + rockWeight + boulderWeight
	 fmt.Println(totalWeight)
 }