// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// go test *.go -bench=".*" -benchmem

package guuid_test

import (
	"testing"

	"github.com/gogf/guuid/v2"
)

func Benchmark_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		guuid.New().String()
	}
}

func Benchmark_NewUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if u, err := guuid.NewUUID(); err != nil {
			panic(err)
		} else {
			u.String()
		}
	}
}

func Benchmark_NewDCEGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if u, err := guuid.NewDCEGroup(); err != nil {
			panic(err)
		} else {
			u.String()
		}
	}
}

func Benchmark_NewDCEPerson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if u, err := guuid.NewDCEPerson(); err != nil {
			panic(err)
		} else {
			u.String()
		}
	}
}

func Benchmark_NewRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if u, err := guuid.NewRandom(); err != nil {
			panic(err)
		} else {
			u.String()
		}
	}
}

func Benchmark_NewMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		guuid.NewMD5(guuid.UUID{}, []byte("")).String()
	}
}

func Benchmark_NewSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		guuid.NewSHA1(guuid.UUID{}, []byte("")).String()
	}
}
