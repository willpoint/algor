package ds

// PriorityQueue ...
// A Priority Queue needs a  comparison rule that never
// contradicts itself. In order for a comparison rule, which
// we denote by <=, to be robust in this way, it must
// define a `total order relation`, which is to say that
// the comparison rule is defined for every pair of keys and
// it must satisfy the following properties:
// Reflexive property: k <= k
// Antisymmetric property: if k1  <= k2 and k2 <= k1, then k1 == k2
// Transitive property: if k1 <= k2 and k2 <= k3, then k1 <= k3
// Following these rules will prevent a comparison contradiction and
// defines a linear ordering relationship between a pair of keys.
// In this situation - the notion of smallest key Kmin, is well
// defined such Kmin <= k, for any relation k in our collection

// PriorityQueue ...
type PriorityQueue struct {
}

// Insert inserts the element p which satisfies the
// PQueuer interface into the PriorityQueue
func (pq *PriorityQueue) Insert(p PQueuer) error {
	return nil
}

// Min returns an element of PQ with the smallest asscociated
// key value
func (pq *PriorityQueue) Min() PQueuer {
	return nil
}

// RemoveMin removes from PQ the element
// associated with the return value of Min()
func (pq *PriorityQueue) RemoveMin() PQueuer {
	return nil
}

// PQueueItem ...
type PQueueItem struct {
}

// PQueuer ...
type PQueuer interface {
	Key() interface{}
	ChangeKey(key interface{}) error
}
