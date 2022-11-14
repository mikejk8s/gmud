# gmud

A mud in go

## Features

What's broken:
1. //TODO: SSH only takes one key and right now I'm getting errors accessing it
2. //TODO: gmud container cant connect to database
   1. This is NOT a problem when running ./app from localhost

## Todo

~~1. Let user create a new character~~ 

![TODO1](https://media4.giphy.com/media/DtcIXipywWrSlvXPrp/giphy.gif?cid=790b76115f4a0964390d82e9bc76ed9bd151e5d2ee43f9a3&rid=giphy.gif&ct=g)

1.5 Assign user to character

2. Let user login to an existing character

3. Put user into beginner room with a description

4. ~~DeleteUser currently 404s~~

5. CharacterRoutes GetUser returns the wrong list of characters, doesn't do its query

6. Require unique names sql queries Character.Name

7. Make a map of public keys

8. *AccountLogin function runs login page on the server*

9. Docker-compose broken, gmud connection refused to mysql/localhost:3306 - can access w/ sqlstudio fine

10. Alive reports false need to switch to reverse bool or dead

11. Level shouldnt be 0

12.


## Api Paths

Gin Stats http://127.0.0.1:8080/stats
http://127.0.0.1:8081/characters/9 {id}

``` go
	a := r.Group("/api")
	{
		a.POST("/token", controllers.GenerateToken)
		a.POST("/user/register", controllers.RegisterUser)
		r.GET("/characters", cr.GetCharacters)
		s := a.Group("/secured").Use(middlewares.Auth())
		{
			s.GET("/user", controllers.GetUser)
			s.POST("/token", controllers.GenerateToken)
			s.GET("/characters/:id", cr.GetCharacter)
			s.POST("/characters", cr.CreateCharacter)
			s.PUT("/characters/:id", cr.UpdateCharacters)
			s.DELETE("/characters/:id", cr.DeleteCharacter)
		}
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
