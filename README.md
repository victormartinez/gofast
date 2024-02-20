# Resumo

- There are no Exceptions in Go
- Variables declared in global scope are visible globally

```
go env
go env | grep GOOS
go env | grep ARCH
go tool dist list
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

```go

// Precisa do go mod antes
go mod init gofast  // gofast é o nome do projeto
go get golang.org/x/exp/constraints
go get github.com/google/uuid
```


```go

go build -o gofast .

./gofast

GOOS=windows go build -o gofastwin .
GOOS=windows GOARC=arm go build -o gofastwin .

GOOS=linux go build -o gofastlinux .

GOOS=darwin go build -o gofastlinux .
```