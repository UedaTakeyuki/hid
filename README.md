# hid
my supplimentals of hashids

# features
## 1. The support function to create hashids.HashID.
Create original ``HashID`` handle by function [func CreateHashID(salt string, alphabet string, minlength int) (h *(hashids.HashID), err error)](https://github.com/UedaTakeyuki/hid/blob/master/hid.go#L19)

## 2. ``hidType`` and their methods to support ``Int`` (not []int) <==> ``string`` use-case
- [func CreateHID(salt string, alphabet string, minlength int) (h *hidType, err error)](https://github.com/UedaTakeyuki/hid/blob/master/hid.go#L29)
- [func (h *hidType) Encode(i int) (str string, err error)](https://github.com/UedaTakeyuki/hid/blob/master/hid.go#L42)
- [func (h *hidType) EncodeInt64(i int64) (str string, err error)](https://github.com/UedaTakeyuki/hid/blob/master/hid.go#L47)
- [func (h *hidType) Decode(str string) (i int, err error)](https://github.com/UedaTakeyuki/hid/blob/master/hid.go#L52)
- [func (h *hidType) DecodeInt64(str string) (i int64, err error)](https://github.com/UedaTakeyuki/hid/blob/master/hid.go#L62)

## 3. How to use
Refer [test code](https://github.com/UedaTakeyuki/hid/blob/master/test/hid_test.go).
