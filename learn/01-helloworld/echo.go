/*
Date: 07/10/2025
CLI Arguments in Go, by example of UNIX echo command 

- The "os" package in standard library provides functions and values dealing with operating system in a platform independent fashion.
- Slices in Go:
  - Dynamic sized sequence 's',  of array elements, where individual elements can be accesses as s[i] and contigous 
  subsequence can be accessed as s[m:n].
  - len(s): Gives the number of elements in s.
  - s[m:n] - gives elements from mth index upto nth index(excluding the nth index).
  - s[:n] - gives elements from 0th index utp nth.
  - s[m:] - gives elements from mth index upto len(s).
*/

package main

import (
  "fmt"
  "os"
)

func main(){
  /*
  - Variables can be initialized at the time of declaration.
  - If not initialized, they are assigned 'Zero Values' of its type. Numericals are assigned 0, and strins are initialized with
    empty strings ("").
  */
  var s, sep string

  /*
  - Short variable declaration - ":=" is used, a statement that declares one or more variables and gives appropriate types base
    on initalizer values.
  - i++ is equivalent to i+=1 => i = i + 1. Theses are statments NOT expressions. And only postfix are allowed.
  - j = i++ is ILLEGAL.
  - --i or ++i are also ILLEGAL.
  - paraentheses are not used in for loops
  */

  /*
  - the += operator creates a new string after concatenation, hence s will be remove via the garbage collector.
  */
  for i:=1; i<len(os.Args); i++ {
    s += sep + os.Args[i] // string concatenation
    sep = " " 
  }
  /*
  For loops in go:

  1. Type I
    for initialization; condition; post {
      // zero or more statements
    }

    - intialization statement is optional. It's executed before the loop starts. If it's present it MUST be a simple 
    statement, i.e a short variable declaration, an increment, or assignment statement, or a function call.
    - condition is a boolean statement, that is evaluated before execution of each iteration.
    - post statement is executed after the body of the loop, then the condition is evaluted again.
    - the loop ends when condtiion becomes false.

  2. Type II - tradition while loop
    for condition {
      // zero or more statements
    }

    - if the condition is also ommitted it becomes a infinite loop
    for {
      // zero or more statements
    }

  3. Type III - range based
    for idx, elem = range data_structure {
      // zero or more statements
    }
    - range based loops provides index and element value at that index.
    - Go doesn't allow unused local variables, so use '_' for ignoring the index
  */

  fmt.Println(s)
}

