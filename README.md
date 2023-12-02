# Calcit2

[![Go](https://github.com/Qinbeans/Calcit2/actions/workflows/go.yml/badge.svg)](https://github.com/Qinbeans/Calcit2/actions/workflows/go.yml)

## About
Grammar is loosely based off of Calcit, but due to some limitations of Golang, cannot be exactly the same. The main difference between Calcit and Calcit2 is that Calcit2 uses ANTLR4 and Wails, while Calcit used Tauri with a custom parser. The idea for Calcit2 was to help me understand how to use ANTLR4.

## Usage
### Build
```bash
wails build
```
### Dev
```bash
wails dev
```
### Test
```bash
go test .
```

## Grammar
`pcast` is one of the differences between Calcit and Calcit2. Calcit implicitly bit casted floats to ints while Calcit2 requires the user to explicitly bit cast floats to ints. This is mostly a limitation from ANTLR4 that I came across and I didn't want to sacrifice readability for a feature that probably could benefit from being explicit.

```
pcast(12.5) >> 32
```

Aside from that, I added more explicitness to the grammar. For example, `1 + 2.5` is not valid in Calcit2, but `1 + int(2.5)` and `float(1) + 2.5` is. The UI for Calcit2 was simplified slightly by removing the drop down menu for the type of the variable. Instead, the user can just use `int()` or `float()` to cast the variable to the desired type. Casts for octal, binary, and hexadecimal are also supported, but it doesn't change the type of the variable. For example, `int(0b101)` is still an int, but it's value is 5, and the parser will see it as such. This effectively means that all octal, binary, and hexadecimal numbers are treated as ints. What separates them from regular ints is that they are displayed in their respective bases--the left most number retains the display base. For example, `1 + 0b101` will display `6` instead of `0b110`. This is because the left most number is a regular int, and the right most number is a binary int.

## Disclaimers
There's likely a perceivable performance difference between Calcit and Calcit2, but I haven't tested it. I would take a guess that Calcit is faster.