# gmud

A mud in go
## How to run?
````shell
docker compose up
````
Change initdecoy.sql to init.sql and replace username and passwords in line 23-30.

And connect to 127.0.0.1:2222 with your favourite Telnet client.
It doesn't work on my side due to my college network, but it should work on yours.

## Ideas

- I am thinking something else for MP.
- Above seems like overengineering it best.
- If we can handle the multiplayer, then return to implementing the game again. Such as statistics.
- D&D 5e ruleset can be used for characters, and one-page adventures can be used for quests. Lobby system will be used for multiplayer.
- We may say just screw it, and let the people create their own adventures. But that will be a lot of work. And Path of Exile's new league is near. We can just roll with 3-5 quests and keep it as proof of concept not a game engine.

 
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

~~9. Docker-compose broken, gmud connection refused to mysql/localhost:3306 - can access w/ sqlstudio fine~~ Rename initdecoy.sql to init.sql and replace XXXXX values on line 25-30. Then write username, password, Host (in (127.0.0.1:3306) form) and password in docker.composer.yml


~~10. Alive reports false need to switch to reverse bool or dead~~  1 means alive and 0 means dead? Cant change it.

~~11. Level shouldnt be 0~~ It is level 1 on every new created character.

12. Create a stats model for character stats.

Currently after class selection, character is saved in database. Create a new stat selection model, expand characters table (or however you want) and send character to database after stat selection.

![image](https://user-images.githubusercontent.com/92731060/204088110-59cc9580-e76b-4a89-8fa1-2a949ccbcbbe.png)


14. Take barebones signup server (html templates in cmd/app/templates, create static files at wherever you want I didnt code it. backend code itself is in pkg/backend) and make it actually a real server. mike probably had a site for that.



## Api Paths

~~Gin Stats http://127.0.0.1:8080/stats
http://127.0.0.1:8081/characters/9 {id}~~

This part is completely yeeted off. We still need a create account endpoint on a site that creates new users at users table.



# To fix go mod  paths

```go
➜ charactersDb git:(main) ✗ go env -w GOPRIVATE=github.com/mikejk8s
➜  charactersDb git:(main) ✗ go mod tidy
Found existing alias for "go mod". You should use: "gom"
go: finding module for package github.com/mikejk8s/gmud
go: finding module for package github.com/go-sql-driver/mysql
go: downloading github.com/mikejk8s/gmud v0.0.0-20220821060920-758a6a03bc00
go: found github.com/go-sql-driver/mysql in github.com/go-sql-driver/mysql v1.6.0
go: found github.com/mikejk8s/gmud in github.com/mikejk8s/gmud v0.0.0-20220821060920-758a6a03bc00
```

These are working I guess? Didnt have any problem with fresh installation WSL.
