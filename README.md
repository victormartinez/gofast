# Resumo

- There are no Exceptions in Go
- Variables declared in global scope are visible globally

```
go run .
```

```go

fmt.Printf("Type of var is %T and the value is %v", e, e)

const myConstant = "Hello, World!"

type MyCustomType int

noNeedToSpecifyType := "X" 

var b bool
b = true

```


## Arrays

```go

var myArray [3]int
myArray[0] = 10

for i, v := range meuArray {
    fmt.Printf("Index %d with value %d\n", i, v)
}


slicedArray := []int{1000, 2000, 3000, 4000}
slicedArray = append(slicedArray, 5000)
fmt.Printf("len=%d capacity=%d %v\n", len(slicedArray), cap(slicedArray), slicedArray)

// capacidade = 8 apesar de 5 elementos.
// Go dobra o tamanho do slice sempre que precisa redimensionar. Por isso tente trabalhar
// com um slice próximo da capacidade que você vai trabalhar.

```
