# Binary Search Tree

## Performance:

The time complexity of these operations is bound by *h*, the height of the tree. Provided that the tree is *balanced* (i.e. the height is log(n) of the elements in the tree), dictionary operations take O(h) time.

| Operation | Time Complexity |
|-----------|-----------------|
| Insert    | O(h)            |
| Delete    | O(h)            |
| Search    | O(h)            |


## TODOs/Ideas

- tree.Traverse() should return a slice, not print stuff
- Make this threadsafe with a mutex on [insert, delete, traverse]
