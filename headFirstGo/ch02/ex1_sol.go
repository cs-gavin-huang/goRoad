/*
 * @Author: geekli
 * @Date: 2020-12-24 09:00:00
 * @LastEditTime: 2020-12-24 09:00:00
 * @LastEditors: your name
 * @Description: 
 * @FilePath: /goRoad/headFirstGo/ch02/ex1_sol.go
 */
 package main

 import (
	 "fmt"
	 "strings"
 )
 
 func main() {
	 // The "builder" variable holds a "strings.Builder"
	 // value. We're not assigning a value to "builder",
	 // but that's okay; you can call methods on the
	 // zero value for the "strings.Builder" type.
	 var builder strings.Builder
	 
	 // YOUR CODE HERE: Call the "WriteRune" method on
	 // "builder" 3 times, with the runes 'a', 'b', and 'c'.
	 builder.WriteRune('a')
	 builder.WriteRune('b')
	 builder.WriteRune('c')
	 
	 // YOUR CODE HERE: Call the "WriteString" method on
	 // "builder" once, with the string "defg".
	 builder.WriteString("defg")
	 
	 // YOUR CODE HERE: Call the "String" method on
	 // "builder", and pass the result to "fmt.Println".
	 fmt.Println(builder.String())
 }