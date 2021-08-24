package concurrent_map


type ConcurrentMap interface {
	// Concurrent 返回并发量
	Concurrent() int

	// Put 保存一个键值对
	// 若 key 已存在, 则覆盖
	Put(key string, val interface{}) (bool, error)

	// Get 获取 map 中的值
	// 返回 nil 表示不存在
	Get(key string) interface{}

	// Delete 删除指定的 key
	// 如果 key 不存在则返回 false
	Delete(key string) bool

	// Len 返回 map 的长度
	Len() uint64
}

type defaultConcurrentMap struct {
	concurrent int

	len uint64
}
