/*
 * @Author: geekli
 * @Date: 2020-12-24 09:00:54
 * @LastEditTime: 2020-12-24 09:00:54
 * @LastEditors: your name
 * @Description: 
 * @FilePath: /goRoad/headFirstGo/ch02/ex2_sol.go
 */
 package main

 import (
	 "bufio"
	 "fmt"
	 "log"
	 "os"
	 "strconv"
	 "strings"
 )
 
 func main() {
	 reader := bufio.NewReader(os.Stdin)
 
	 fmt.Print("Enter racer name: ")
	 name, err := reader.ReadString('\n')
	 if err != nil {
		 log.Fatal(err)
	 }
	 name = strings.TrimSpace(name)
 
	 fmt.Print("Enter racer rank: ")
	 input, err := reader.ReadString('\n')
	 if err != nil {
		 log.Fatal(err)
	 }
	 input = strings.TrimSpace(input)
	 rank, err := strconv.ParseInt(input, 10, 64)
	 if err != nil {
		 log.Fatal(err)
	 }
 
	 // Need to declare "medal" outside the "if" statements,
	 // so that it's still in scope afterwards.
	 var medal string
	 if rank == 1 {
		 medal = "gold"
	 } else if rank == 2 {
		 medal = "silver"
	 } else if rank == 3 {
		 medal = "bronze"
	 } else { // Runs if none of the above are true.
		 medal = "participant"
	 }
 
	 fmt.Println(name, "gets a", medal, "medal!")
 }