package hashmap

type KV struct {
	Key int
	Val string
}

type HashMap struct {
	KvPair []KV
	Size   int
}

type Hashmapi interface {
	put(k int, v string)
	get(k int)
	remove(k int)
	genHash()
}

// methods required to design
