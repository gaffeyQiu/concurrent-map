package concurrent_map

import (
	"bytes"
	"fmt"
	"sync/atomic"
	"unsafe"
)

// linkedPair 单向链表
type linkedPair interface {
	
	Next() Pair

	SetNext(nextPair Pair) error
}

// Pair 并发安全的 key val 接口
type Pair interface {
	linkedPair

	Key() string

	Hash() uint64

	Element() interface{}

	SetElement(element interface{}) error

	Copy() Pair

	String() string
}

type pair struct {
	key string

	// key 的哈希值
	hash uint64
	element unsafe.Pointer
	next unsafe.Pointer
}

// newPair 初始化一个 pair 类型
func newPair(key string, element interface{}) (Pair, error) {
	p := &pair{
		key: key,
		hash: hash(key),
	}

	if element == nil {
		return nil, newIllegalParameterError("element is nil")
	}

	p.element = unsafe.Pointer(&element)
	return p, nil
}

func (p *pair) Key() string {
	return p.key
}

func (p *pair) Hash() uint64 {
	return p.hash
}

func (p *pair) Element() interface{} {
	pointer := atomic.LoadPointer(&p.element)
	if pointer == nil {
		return nil
	}

	return *(*interface{})(pointer)
}

func (p *pair) SetElement(element interface{}) error {
	if element == nil {
		return newIllegalParameterError("element is nil")
	}

	atomic.StorePointer(&p.element, unsafe.Pointer(&element))
	return nil
}

func (p *pair) Next() Pair {
	pointer := atomic.LoadPointer(&p.next)
	if pointer == nil {
		return nil
	}
	return (*pair)(pointer)
}

func (p *pair) SetNext(nextPair Pair) error {
	if nextPair == nil {
		atomic.StorePointer(&p.next, nil)
		return nil
	}

	pp, ok := nextPair.(*pair)
	if !ok {
		return newIllegalParameterError("nextPair not vaildate")
	}

	atomic.StorePointer(&p.next, unsafe.Pointer(pp))
	return nil
}

func (p *pair) Copy() Pair {
	pCopy, _ := newPair(p.Key(), p.Element())
	return pCopy
}

func (p *pair) String() string {
	return p.genString(false)
}

func (p *pair) genString(nextDetail bool) string {
	var buf bytes.Buffer
	buf.WriteString("pair{key:")
	buf.WriteString(p.Key())
	buf.WriteString(", hash:")
	buf.WriteString(fmt.Sprintf("%d", p.Hash()))
	buf.WriteString(", element:")
	buf.WriteString(fmt.Sprintf("%+v", p.Element()))

	if nextDetail {
		buf.WriteString(", next:")
		if next := p.Next(); next != nil {
			if npp, ok := next.(*pair); ok {
				buf.WriteString(npp.genString(nextDetail))
			} else {
				buf.WriteString("<ignore>")
			}
		}
	} else {
		buf.WriteString(", nextKey:")
		if next := p.Next(); next != nil {
			buf.WriteString(next.Key())
		}
	}
	buf.WriteString("}")
	return buf.String()
} 