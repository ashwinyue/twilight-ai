package completions

import (
	"crypto/rand"
	"fmt"
)

func generateID() string {
	b := make([]byte, 12)
	rand.Read(b)
	return fmt.Sprintf("call_%x", b)
}
