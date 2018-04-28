# how to deploy this app using Docker

$ docker build -t go-web-server .
$ docker run -it --rm --name web-app-server go-web-server

# go dep package management tool
go dep => package management tool. dep init 

dep ensure => check changes and update the Gopkg.lock file