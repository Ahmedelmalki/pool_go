package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)


/*
 * Complete the 'findSubstring' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func findSubstring(s string, k int32) string {
    maxsubstring := ""
    maxv := 0
    for i :=0 ; i<len(s)-int(k) ; i++{
        substring := s[i:i+int(k)]
        vcount := countv(substring)
        if vcount > maxv {
            maxv = vcount
            maxsubstring = substring
        }
    }
    return maxsubstring
}

func countv(str string) int{
    c := 0
    for i := 0 ; i < len(str) ; i++{
        if isvowel(string(str[i])){
            c++
        }
    }
    return c
}

func isvowel(str string) bool {
    r := []rune(str)
    for i := 0 ; i <len(r); i++{
        if r[i] == 'a' || r[i] == 'e' || r[i] == 'i' || r[i] == 'o' || r[i] == 'u'{
            return true
        }
    }
    return false
}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := findSubstring(s, k)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}