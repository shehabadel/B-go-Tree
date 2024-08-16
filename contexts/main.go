package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type CtxKey struct {
}

type CtxValue struct {
	id string
}
type Response struct {
	val int
	err error
}

func main() {
	// jobs := []string{"1", "2", "3"}
	// c := context.Background()
	// c1 := context.WithValue(c, CtxKey{}, CtxValue{id: primitive.NewObjectID().Hex()})
	// c, cancelTimeout := context.WithTimeout(c1, time.Second*6)
	// defer cancelTimeout()
	// c, cancel := context.WithTimeout(c, time.Second*4)
	// defer cancel()
	// go RunJob(c, cancel, jobs)
	// time.Sleep(time.Second * 6)
	// log.Default().Println("finished")
	ctx := context.Background()

	val, err := fetchUserData(ctx, 6)
	if err != nil {
		log.Default().Panicln("error ", err)
	}
	log.Default().Println(val)

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

func fetchUserData(ctx context.Context, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	respch := make(chan Response)

	go func() {
		val, err := fetchThirdParty()
		respch <- Response{
			val: val,
			err: err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("third party timeout")

		case resp := <-respch:
			return resp.val, resp.err
		}
	}
}

func fetchThirdParty() (int, error) {
	time.Sleep(time.Millisecond * 600)
	return 666, nil
}
