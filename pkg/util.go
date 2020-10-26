package pkg

import (
	"crypto/sha1"
	"fmt"
	"github.com/google/uuid"
)


func GetUniqueToken()string{
	return uuid.New().String()
}

func GetPasswordHash(pass string)(string,error){
	h := sha1.New()
	h.Write([]byte(pass))
	hashedPassword := fmt.Sprintf("%x",h.Sum(nil))
	return hashedPassword,nil
}

