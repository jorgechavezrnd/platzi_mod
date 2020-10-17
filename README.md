### Command for init go mod
```
go mod init platzi/go/curso/jorgechavezrnd
```

- Then `go.mod` file will be created
- Now we can add dependecies on go.mod, example:
```
require github.com/jorgechavezrnd/mycalculator v1.0.4
```
- Now we just execute `go run main.go` and dependencies will be downloaded, and `go.sum` file will be created
