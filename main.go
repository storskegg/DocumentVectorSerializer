package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"runtime"

	"github.com/golang/protobuf/proto"
	"github.com/martinlindhe/base36"
)

func main() {
	runtime.GOMAXPROCS(1)

	vec := make([]float32, 300)

	for i := 0; i < 300; i++ {
		if i%2 == 0 {
			vec[i] = randFloat32()
		} else {
			vec[i] = randFloat32() * float32(-1)
		}
	}

	dVec := &DocumentVector{
		Vector: vec,
	}

	dVecM, _ := proto.Marshal(dVec)

	dVec64 := base64.URLEncoding.EncodeToString(dVecM)
	dVec36 := base36.EncodeBytes(dVecM)

	fmt.Printf("Serialized 36... (%v)\n\n", len(dVec36))
	fmt.Println(dVec36)
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("Serialized 64... (%v)\n\n", len(dVec64))
	fmt.Println(dVec64)
}

func randFloat32() (f float32) {
	f = 50
	x := make([]byte, 1)
	y := make([]byte, 1)

	for f < float32(-1) || f > float32(1) {
		n, err := rand.Read(x)
		if err != nil || n != 1 {
			log.Panicln(err)
		}
		_, err = rand.Read(y)
		if err != nil || n != 1 {
			log.Panicln(err)
		}

		f = float32(x[0]) / float32(y[0])
	}

	return f
}
