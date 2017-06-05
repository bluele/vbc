package vbc

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestVBC32(t *testing.T) {
	for i := 0; i < 1000; i++ {
		vs := getRandomUInt32Values(1000, 1000000)
		encoded, err := Encode32(vs)
		if err != nil {
			t.Error(err)
		}
		org := Decode32(encoded)
		if len(org) != len(vs) {
			t.Errorf("%v != %v", len(org), len(vs))
		}
		for i, v1 := range org {
			if v1 != vs[i] {
				t.Errorf("idx: %v, %v != %v", i, v1, vs[i])
			}
		}
	}
}

func TestVBC64(t *testing.T) {
	for i := 0; i < 1000; i++ {
		vs := getRandomUInt64Values(1000, 1000000)
		encoded, err := Encode64(vs)
		if err != nil {
			t.Error(err)
		}
		org := Decode64(encoded)
		if len(org) != len(vs) {
			t.Errorf("%v != %v", len(org), len(vs))
		}
		for i, v1 := range org {
			if v1 != vs[i] {
				t.Errorf("idx: %v, %v != %v", i, v1, vs[i])
			}
		}
	}
}

func getRandomUInt32Values(size, rng int) []uint32 {
	v := []uint32{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		v = append(v, uint32(rand.Intn(rng)+1))
	}
	sort.Slice(v, func(i int, j int) bool {
		return v[i] < v[j]
	})
	return v
}

func getRandomUInt64Values(size, rng int) []uint64 {
	v := []uint64{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		v = append(v, uint64(rand.Intn(rng)+1))
	}
	sort.Slice(v, func(i int, j int) bool {
		return v[i] < v[j]
	})
	return v
}
