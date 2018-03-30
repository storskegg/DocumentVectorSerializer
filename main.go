package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/eknkc/basex"
	"github.com/golang/protobuf/proto"
)

const (
	base36dict = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base62dict = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	baseXdict  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ,.;:|[]{}`~!@#$%^&*()-_=+?'\"\\"
)

func main() {
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

	base62, _ := basex.NewEncoding(base62dict)
	base36, _ := basex.NewEncoding(base36dict)
	baseX, _ := basex.NewEncoding(baseXdict)

	dVec64 := base64.URLEncoding.EncodeToString(dVecM)
	dVec62 := base62.Encode(dVecM)
	dVec36 := base36.Encode(dVecM)
	dVecX := baseX.Encode(dVecM)

	fmt.Printf("Serialized 36... (%v)\n\n", len(dVec36))
	fmt.Println(dVec36)
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("Serialized 62... (%v)\n\n", len(dVec62))
	fmt.Println(dVec64)
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("Serialized 64... (%v)\n\n", len(dVec64))
	fmt.Println(dVec64)
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("Serialized X... (%v)\n\n", len(dVecX))
	fmt.Println(dVecX)
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
