package main

import "testing"

// TestGraphDFS runs the different versions of the DFS implementations.
// Execute this using: go test -v -run TestGraphDFS
func TestGraphDFS(t *testing.T) {
	g := NewGraph(6)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)
	g.AddEdge(4, 5)

	// Execute your original active test lineup
	g.DFS(0)

	// Un-comment if you wish to verify intermediate versions:
	// g.DFSV2(0)
	// g.DFSV3(0)

	g.DFSV4(0)
	g.DFSV5(0)
}
