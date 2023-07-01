package consistenhash

import (
	"sort"
	"strconv"
)

func New(replicas int) *HashRing {
	return &HashRing{
		replicas:   replicas,
		hashMap:    map[uint32]string{},
		sortedNode: make([]uint32, 0),
	}
}

// AddNode 将单个节点的地址和它的虚拟节点添加到hash环中
func (hr *HashRing) AddNode(nodeAddr string) {
	//为单个节点添加虚拟节点
	for i := 0; i < hr.replicas; i++ {
		//获取hash值
		key := hr.hashKey(strconv.Itoa(i) + nodeAddr)
		hr.hashMap[key] = nodeAddr
		//将hash值添加到hash环中
		hr.sortedNode = append(hr.sortedNode, key)
	}

	sort.Slice(hr.sortedNode, func(i, j int) bool {
		return hr.sortedNode[i] < hr.sortedNode[j]
	})
}

// AddNodes 将多个节点添加到hash环中
func (hr *HashRing) AddNodes(nodeAddr []string) {
	for _, node := range nodeAddr {
		hr.AddNode(node)
	}

}

// RemoveNode 移除节点和虚拟节点
func (hr *HashRing) RemoveNode(nodeAddr string) {
	for i := 0; i < hr.replicas; i++ {
		//根据传入的地址值算出它对应的hash值
		key := hr.hashKey(nodeAddr)
		delete(hr.hashMap, key)

		//在hash环上删除key
		if ok, index := hr.getKeyIndex(key); ok {
			hr.sortedNode = append(hr.sortedNode[:index], hr.sortedNode[index+1:]...)
		}
	}
}

// 在hash环上获取key的索引
func (hr *HashRing) getKeyIndex(key uint32) (bool, uint32) {

	//sort.Search()
	for i, k := range hr.sortedNode {
		if k == key {
			return true, uint32(i)
		}
	}

	return false, uint32(0)
}

// GetNode 请求：客户端地址  返回：应当处理该客户端请求的服务器的地址
func (hr *HashRing) GetNode(key string) string {
	if len(hr.hashMap) == 0 {
		return ""
	}

	//计算出地址的hash值
	hashKey := hr.hashKey(key)
	//hash环
	nodes := hr.sortedNode

	//得到hash环上第一个hash值的节点地址
	masterNode := hr.hashMap[nodes[0]]

	for _, nodeKey := range hr.sortedNode {
		//如果客户端的hash值小于当前节点的hash值
		//则由当前节点来进行处理
		if hashKey < nodeKey {
			masterNode = hr.hashMap[nodeKey]
			break
		}
	}

	return masterNode
}
