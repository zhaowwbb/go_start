package main

import "fmt"

// Graph represents a directed graph using an adjacency list model.
type Graph struct {
	vertices int
	adjList  map[int][]int
}

// NewGraph initializes a new Graph with a fixed number of vertices.
// Acts as the constructor equivalent to public Graph(int vertices).
func NewGraph(vertices int) *Graph {
	g := &Graph{
		vertices: vertices,
		adjList:  make(map[int][]int),
	}
	for i := 0; i < vertices; i++ {
		g.adjList[i] = make([]int, 0)
	}
	return g
}

// AddEdge adds a directed edge from source to destination.
func (g *Graph) AddEdge(source, destination int) {
	g.adjList[source] = append(g.adjList[source], destination)
}

// --- DFS Version 1 (Core) ---

func (g *Graph) DFS(startNode int) {
	visited := make([]bool, g.vertices)
	fmt.Printf("Depth First Traversal starting from node %d:\n", startNode)
	g.dfsRecursive(startNode, visited)
	fmt.Println()
	fmt.Println("#################")
}

func (g *Graph) dfsRecursive(currentNode int, visited []bool) {
	visited[currentNode] = true
	fmt.Printf("%d ", currentNode)

	for _, neighbor := range g.adjList[currentNode] {
		if !visited[neighbor] {
			g.dfsRecursive(neighbor, visited)
		}
	}
}

// --- DFS Version 2 ---

func (g *Graph) DFSV2(startNode int) {
	visited := make([]bool, g.vertices)
	fmt.Printf("DFSV2, root node=%d\n", startNode)
	g.recursiveDFSV2(startNode, visited)
	fmt.Println()
	fmt.Println("#################")
}

func (g *Graph) recursiveDFSV2(startNode int, visited []bool) {
	visited[startNode] = true
	fmt.Printf("[%d]", startNode)
	for _, n := range g.adjList[startNode] {
		if !visited[n] {
			g.recursiveDFSV2(n, visited)
		}
	}
}

// --- DFS Version 3 ---

func (g *Graph) DFSV3(startNode int) {
	fmt.Printf("DFSV3(), start node=%d\n", startNode)
	visited := make([]bool, g.vertices)
	g.recursiveV3(startNode, visited)
	fmt.Println()
	fmt.Println("#################")
}

func (g *Graph) recursiveV3(currentNode int, visited []bool) {
	visited[currentNode] = true
	fmt.Printf("[%d]", currentNode)
	for _, neighbor := range g.adjList[currentNode] {
		if !visited[neighbor] {
			g.recursiveV3(neighbor, visited)
		}
	}
}

// --- DFS Version 4 ---

func (g *Graph) DFSV4(startNode int) {
	visited := make([]bool, g.vertices)
	fmt.Println("[V4]")
	g.recursiveDFS4(startNode, visited)
	fmt.Println()
	fmt.Println("#################")
}

func (g *Graph) recursiveDFS4(startNode int, visited []bool) {
	visited[startNode] = true
	fmt.Printf("[%d]", startNode)
	for _, n := range g.adjList[startNode] {
		if !visited[n] {
			g.recursiveDFS4(n, visited)
		}
	}
}

// --- DFS Version 5 ---

func (g *Graph) DFSV5(startNode int) {
	fmt.Println("[V5]")
	visited := make([]bool, g.vertices)
	g.recursiveDFSV5(startNode, visited)
	fmt.Println()
	fmt.Println("#################")
}

func (g *Graph) recursiveDFSV5(node int, visited []bool) {
	if !visited[node] {
		fmt.Printf("[%d]", node)
		visited[node] = true
	} else {
		return
	}

	list := g.adjList[node]
	for _, i := range list {
		g.recursiveDFSV5(i, visited)
	}
}
