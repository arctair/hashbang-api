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
### Run the tests against a deployment
```
$ BASE_URL=https://hashbang.arctair.com go test
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
