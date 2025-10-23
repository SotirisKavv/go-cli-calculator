# 🧮 CLI Calculator (Go) - **Beginner Level**

A simple command-line calculator demonstrating fundamental Go concepts like function design, error handling, and CLI argument parsing. Perfect for learning Go basics and mathematical operations.

---

## 🚀 What is this?

A basic Go CLI tool that performs six mathematical operations with proper error handling. Showcases fundamental Go programming concepts and standard library usage.

---

## ✨ Features

- **Six Mathematical Operations:** Add, subtract, multiply, divide, exponent, modulus
- **Error Handling:** Division by zero protection and input validation
- **CLI Interface:** Flag-based argument parsing with help system
- **Zero Dependencies:** Pure Go standard library implementation

---

## � Go Skills Demonstrated

- **Basic Syntax:** Functions, variables, control structures
- **Flag Package:** Command-line argument parsing
- **Error Handling:** `error` returns and validation
- **Math Package:** `math.Pow`, `math.Mod` functions
- **Switch Statements:** Operation routing logic
- **Standard Library:** `fmt`, `os`, `flag` packages

---

## 🛠️ Usage

```sh
# Basic operations
go run calculator.go -op add -a 10.5 -b 3.2        # Result: 13.7000
go run calculator.go -op divide -a 20 -b 4          # Result: 5.0000
go run calculator.go -op exponent -a 2 -b 8         # Result: 256.0000

# Error handling
go run calculator.go -op divide -a 10 -b 0          # Error: Division by zero
go run calculator.go -h                             # Show help
```

---

## 🎯 Learning Objectives

This project demonstrates:
- **Go Fundamentals:** Basic syntax, functions, and error handling
- **CLI Development:** Argument parsing and user interaction
- **Mathematical Operations:** Implementing arithmetic with validation
- **Code Organization:** Clean function structure and separation of concerns


