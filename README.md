# gmud

A mud in go

## Features

What's broken:
r~~1. //TODO: SSH only takes one key and right now I'm getting errors accessing itr~~ SSH username and password type of authentication is in charge.
2. //TODO: gmud container cant connect to database
   1. This is NOT a problem when running ./app from localhost

## Todo

~~1. Let user create a new character~~ 

~~1.5 Assign user to character~~

~~2. Let user login to an existing character~~

~~3. Put user into beginner room with a description~~

~~3.5 Put a viewport into the tutorial.go as a beginner room with a description, as every new created or level 1 character will fall into model in tutorial.go~~

4. ~~DeleteUser currently 404s~~

~~5. CharacterRoutes GetUser returns the wrong list of characters, doesn't do its query~~ Now returns array of Characters model.

~~6. Require unique names sql queries Character.Name~~ Queries against DB is done with account name.

~~7. Make a map of public keys~~ No more public keys, SSH username and password login.

8. ~~AccountLogin function runs login page on the server*~~

8.25 We need a whole ass sign up site that handles the registers. Data needs to be nuked at users table.

9. Docker-compose broken, gmud connection refused to mysql/localhost:3306 - can access w/ sqlstudio fine

~~10. Alive reports false need to switch to reverse bool or dead~~  1 means alive and 0 means dead? Cant change it.

~~11. Level shouldnt be 0~~ It is level 1 on every new created character.

12. Create a stats model for character stats.

13. Create a combat tutorial room model.


## Api Paths

~~Gin Stats http://127.0.0.1:8080/stats
http://127.0.0.1:8081/characters/9 {id}~~

This part is completely yeeted off. We still need a create account endpoint on a site that creates new users at users table.



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

These are working I guess? Didnt have any problem with fresh installation WSL.
