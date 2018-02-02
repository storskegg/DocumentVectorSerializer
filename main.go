package main

import (
	"crypto/rand"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/martinlindhe/base36"
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

	//fmt.Println("[]float32 Slice:")
	//fmt.Println(vec)

	dVec := &DocumentVector{
		Vector: vec,
	}

	dVecM, _ := proto.Marshal(dVec)
	//fmt.Println("Marshal'd Vector...")
	//fmt.Println(dVecM)

	dVecS := base36.EncodeBytes(dVecM)
	fmt.Println("Serialized...")
	fmt.Println(dVecS)
}

func randFloat32() (f float32) {
	f = 50
	x := make([]byte, 1)
	y := make([]byte, 1)

	for f < float32(-1) || f > float32(1) {
		rand.Read(x)
		rand.Read(y)

		f = float32(x[0]) / float32(y[0])
	}

	return f
}
