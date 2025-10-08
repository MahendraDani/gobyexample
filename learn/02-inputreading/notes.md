# Reading input in Go
Ways of reading input depends on:
1. Source: standard input, files or network streams
2. Granularity: whether input is needed line-by-line, word-by-word, or structured input
3. Performance: buffered vs unbuffered I/O

# Five ways of reading input:
## 1. Using `fmt` package - `fmt.Scan()`, `fmt.Scanf()`, `fmt.Scanln()`
Pros:
- Very simple for reading **small** and **structured** input.
- Good for interactive console programs.

Cons:
- **Blocking** and **unbuffered** - not suited for large input streams.
- No control over partial reads or tokenization.
- Returns error on unexpected input.

## fmt.Scan()
```Go
var name string
fmt.Scan(&name)
```
- Reads space separated tokens.
- Stops scanning at **whitespaces** (space, tab, newline).

Use case:
- Just need to take 1-5 values as input interactively

## fmt.Scanln()
```Go
var name string
var age int
fmt.Scanln(&name,&age)
```
- Reads inputs until it encounters a new line.
- It expects all inputs specified on the same line.
- If there's an extra input than what it expects, it throws an error.

Use case:
- Useful when you want to take a specific number of inputs on the same line.


## fmt.Scanf()
```Go
var name string
var age int
fmt.Scanf("%s %d",&name,&age)
```

- Let's the programmer specify the input format.
- If we want to read multiple words with space, until new line:
```Go
fmt.Scanf("%[^\n]",&name)
```


## 2. Using `bufio.Scanner`
Pros
- **Buffered** + efficient, line-by-line reading.
- Handles continous input smoothly.
- Automatically stops at EOF or errors.
- Perfect for stream processing.

Cons
- Each line is limited to 64KB by default (can be increased via `Buffer()`)
- Can only split by a single delimiter at a time (default: newline)

Use Case
- Reading lines of text, e.g logs, CSV rows, or user input logs.

## 3. Using `bufio.Reader`
Pros
- More flexible than `Scanner` - can read by bytes, runes or, until any delimiter.
- Allows custom parsing logic and manual buffering.
- No 64KB limit like Scanner.

Cons
- Slightly more verbose
- Programmer must handle newlines, trimming, etc manually.

Use case
- when fine-grained control is needed.

## 4. Using `os.Stdin.Read`
Pros
- Lowest level access - reads raw bytes.
- Fast and flexible, minimal overhead.

Cons
- Programmer must handle decoding, tokenization and newlines.
- Not ideal for text-based input.

## 5. From files or other sources
- scanner functions works with files too.


