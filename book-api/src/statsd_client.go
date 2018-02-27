package main

import (
	"log"
	"os"
	//"strconv"
	"gopkg.in/alexcesaro/statsd.v2"
)

func get_addr()(address string){
	var statsd_server string =  os.Getenv("STATSD_SERVER")
	var statsd_port string =  os.Getenv("STATSD_PORT")
	statsd_address := statsd_server + ":" + statsd_port
	return statsd_address
}

func statsd_incr(t string){

	statsd_address := get_addr()
	c, err := statsd.New(
		statsd.Network("tcp"),
		statsd.Address(statsd_address)
	)
	if err != nil {
		log.Print(err)
		//fmt.Println("%v",err)
	}
	defer c.Close()

	c.Increment(t)
}

func statsd_gauge(t string, v float64){
	statsd_address := get_addr()
	c, err := statsd.New(
		statsd.Network("tcp"),
		statsd.Address(statsd_address)
	)
	if err != nil {
		log.Print(err)
		//fmt.Println("%v",err)
	}
	defer c.Close()
	c.Gauge(t, v)
}


// Increment a counter.