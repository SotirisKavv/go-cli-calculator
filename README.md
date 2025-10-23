# CLI Calculator — Simple math operations in Go

Beginner‑friendly CLI calculator demonstrating Go fundamentals: flag parsing, clean error handling, and arithmetic operations using only the standard library.

- Operations: add, subtract, multiply, divide, exponent, modulus
- CLI: flag‑based interface with built‑in help
- Safety: input validation and division‑by‑zero protection
- Dependencies: pure Go standard library

Quick links:
- Entrypoint: `calculator.go`


## What this project showcases

- Basic Go syntax: functions, variables, control structures
- Flag parsing with `flag`
- Error handling via `error` returns and validation
- Math helpers using `math.Pow` and simple modulus logic
- Switch‑driven operation routing


## High‑level flow

Components:
- CLI flags: `-op`, `-a`, `-b`
- Dispatcher: routes operation via `switch`
- Compute: performs calculation with validation
- Output: prints formatted result or error

Flow:
1) Parse flags and validate inputs
2) Dispatch to the appropriate math operation
3) Print result with fixed precision or exit with an error message


## Design tactics

- Minimal, readable functions per operation
- Centralized validation for unsafe cases (e.g., divide by zero)
- Standard library only: `flag`, `fmt`, `os`, `math`


## Run locally

Prereqs: Go 1.20+

```powershell
# Basic operations
go run calculator.go -op add -a 10.5 -b 3.2        # Result: 13.7000
go run calculator.go -op divide -a 20 -b 4          # Result: 5.0000
go run calculator.go -op exponent -a 2 -b 8         # Result: 256.0000

# Error handling
go run calculator.go -op divide -a 10 -b 0          # Error: Division by zero
go run calculator.go -h                             # Show help
```


## Folder map

- `calculator.go`: CLI flags, dispatcher, and math operations


## Next steps (ideas)

- Add integer‑only mode and big number support
- Add interactive REPL mode


## Author

IAmSotiris


