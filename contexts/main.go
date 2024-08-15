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
	c = context.WithValue(c, CtxKey{}, CtxValue{id: primitive.NewObjectID().Hex()})
	ctxWithTimeout, cancel := context.WithTimeout(c, time.Second*19)
	defer cancel()
	RunJob(ctxWithTimeout, cancel, jobs)

}

func RunJob(ctxWithTimeout context.Context, cancel context.CancelFunc, jobs []string) {

	for _, v := range jobs {
		select {
		case <-ctxWithTimeout.Done():
			log.Default().Println("Cancelled: ", v)
			return
		default:
			ctxWithTimeout.Value("alo")
			log.Default().Println("Run job: ", v, ctxWithTimeout.Value(CtxKey{}).(CtxValue).id)
			time.Sleep(time.Second * 20)
		}
	}
}
