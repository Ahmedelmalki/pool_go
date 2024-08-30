package main

import (
    "fmt"
)

type Maybe[T any] struct {
    value *T
}

func Just[T any](value T) Maybe[T] {
    return Maybe[T]{value: &value}
}

func Nothing[T any]() Maybe[T] {
    return Maybe[T]{value: nil}
}

func (m Maybe[T]) IsPresent() bool {
    return m.value != nil
}

func (m Maybe[T]) Get() (T, bool) {
    if m.value == nil {
        var zero T
        return zero, false
    }
    return *m.value, true
}

func (m Maybe[T]) Bind[U any](f func(T) Maybe[U]) Maybe[U] {
    if !m.IsPresent() {
        return Nothing[U]()
    }
    return f(*m.value)
}

func addOne(x int) Maybe[int] {
    return Just(x + 1)
}

func main() {
    // Create a Just Maybe
    a := Just(5)

    // Use Bind to chain operations
    b := a.Bind(addOne).Bind(addOne).Bind(addOne)

    // Check the result
    if b.IsPresent() {
        value, _ := b.Get()
        fmt.Println("Value:", value) // Output: Value: 8
    } else {
        fmt.Println("No value")
    }

    // Create a Nothing Maybe
    c := Nothing[int]()

    // Use Bind with Nothing
    d := c.Bind(addOne).Bind(addOne).Bind(addOne)

    // Check the result
    if d.IsPresent() {
        value, _ := d.Get()
        fmt.Println("Value:", value)
    } else {
        fmt.Println("No value") // Output: No value
    }
}