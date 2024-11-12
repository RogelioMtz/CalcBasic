# CalcBasics

This is a simple command-line calculator application written in Go. It supports basic arithmetic operations as well as some advanced operations like exponential and square root calculations.

## Features

- Basic Operations:
  - Addition (`+`)
  - Subtraction (`-`)
  - Multiplication (`*`)
  - Division (`/`)
- Advanced Operations:
  - Exponential (`^`)
  - Square Root (`v`)
- Quit the application (`q`)

## Usage

1. Clone the repository:
    ```sh
    git clone https://github.com/RogelioMtz/CalcBasics.git
    cd CalcBasics
    ```

2. Build the application:
    ```sh
    go build -o calcBasics
    ```

3. Run the application:
    ```sh
    ./calcBasics
    ```

4. Follow the on-screen instructions to perform calculations.

## Code Overview

### Main Function

The `main` function runs an infinite loop to continuously prompt the user for an operation and perform the corresponding calculation.

```go
func main() {
    for {
        operation := getOperation()
        if operation == Quit {
            fmt.Println("Quitting the program.")
            break
        }
        x, y := getUserInput(operation)
        result, err := performOperation(x, y, operation)
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Println("Result:", result)
        }
    }
}

## License
This project is licensed under the Mozilla Public License 2.0.
