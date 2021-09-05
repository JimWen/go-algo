package trie

import "fmt"

type TrieNode struct {
	data         byte
	link         map[byte]*TrieNode // 多叉树存储所有下级子节点(26个字母)
	isEndingChar bool               // 标记当前点是否为字符串结束点，判断是否完全匹配
}

func getTrieNode(c byte) *TrieNode {
	return &TrieNode{c, make(map[byte]*TrieNode), false}
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{getTrieNode('/')}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	p := this.root

	for _, v := range []byte(word) {
		if _, ok := p.link[v]; !ok { // 不存在
			p.link[v] = getTrieNode(v)
		}

		p = p.link[v]
	}

	p.isEndingChar = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this.root

	for _, v := range []byte(word) { // 不存在
		if _, ok := p.link[v]; !ok {
			return false
		}

		p = p.link[v]
	}

	return p.isEndingChar
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this.root

	for _, v := range []byte(prefix) {
		if _, ok := p.link[v]; !ok { // 不存在
			return false
		}

		p = p.link[v]
	}

	return true
}

func (this *Trie) GetPrefixList(prefix string) []string {
	ret := make([]string, 0)

	p := this.root

	for _, v := range []byte(prefix) {
		if _, ok := p.link[v]; !ok { // 不存在
			return ret
		}

		p = p.link[v]
	}

	// 匹配前缀后遍历输出全部匹配
	var getFullStringFromNode func(prefix string, p *TrieNode)
	getFullStringFromNode = func(prefix string, p *TrieNode) {
		if p.isEndingChar {
			ret = append(ret, prefix)
		}

		for _, v := range p.link {
			getFullStringFromNode(fmt.Sprintf("%s%c", prefix, v.data), v)
		}
	}

	getFullStringFromNode(prefix, p)
	return ret
}
