// A simple interface for loading an array from
// a JSON file and access a random element of it.
package jsonrand

type JSONRand interface {
	Load()
	Element() interface{}
}
