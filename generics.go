package main

import "fmt"
// Starting with version 1.18, Go has added support for generics, also known as type parameters.

/*
func MapKeys[K comparable, V any][m map[K]V) []K {
    r := make([]K, 0, len(m))
    for k:= range m {
        r = append(r, k)
    }
    return r
}

type List[T any] struct {
}

type element[T any] struct {
}

func (lst *List[T]) Push(v T) {
}

func (lst *List[T]) GetAll() []T {
}

func main() {
    var m = map[int]String{1: "2", 2: "4", 4: "8"}
    fmt.Println("keys m:", MapKeys(m))

    _ = MapKeys[int, string](m)

    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.GetAll())
}
*/
