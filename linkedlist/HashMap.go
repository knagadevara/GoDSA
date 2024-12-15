package linkedlist

type KV[T1, T2 comparable] struct {
	Key T1
	Val T2
}

func CreateHashNode[T1, T2 comparable](key T1, value T2) *Node[KV[T1, T2]] {
	var nwKV Node[KV[T1, T2]]
	nwKV.value.Key = key
	nwKV.value.Val = value
	nwKV.prvNode = nil
	nwKV.nextNode = nil
	return &nwKV
}

func (kv *KV[T1, T2]) settKV(k T1, v T2) *KV[T1, T2] {
	kv.Key = k
	kv.Val = v
	return kv
}

func hashPlace() {

}
