package data_structures

import (
	"fmt"
	"strings"
)

type node struct {
	key   int
	value string
	next  *node
}

type hashMap struct {
	buckets []*node
}

func hash(key int) int {
	return key % 10
}

func NewHashMap() *hashMap {
	return &hashMap{buckets: make([]*node, 10)}
}

func (h *hashMap) Get(key int) string {
	indexPosition := hash(key)
	if h.buckets[indexPosition] != nil {
		//check if it is in the linked list of the index position
		firstNodeInList := h.buckets[indexPosition]
		for n := firstNodeInList; n != nil; n = n.next {
			if n.key == key {
				return n.value
			}
		}
	}
	return "not found"
}

func (h *hashMap) Set(key int, value string) string {
	index := hash(key)
	newNode := &node{
		key:   key,
		value: value,
	}
	//if index is empty, can insert the node
	if h.buckets[index] == nil {
		h.buckets[index] = newNode
		return value
	}
	//if some nodes already exists at that index position, loop through the linked-list
	firstNodeInList := h.buckets[index]
	for n := firstNodeInList; n != nil; n = n.next {
		//if the key exists within the list, update its value
		if n.key == key {
			n.value = value
			return n.value
		}
		// If it doesn't exist, add the new node to the end of the list
		if n.next == nil {
			n.next = newNode
			return value
		}
	}
	return ""
}

func (h *hashMap) String() string {
	var builder strings.Builder
	builder.WriteString("HashMap contents:\n")

	for i, bucket := range h.buckets {
		builder.WriteString(fmt.Sprintf("%d: ", i))
		for node := bucket; node != nil; node = node.next {
			builder.WriteString(fmt.Sprintf("(%d, %s) ", node.key, node.value))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
