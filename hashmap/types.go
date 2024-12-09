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
	insert()
	delete()
	lookUp()
	genHash()
	getKey()
	getValue()
	equals()
}

// methods required to design
