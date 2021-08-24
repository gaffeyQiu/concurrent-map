package concurrent_map

type linkedPair interface {
	Next() Pair

	SetNext(nextPair Pair) error
}

type Pair interface {
	Key() string

	Hash() uint64

	Element() interface{}

	SetElement(element interface{}) error

	Copy() Pair

	String() string
}


