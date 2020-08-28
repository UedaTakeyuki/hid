// hid - my supplimentals of hashids.
//
// Copyright 2020 Aterier UEDA
// Author: Takeyuki UEDA

package hid

import (
	"github.com/speps/go-hashids"
)

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
