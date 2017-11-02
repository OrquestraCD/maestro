package stringid

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateID() string {
	seed := strconv.Itoa(int(time.Now().Unix() + rand.Int63()))

	return fmt.Sprintf("%x", sha1.Sum([]byte(seed)))[0:16]
}
