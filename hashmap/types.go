package hashmap

type KV[T1, T2 comparable] struct {
	Key T1
	Val T2
}

type HashMap[T1, T2 comparable] struct {
	KvPair []KV[T1, T2]
	Size   uint
}

type Hashmapi interface {
	put(k int, v string)
	get(k int)
	remove(k int)
	genHash()
}
