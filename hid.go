// hid - my supplimentals of hashids.
//
// Copyright 2020 Aterier UEDA
// Author: Takeyuki UEDA

package hid

import (
	"log"

	"github.com/speps/go-hashids"
)

type hidType hashids.HashID

/*
 * CreateHashID
 */
func CreateHashID(salt string, alphabet string, minlength int) (h *(hashids.HashID), err error) {
	// hashid
	hd := hashids.NewData()
	hd.Salt = salt
	hd.Alphabet = alphabet
	hd.MinLength = minlength
	h, err = hashids.NewWithData(hd)
	return
}

func CreateHID(salt string, alphabet string, minlength int) (h *hidType, err error) {
	// hashid
	var hid *(hashids.HashID)
	hid, err = CreateHashID(salt, alphabet, minlength)
	if err == nil {
		h = (*hidType)(hid)
	}
	return
}

/*
 * Encode
 */
func (h *hidType) Encode(i int) (str string, err error) {
	str, err = (*(hashids.HashID))(h).Encode([]int{i})
	return
}

func (h *hidType) EncodeInt64(i int64) (str string, err error) {
	str, err = (*(hashids.HashID))(h).EncodeInt64([]int64{i})
	return
}

func (h *hidType) Decode(str string) (i int, err error) {
	var res []int
	res, err = (*(hashids.HashID))(h).DecodeWithError(str)
	if err == nil {
		log.Println("res", res)
		i = res[0]
	}
	return
}

func (h *hidType) DecodeInt64(str string) (i int64, err error) {
	var res []int64
	res, err = (*(hashids.HashID))(h).DecodeInt64WithError(str)
	if err == nil {
		i = res[0]
	}
	return
}
