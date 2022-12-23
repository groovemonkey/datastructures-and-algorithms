# Priority Queue, three ways

## Performance:

The trick here, regardless of which datastructure is backing the Queue, is to maintain a pointer to the minimum value to make the find-min operation cheap.

| Operation  | Unsorted Array | Sorted Array | Balanced Tree |
|------------|----------------|--------------|---------------|
| Insert     | O(1)           | O(n)         | O(log n)      |
| Find-Min   | O(1)           | O(1)         | O(1)          |
| Delete-Min | O(n)           | O(1)         | O(log n)      |


## TODOs/Ideas

- implement as sorted array
- use a balanced tree, not a regular binary search treee
