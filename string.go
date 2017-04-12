package jsonrand

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// A jsonrand.String is a simple JSON array of strings.
type String struct {
	elements []string
}

// Load takes a filename and attempts to parse the JSON it contains.
func (r *String) Load(filename string) {
	rand.Seed(time.Now().UTC().UnixNano())

	file, e := ioutil.ReadFile(filename)
	if e != nil {
		panic(fmt.Sprintf("File Error: [%v]\n", e))
	}
	e = json.Unmarshal(file, &r.elements)
	if e != nil {
		panic(fmt.Sprintf("File Error: [%v]\n", e))
	}
}

// Element returns a random element from the loaded file.
func (r *String) Element() string {
	i := rand.Intn(len(r.elements) - 1)
	fmt.Println(i)
	return r.elements[i]
}
