# Go by Example
This is my hands-on studying https://gobyexample.com/.

## 6. Switch
Q. What does `i interface{}` mean?
```go
whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
        fmt.Println("I'm a bool")
```

## 8. Slices
- Unlike arrays, slices are typed only by the elements they contain (not the number of elements)
- Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.

### Methods
Q. Does it copy rect even when we pass a pointer of rect?
```
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}
```

### Generics
Requires version 1.18.
https://gobyexample.com/generics
