# gmud

A mud in go

## Features

## Todo

1. Let user create a new character
2. Let user login to an existing character
3. Put user into beginner room with a description
4. ~~DeleteUser currently 404s~~
5. CharacterRoutes GetUser returns the wrong list of characters, doesn't do its query
6. Require unique names sql queries Character.Name
7. Make a map of public keys
8. AccountLogin function runs login page on the server 


## Api Paths

http://127.0.0.1:8080/characters/9 {id}

``` go
    r.GET("/characters", GetCharacters)
	r.GET("/characters/:id", GetCharacter)
	r.POST("/characters", CreateCharacter)
	r.PUT("/characters/:id", UpdateCharacters)
	r.DELETE("/characters/:id", DeleteCharacter)
	r.Run(":8080")
```



# To fix go mod  paths

```go
➜  charactersDb git:(main) ✗ go env -w GOPRIVATE=github.com/mikejk8s
➜  charactersDb git:(main) ✗ go mod tidy
Found existing alias for "go mod". You should use: "gom"
go: finding module for package github.com/mikejk8s/gmud
go: finding module for package github.com/go-sql-driver/mysql
go: downloading github.com/mikejk8s/gmud v0.0.0-20220821060920-758a6a03bc00
go: found github.com/go-sql-driver/mysql in github.com/go-sql-driver/mysql v1.6.0
go: found github.com/mikejk8s/gmud in github.com/mikejk8s/gmud v0.0.0-20220821060920-758a6a03bc00
```