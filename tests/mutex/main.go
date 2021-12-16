package main

import (
	"context"
	"log"
	"time"

	"git.pleamon.com/p/petcd"
)

func main() {
	ctx := context.TODO()
	etcd, err := petcd.New("192.168.1.220:2379")
	if err != nil {
		panic(err)
	}
	mutex, err := etcd.NewMutex("test")
	if err != nil {
		panic(err)
	}
	log.Println(mutex)
	err = mutex.Lock(ctx)
	if err != nil {
		panic(err)
	}
	// log.Println(4)
	log.Println("locked")
	<-time.After(10 * time.Second)
	mutex.Unlock(ctx)
	log.Println("unlock")
}
