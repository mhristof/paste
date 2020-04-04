package sha

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

func Sha() string {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return convert(data)
}

func convert(lines []byte) string {
	sum := sha256.Sum256([]byte(lines))
	return fmt.Sprintf("%x", sum)
}
