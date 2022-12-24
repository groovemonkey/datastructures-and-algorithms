# Dave's Datastructures and Algorithms learning playground

This repo contains my implementations of some core datastructures and algorithms (in Go for now, maybe later in other languages).

My general plan is, for each datastructure, e.g. Binary Search Trees:

- `./binary-search-tree/`: A separate module/directory, containing
  - `binary-search-tree.go`: the implementation, along with 
  - `bst-test.go`: some unit tests (and maybe a benchmark test or two).
  - `problem-1-name.go`, `problem-2-name.go`: A few leetcode-like problems that lend themselves to solutions with this datastructure, possibly with benchmark tests that show the difference between naiive solutions and those using the optimal datastructure/algorithm.

My first crack at this is going to follow Skiena's "The Algorithm Design Manual" fairly closely. I may add stuff from CLRS and other textbooks as well if this format works well for my learning.


## Tests

Running `go test ./...` should always work:

```
➜  datastructures-and-algorithms git:(master) go test ./...
ok  	github.com/groovemonkey/datastructures-and-algorithms/binarysearchtree	(cached)
ok  	github.com/groovemonkey/datastructures-and-algorithms/priorityqueue	(cached)
```
