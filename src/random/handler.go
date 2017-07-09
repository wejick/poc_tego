package random

import (
	"log"
	"math/rand"
	"net/http"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	"time"

	"github.com/julienschmidt/httprouter"
	tegoHTTP "github.com/wejick/tego/http"
)

type RandomS struct{}

//GetRandomHTTP http handler for get random
func GetRandomHTTP(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewRandomClient(conn)

	r, err := c.Getrandom(context.Background(), &RandomNumberRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	randomNum := randomNumber{
		Number:  int(r.Number),
		TimeNow: r.Now,
	}
	tegoHTTP.ResponseJSON(w, randomNum, 200, "")
}

func getRandom() randomNumber {
	randomNum := randomNumber{
		Number:  rand.Int(),
		TimeNow: time.Now().String(),
	}

	return randomNum
}

//Getrandom grpc handler for get random
func (s *RandomS) Getrandom(ctx context.Context, req *RandomNumberRequest) (response *RandomNumberResponse, err error) {
	randomNum := getRandom()
	response = &RandomNumberResponse{
		Number: int32(randomNum.Number),
		Now:    randomNum.TimeNow,
	}
	return
}
