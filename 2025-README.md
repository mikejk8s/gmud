# Lets go again!

### 

### State: 

12/12/2024
Prior to bailing for 2 years I was converting the ORM to point to postgres instead of mysql and got caught in the weeds for some reason.

### WSL2 Specifics

* Use docker-compose.exe up to start to use Windows GUI desktop

# TODO: 

In no particular order here

* Upgrade go things
    * go version 19 -> latest
    * go modules
    * container images
    * reconnect swimm docs
* Security
    * Clean up image, /code etc are passing through


## To get the engine cranking again

1. main.go has an ssh key required to start the server 		wish.WithHostKeyPath(".ssh/term_info_ed25519"), for ease of use there's a key in this repo

