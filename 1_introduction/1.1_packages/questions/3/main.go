package main

import (
    "fmt"
    "os"
    "strconv"
)

func run() error {
    if len(os.Args) < 4 {
        return fmt.Errorf("not enough arguments: expected 3 arguments, got %d", len(os.Args)-1)
    }

    width, err := strconv.Atoi(os.Args[1])
    if err != nil {
        return fmt.Errorf("invalid width: %v", err)
    }

    height, err := strconv.Atoi(os.Args[2])
    if err != nil {
        return fmt.Errorf("invalid height: %v", err)
    }

    percent, err := strconv.Atoi(os.Args[3])
    if err != nil {
        return fmt.Errorf("invalid percentage: %v", err)
    }

    file, err := os.Create("config.txt") // Исправлено имя файла
    if err != nil {
        return fmt.Errorf("something wrong with file: %v", err)
    }

    defer file.Close()

    _, err = fmt.Fprintf(file, "%dx%d %d%%", width, height, percent) // Исправлено на %% для отображения процента
    if err != nil {
        return fmt.Errorf("could not write to config file: %v", err)
    }

    return nil
}

func main() {
    if err := run(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
