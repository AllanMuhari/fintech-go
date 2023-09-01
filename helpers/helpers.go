package helpers

import (
	"golang.org/x/crypto/bycrypt"
)
func HandleErr(err error){
	if err != nil {
		panic(err.Error())
	}
}
func Hash
