/*
Date: 07/10/2025

- Go is a compiled language.
- go run __.go __.go ... => compiles the source code from one or more .go files, links it with libraries, and runs resulting executable file.
- Packages
  - Go code is organized into packages, which is similar to libraries or modules in other languages. 
  - Consists of one or more .go files in a single directory that defines what the package does.
- Each source file begins with a package declaration.
*/

/*
- "main" is a special package. It defines a standalone executable program, NOT a library.
*/
package main

/*
- Import declarations are used to tell the compiler, which packages are needed to execute the program.
- The "fmt" package contains functions for formatted output and scanning input. 
*/
import "fmt"

/*
- "main" is a special function - it's where the execution of the program begins.
- the "{" bracket must be on the same line as the function declaration.
*/
func main(){
  fmt.Println("Hello World\n")
}
