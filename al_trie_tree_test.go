package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestTrieNode() *TrieNode {
	node := InitTrieNode()
	node.Insert("abc")
	node.Insert("ab")
	node.Insert("bcd")
	node.Insert("def")
	return node
}

func TestTrieNode(t *testing.T) {
	node := getTestTrieNode()
	actual := node.Search("abc")
	assert.Equal(t, 1, actual, "")
	node.Insert("abc")
	actual = node.Search("abc")
	assert.Equal(t, 2, actual, "")
	node.Delete("abc")
	actual = node.Search("abc")
	assert.Equal(t, 1, actual, "")
	node.Delete("def")
	actual = node.Search("abc")
	assert.Equal(t, 1, actual, "")
	actual = node.PrefixNumber("ab")
	assert.Equal(t, 2, actual, "")
	node.Insert("abbsdf")
	actual = node.PrefixNumber("ab")
	assert.Equal(t, 3, actual, "")
}

func TestNQueue(t *testing.T) {
	maps := make(map[int]int)
	maps[1] = 1
	maps[2] = 0
	maps[3] = 0
	maps[4] = 2
	maps[5] = 10
	maps[6] = 4
	maps[7] = 40
	maps[8] = 92
	maps[9] = 352
	maps[10] = 724
	// maps[11] = 2680
	// maps[12] = 14200
	// maps[13] = 73712
	// maps[14] = 365596
	// maps[15] = 2279184
	// maps[16] = 14772512
	// maps[17] = 95815104
	// maps[18] = 666090624
	// maps[19] = 4968057848
	// maps[20] = 39029188884
	// maps[21] = 314666222712
	// maps[22] = 2691008701644
	// maps[23] = 24233937684440
	// maps[24] = 227514171973736
	// maps[25] = 2207893435808352
	for k, v := range maps {
		actual := NQueue(k)
		assert.Equal(t, v, actual, "")
	}
}

func BenchmarkNQueue(b *testing.B) {
	maps := make(map[int]int)
	// maps[1] = 1
	// maps[2] = 0
	// maps[3] = 0
	// maps[4] = 2
	// maps[5] = 10
	// maps[6] = 4
	// maps[7] = 40
	// maps[8] = 92
	// maps[9] = 352
	// maps[10] = 724
	// maps[11] = 2680
	// maps[12] = 14200
	//maps[13] = 73712
	// maps[14] = 365596
	maps[15] = 2279184
	// maps[16] = 14772512
	// maps[17] = 95815104
	// maps[18] = 666090624
	// maps[19] = 4968057848
	// maps[20] = 39029188884
	// maps[21] = 314666222712
	// maps[22] = 2691008701644
	// maps[23] = 24233937684440
	// maps[24] = 227514171973736
	// maps[25] = 2207893435808352
	for k, v := range maps {
		actual := NQueue(k)
		assert.Equal(b, v, actual, "")
	}
}

func BenchmarkNQueue_2(b *testing.B) {
	maps := make(map[int]int)
	// maps[1] = 1
	// maps[2] = 0
	// maps[3] = 0
	// maps[4] = 2
	// maps[5] = 10
	// maps[6] = 4
	// maps[7] = 40
	// maps[8] = 92
	// maps[9] = 352
	// maps[10] = 724
	// maps[11] = 2680
	// maps[12] = 14200
	//maps[13] = 73712
	// maps[14] = 365596
	//maps[15] = 2279184
	maps[16] = 14772512
	// maps[17] = 95815104
	// maps[18] = 666090624
	// maps[19] = 4968057848
	// maps[20] = 39029188884
	// maps[21] = 314666222712
	// maps[22] = 2691008701644
	// maps[23] = 24233937684440
	// maps[24] = 227514171973736
	// maps[25] = 2207893435808352
	for k, v := range maps {
		actual := NQueue_2(k)
		assert.Equal(b, v, actual, "")
	}
}
