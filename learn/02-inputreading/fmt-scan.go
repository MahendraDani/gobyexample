package main

import (
  "fmt"
)

func main(){
  var name string
  var age int

  fmt.Print("Enter name and age: ")
  /*
  - Does not read full line. Will break if you input two or more space separated words.
  */
  _, err := fmt.Scan(&name,&age)

  if err != nil {
    fmt.Println("Error: ", err)
    return
  }
  fmt.Printf("%s\t%d\n",name,age)
}
