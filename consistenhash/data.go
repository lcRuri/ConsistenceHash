package consistenhash

type HashRing struct {
	replicas   int               //复制因子
	hashMap    map[uint32]string //hash环
	sortedNode []uint32          //将节点排序
}
