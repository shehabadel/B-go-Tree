package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CtxKey struct {
}

type CtxValue struct {
	id string
}

func main() {
	jobs := []string{"1", "2", "3"}
	c := context.Background()
	c1 := context.WithValue(c, CtxKey{}, CtxValue{id: primitive.NewObjectID().Hex()})
	c, cancelTimeout := context.WithTimeout(c1, time.Second*6)
	defer cancelTimeout()
	ctxWithTimeout, cancel := context.WithCancel(c)
	defer cancel()
	go RunJob(ctxWithTimeout, cancel, jobs)
	time.Sleep(time.Second * 6)
	log.Default().Println("finished")

}

func RunJob(ctxWithTimeout context.Context, cancel context.CancelFunc, jobs []string) {

	for i, v := range jobs {
		select {
		case <-ctxWithTimeout.Done():
			log.Default().Println("Cancelled: ", v, ctxWithTimeout.Err().Error())
			return
		default:
			if i == 1 {
				cancel()
			}
			log.Default().Println("Run job: ", v, ctxWithTimeout.Value(CtxKey{}).(CtxValue).id)
			time.Sleep(time.Second * 5)
		}
	}
}
