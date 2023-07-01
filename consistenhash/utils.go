package consistenhash

import "hash/crc32"

// 节点或客户端地址
// 返回：地址所对应的哈希值
func (hr *HashRing) hashKey(key string) uint32 {
	scratch := []byte(key)
	return crc32.ChecksumIEEE(scratch)
}
