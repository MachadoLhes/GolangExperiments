package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func RandomInRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {
	silabas := []string{"a", "i", "u", "e", "o", "ka", "ki", "ku", "ke", "ko", "ba", "bi", "bu", "be", "bo",
		"sa", "shi", "su", "se", "so", "ta", "tchi", "tsu", "te", "to", "ra", "ri", "ru", "re", "ro"}
	num_silabas := RandomInRange(2, 5)
	var buffer bytes.Buffer
	for i := 0; i < num_silabas; i++ {
		buffer.WriteString(silabas[rand.Intn(len(silabas))])
	}
	fmt.Println(buffer.String())
}
