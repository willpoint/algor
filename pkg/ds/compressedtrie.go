package ds

// CompressedTrie is essentially a Trie with more constraints
// enabling it to utilize memory better. It ensures that each internal
// node in the trie has at least two children, by compressing chains
// of single-child nodes into individual edges and avoiding redundant
// nodes within the Trie
// eg. if vi is redundant for i = 1, ..., k - 1
// then v0 and vk are not redundant prevented by a concatenation
// of the labels vi, ..., vk
