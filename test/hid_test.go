package hid

import (
	cp "github.com/UedaTakeyuki/compare"
	"local.packages/hid"
	"log"
	"testing"
)

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
