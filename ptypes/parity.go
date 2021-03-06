// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package ptypes

// pt is a parity cache for all 16-bit non-negative integers.
var pt [1 << 16]uint16

// initialize parity table pt.
func init() {
	for i := 0; i < len(pt); i++ {
		pt[i] = Parity(uint64(i))
	}
}

// Parity returns 1 if the number of bits set to 1 in x is odd, otherwise O.
// The time complexity is O(log(n)) where n is the word size.
// The space complexity is O(1).
func Parity(x uint64) (p uint16) {
	x ^= x >> 32
	x ^= x >> 16
	x ^= x >> 8
	x ^= x >> 4
	x ^= x >> 2
	x ^= x >> 1
	return uint16(x & 1)
}

// ParityAlt returns 1 if the number of bits set to 1 in x is odd, otherwise O.
// The time complexity is O(k) where k is the number of bits in x set to 1.
// The space complexity is O(1).
func ParityAlt(x uint64) (p uint16) {
	for x != 0 {
		p ^= 1
		x &= x - 1
	}
	return p
}

// ParityLookup returns 1 if the number of bits set to 1 in x is odd, otherwise O.
// The time complexity is O(n/l) where n is the word size and l is the width
// of a word of cache key. The space complexity is O(1) beyond the 1<<l space
// is needed to cache precomputed results, which is constant.
func ParityLookup(x uint64) uint16 {
	return pt[(x>>48)&0xffff] ^ pt[(x>>32)&0xffff] ^ pt[(x>>16)&0xffff] ^ pt[x&0xffff]
}
