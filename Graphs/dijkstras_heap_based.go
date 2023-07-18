package main

import "math"

// DijkstrasAlgorithm finds the shortest distances from the start vertex to all other vertices using Dijkstra's algorithm.
func DijkstrasAlgorithm(start int, edges [][][]int) []int {
	// Get the number of vertices in the graph.
	numberOfVertices := len(edges)

	// Initialize the minDistances array with maximum integer values except for the start vertex, which is set to 0.
	minDistances := make([]int, 0, numberOfVertices)
	for range edges {
		minDistances = append(minDistances, math.MaxInt32)
	}
	minDistances[start] = 0

	// Create a min-heap to store vertices and their distances from the start vertex.
	minDistancePairs := make([]Item, 0, len(edges))
	for i := range edges {
		minDistancePairs = append(minDistancePairs, Item{i, math.MaxInt32})
	}
	minDistancesHeap := NewMinHeap(minDistancePairs)
	minDistancesHeap.Update(start, 0)

	// Explore vertices in the graph using Dijkstra's algorithm until all vertices have been visited.
	for !minDistancesHeap.IsEmpty() {
		vertex, currentMinDistance := minDistancesHeap.Remove()

		// If the currentMinDistance is still set to the maximum integer value, there is no path from the start vertex to this vertex.
		if currentMinDistance == math.MaxInt32 {
			break
		}

		// Explore the outgoing edges of the current vertex.
		for _, edge := range edges[vertex] {
			destination, distanceToDestination := edge[0], edge[1]

			// Calculate the new path distance from the start vertex to the destination vertex through the current vertex.
			newPathDistance := currentMinDistance + distanceToDestination

			// Update the minDistances array and the minDistancesHeap with the new distance if it is smaller.
			currentDestinationDistance := minDistances[destination]
			if newPathDistance < currentDestinationDistance {
				minDistances[destination] = newPathDistance
				minDistancesHeap.Update(destination, newPathDistance)
			}
		}
	}

	// Construct the finalDistances array based on the minDistances array.
	finalDistances := make([]int, 0, len(minDistances))
	for _, distance := range minDistances {
		// If a vertex's distance is still set to the maximum integer value, there is no path from the start vertex to that vertex.
		// So, -1 is stored instead.
		if distance == math.MaxInt32 {
			finalDistances = append(finalDistances, -1)
		} else {
			finalDistances = append(finalDistances, distance)
		}
	}

	return finalDistances
}

// Item represents a vertex and its distance in the min-heap.
type Item struct {
	Vertex   int // Vertex represents the index of the vertex.
	Distance int // Distance represents the distance from the start vertex to the current vertex.
}

// MinHeap represents a min-heap data structure.
type MinHeap struct {
	array     []Item   // array is an array of Item structs to represent the heap.
	vertexMap map[int]int // vertexMap is a map to efficiently update and access items in the heap.
}

// NewMinHeap creates a new MinHeap instance with the given array of items.
func NewMinHeap(array []Item) *MinHeap {
	// Initialize the vertexMap to store the index of each vertex in the heap.
	vertexMap := map[int]int{}
	for _, item := range array {
		vertexMap[item.Vertex] = item.Vertex
	}

	// Create the MinHeap with the array and vertexMap.
	heap := &MinHeap{array: array, vertexMap: vertexMap}
	heap.buildHeap()
	return heap
}

// IsEmpty checks if the min-heap is empty.
func (h *MinHeap) IsEmpty() bool {
	return h.length() == 0
}

// Remove removes the item with the minimum distance from the heap and returns its vertex and distance.
func (h *MinHeap) Remove() (int, int) {
	l := h.length()
	h.swap(0, l-1)
	peeked := h.array[l-1]
	h.array = h.array[0:l-1]
	delete(h.vertexMap, peeked.Vertex)
	h.siftDown(0, l-2)
	return peeked.Vertex, peeked.Distance
}

// Update updates the distance of a vertex in the heap and maintains the heap property.
func (h *MinHeap) Update(vertex int, value int) {
	h.array[h.vertexMap[vertex]] = Item{vertex, value}
	h.siftUp(h.vertexMap[vertex])
}

// swap swaps two items in the min-heap.
func (h MinHeap) swap(i, j int) {
	h.vertexMap[h.array[i].Vertex] = j
	h.vertexMap[h.array[j].Vertex] = i
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

// length returns the length of the min-heap.
func (h MinHeap) length() int {
	return len(h.array)
}

// buildHeap performs the build-heap operation on the min-heap.
func (h *MinHeap) buildHeap() {
	first := (len(h.array) - 2) / 2
	for currentIdx := first + 1; currentIdx >= 0; currentIdx-- {
		h.siftDown(currentIdx, len(h.array)-1)
	}
}

// siftDown performs the sift-down operation to maintain the heap property.
func (h *MinHeap) siftDown(currentIdx, endIdx int) {
	childOneIdx := currentIdx*2 + 1
	for childOneIdx <= endIdx {
		childTwoIdx := -1
		if currentIdx*2+2 <= endIdx {
			childTwoIdx = currentIdx*2 + 2
		}
		indexToSwap := childOneIdx
		if childTwoIdx > -1 && h.array[childTwoIdx].Distance < h.array[childOneIdx].Distance {
			indexToSwap = childTwoIdx
		}
		if h.array[indexToSwap].Distance < h.array[currentIdx].Distance {
			h.swap(currentIdx, indexToSwap)
			currentIdx = indexToSwap
			childOneIdx = currentIdx*2 + 1
		} else {
			return
		}
	}
}

// siftUp performs the sift-up operation to maintain the heap property.
func (h *MinHeap) siftUp(currentIdx int) {
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 && h.array[currentIdx].Distance < h.array[parentIdx].Distance {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}
