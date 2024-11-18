package main

import (
	"io"
	"fmt"
	"bytes"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
    buffer := make([]byte, 1024) // Фиксированный размер буфера

    for n > 0 {
        toRead := len(buffer)
        if uint(toRead) > n {
            toRead = int(n)
        }

        // Читаем данные из r
        readBytes, err := r.Read(buffer[:toRead])
        if err != nil && err != io.EOF {
            return err
        }

        // Пишем данные в w
        writtenBytes, err := w.Write(buffer[:readBytes])
        if err != nil {
            return err
        }

        // Уменьшаем n на количество записанных байт
        n -= uint(writtenBytes)

        // Если достигнут конец данных, выходим
        if err == io.EOF || readBytes == 0 {
            break
        }
    }

    return nil
}

func main() {
    // Источник данных
    reader := bytes.NewReader([]byte("Hello, world!"))
    // Назначение для записи
    var writer bytes.Buffer

    // Копируем не более 5 байт
    err := Copy(reader, &writer, 5)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Проверяем результат
    fmt.Println("Result:", writer.String()) // Вывод: "Hello"
}
