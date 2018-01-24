#!/bin/sh

#export GOROOT=/usr/local/go
 
 go get 'github.com/graphql-go/graphql'
 go get 'gopkg.in/alexcesaro/statsd.v2'
 go get 'github.com/sirupsen/logrus'
 go get 'github.com/hashicorp/consul/api'
 go get -u 'github.com/go-redis/redis'