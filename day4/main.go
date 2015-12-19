package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const key string = "yzbqklnj"

func main() {
	for _, prefix := range []string{"00000", "000000"} {
		var i int

		for {
			checkSum := md5.Sum([]byte(fmt.Sprintf("%s%d", key, i)))
			if strings.HasPrefix(fmt.Sprintf("%x", checkSum), prefix) {
				break
			}
			i++
		}

		fmt.Printf("First number with %s hash prefix is %d\n", prefix, i)
	}
}
