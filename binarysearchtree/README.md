# Binary Search Tree

## Performance:

The time complexity of these operations is bound by *h*, the height of the tree. Provided that the tree is *balanced* (i.e. the height is log(n) of the elements in the tree), dictionary operations take O(h) time.

| Operation | Time Complexity |
|-----------|-----------------|
| Insert    | O(h)            |
| Delete    | O(h)            |
| Search    | O(h)            |

### Worst Case

Worst case performance is much like a linked list, which happens when this (unbalanced) tree is built with sorted input (e.g. inserts are ordered). This results in maximum skew, where e.g. every node has a single child on the right side, all the way down.

## TODOs/Ideas

- Genericize? Create a BST initialization function that takes a datatype and a less() function? -- that way any concrete type could be used as the stored datatype
- tree.Traverse() should return a slice, not print stuff
- Make this threadsafe with a mutex on [insert, delete, traverse]
