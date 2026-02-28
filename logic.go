package main

import(
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const base = uint64(len(charset))


//Turns a database ID into Base64 string
func Encode(id uint64)string{
	if id == 0{
		return string(charset[0])
	}
	var sb strings.Builder
	for id > 0{
		rem:=id%base
		sb.WriteByte(charset[rem])
		id = id / base
	}
	return reverse(sb.String())
}

func reverse(s string)string{
	runes:= []rune(s)

	for i,j :=0,len(runes)-1; i<j;i,j=i+1,j-1{
		runes[i],runes[j]=runes[j],runes[i]
	}
	return string(runes)
}