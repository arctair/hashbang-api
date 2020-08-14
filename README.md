# hashbang
Manage Instagram drafts and their tags.  
https://hashbang.arctair.com  
## Run the tests
```
$ go test
```
or
```
$ nodemon
```
## Run the server
```
$ go run .
$ curl localhost:5000
```
## Build a docker image
```
$ go build -o bin/hashbang
$ docker build -t arctair/hashbang .
```