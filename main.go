package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
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

	var x float32
	var y float32

	for f < float32(-1) || f > float32(1) {
		x = rand.Float32()
		y = rand.Float32()
		f = x / y
	}

	return f
}
