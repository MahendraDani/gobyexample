# Names
- If the name begins with an upper-case letter, it is exported, which means that it is visible and accessible outside of
its own package and may be referred to by other parts of the program, as with Printf in the fmt package. 
- Package names themselves are always in lower case
- Go programmers use camel case naming convetion. Ex - `parseRequestLine` over `parse_request_line`
- The letters of acronyms and initialisms like ASCII and HTML are always rendered in the same case, so a function might be called htmlEscape, HTMLEscape, or escapeHTML, but not escapeHtml.

# Variables
The variable declarations has the following form
```Go
var name type = expression
```
- if the type is omitted, it is determined by the initialzer expression.
- if the expression is omitted, the initial value is _zero value_ of the type:
    - 0 for numbers
    - false for booleans
    - "" for strings
    - nil for interfaces and reference types (slice, maps, pointer, channel, function)
    - the zero value of an aggregate type like an array or a struct has the zero value of all its elements or fields.
- In GO, there is no such thing as uninitialized variable.

# Short variable declarations
Short variable declarations has the following form
```
name := expression
```

- `:=` is a declaration, where as `=` is an assignment
- short variable declaration does not necessarily declare all the variables on its left-hand side. If some of them were already declared in the same lexical block, then the short variable declaration acts like an assignment to those variables
```Go
in, err := os.Open(infile)
// ...
out, err := os.Create(outfile)
```
Here, the first statement declares both `in` and `err`, but the second one only declares `out` and assigns a value to existing
variable `err`.

# Pointers
- A pointer value is the address of a variable. A pointer is thus the location at which a value is stored.
```Go
var x = 3 // &x is the "address of x"
*int p = &x / *int is "pointer to int"

fmt.Println(*p) // 3, prints value of the address of variable it holds
fmt.Println(p) // prints the address of x

*p = 4 // updates the value of x to 4
```

# The new() function
- it creates a unnamed variable of type T, initializes it to the zero value of T, and returns its address, which is a value
of type *T

```Go
p := new(int) // p, of type *int, points to an unnamed int variable
fmt.Println(*p) // 0

*p = 2 // sets unnamed int to 2
fmt.Println(*p) // 2

```
