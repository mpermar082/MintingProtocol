// cmd/mintingprotocol/main.go
package main

import (
    "flag"
    "log"
    "os"

    "mintingprotocol/internal/mintingprotocol"
)

func main() {
    // Define flags with descriptions for better CLI usage
    verbose := flag.Bool("verbose", false, "Enable verbose logging")
    input := flag.String("input", "", "Input file path (required)")
    output := flag.String("output", "", "Output file path (required)")
    flag.Parse()

    // Validate input and output flags are provided
    if *input == "" || *output == "" {
        log.Fatal("Input and output file paths are required")
    }

    app := mintingprotocol.NewApp(*verbose)
    if err := app.Run(*input, *output); err != nil {
        log.Printf("Error: %v", err)
        os.Exit(1)
    }
}