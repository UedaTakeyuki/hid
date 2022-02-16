package hid

import (
	cp "github.com/UedaTakeyuki/compare"
	"local.packages/hid"
	"log"
	"testing"
)

// Test of hid itself
func Test_01(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	salt := string("For hid test")
	alphabet := string("abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNOPQRSTUVWXYZ") // 50 chars, no "l" nor "I"
	h, err := hid.CreateHashID(salt, alphabet, 8)
	if err != nil {
		log.Println(err)
	}
	var id string
	id, err = h.EncodeInt64([]int64{1})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("id", id)
	}
	cp.Compare(t, id, string("nLBjDMRp"))

	var i []int64
	i, err = h.DecodeInt64WithError(id)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("i", i)
	}
	cp.Compare(t, i[0], int64(1))
}

// Test for attached member functions of hid
func Test_02(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	salt := string("For hid test")
	alphabet := string("abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNOPQRSTUVWXYZ") // 50 chars, no "l" nor "I"
	h, err := hid.CreateHID(salt, alphabet, 8)

	var originali, actuali int = 1, 0
	var originalii, actualii int64 = 1, 0
	var str string

	str, err = h.Encode(originali)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("str", str)
	}

	actuali, err = h.Decode(str)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("actuali", actuali)
	}
	cp.Compare(t, actuali, originali)

	str, err = h.EncodeInt64(originalii)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("str", str)
	}

	actualii, err = h.DecodeInt64(str)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("actualii", actualii)
	}
	cp.Compare(t, actualii, originalii)
}
