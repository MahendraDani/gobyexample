# Why Go?
- Borrows concepts from multiple langugaes including C.
- compiled programming lanugage for writing robust, reliable and efficient code.
- Facilitates concurrency and provides flexibility in object-oriented programming.
- Has automatic memory management (garbage collection).

# Where to use Go?
- network servers, tools, systems programs, graphics, mobile applications and machine learning.
- Has become a popular replacement for untyped scripting languages.
- Runs faster than programs written in dynamic languages and suffer far few crashes due to unexpected type errors.

# Go's Origin (Relations with C)
- It's called as "C for 21st century"
- Go inherited its expression syntax, control-flow statements, basic data types, call-by value parameter passing, pointers 
and C's emphasis on programs that compile to efficient machine code and cooperate naturally with abstractions of current OS.
- Communicating Sequential Processes (CSP):  In CSP, a program is a parallel composition of processes that have no shared state;
the pro cesses communic ate and synchronize using channels.

# The Go Project
- Go brings garbage collection, a package system, first-class functions, lexical scope, a system call interface, and immutable
strings in which text is generally UTF-8 encoded.
- Go DOES NOT bring:
    - No implicit numeric conversions
    - No constructors or deconstructors
    - No operator overloading
    - No default parameter values
    - No inheritance
    - No generics
    - No exceptions
    - No macros
    - No function annotations
    - No thread-local storage
- Go's builtin data types work naturally without explicit initialization or implicit constructors, so relatively fewer memory alocations and memory writes are hidden in code.
- Go has concurrency features based on CSP. Go's lightweight threads variable stacks size are small enough to create one is cheap and creating a million is practical.

# The Go Standard Library
- Provides builing blocks and APIs for:
    - I/O
    - text processing
    - graphics
    - cryptography
    - networking
    - distributed applications

- Go does not require semicolons at the end of the statements or declarations, except where two or more appear on the same line.

Q. When to used short variable declaration and when to use explicit data-types?
- short variable declaration, can only be used within function, not for package-level variables1
