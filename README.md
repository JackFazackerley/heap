# Heap

Heap is an implementation of Go's [container/heap](https://github.com/golang/go/blob/master/src/container/heap/heap.go) package but with the use of generics from 1.18.

It makes use of [exp/constraints](golang.org/x/exp/constraints) from 1.18 in order to use the following operators `< <= >= >`.

## Using
This package is not recommended to be used, at the moment it is only a test to see how generics could be used with heaps.

## Example
An example of how to use it can be found [here](example/main.go).