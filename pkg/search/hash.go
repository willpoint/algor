package search

import (
	"github.com/willpoint/algor/pkg/ds"
)

// Hash search algorithm uses a hash function h=hash(e)
// to transform the items in the collection into a value
// that is used to index into a hashtable - which is basically
// a seperate collection albeit one with a different behaviour
// To achieve its objective
// n elements of the collection C are first loaded into a hashtable
// that has `b` bins. A key is generated to aid this
// Each element e∈C can be mapped to a key value k=key(e) such
// that if Ei = Ej then key(Ei)=key(Ej).
// The hash function then uses the key(e) to determine the bin A[h]
// into which to insert e, where 0<=h<b.
// Eventually, the target element t is also transformed into
// a search for A[h] where h=hash(t)
// If it is found then true is returned otherwise - target
// does not exist.
// This algorithm has constant time 0(1) for both it's best
// and average case and 0(n) for it's worst case
// A cons against this algorithm can be with respect to storage
// space. The hashTable `A` must be large enough to hold all
// the keys with enough space for storing the collision keys.
// A good choice of A is 2^k - 1. The collision possibility means
// the element of the collection|array that serves A can be a linked list
// data structure
func Hash(ss []string, t string) int {
	// load hash table using 2^k - 1
	size := 1<<uint(len(ss)) - 1
	lt := loadHashTable(ss, size)
	_ = lt
	return hash(ss[0], size)
}

func loadHashTable(ss []string, size int) []*ds.LinkedList {
	return nil
}

// hash computes the key for a string using the popular
// technique to produce a value based on each piece of
// information from the original string
// Studies over the years about hashcode shows a better hashcode
// than this can be used for better result and reduction in
// collision.
// The value of the hashcode and the length of the bins created
// is used to find the hash by using `hashcode % bin`
func hash(s string, blen int) int {
	var k int
	for i := 0; i < len(s); i++ {
		k = 31*k + int(byte(s[i]))
	}
	return k % blen
}
