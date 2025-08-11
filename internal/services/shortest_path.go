package services

import (
	"dj-ro/internal/types"
	"fmt"
	"math"
	"sort"
	"strconv"
)

type Dijkstra struct{}

func NewDijkstra() *Dijkstra {
	return &Dijkstra{}
}

/*
Initialise un tableau de noeuds avec des propriétés
*/
func (d *Dijkstra) InitializeNodeArray(nodes types.NodesWithId, startingNodeID string, endingNodeID string) []types.ResolutionNode {
	result := make([]types.ResolutionNode, 0, len(nodes))

	// Always put starting node first
	if startNode, ok := nodes[startingNodeID]; ok {
		result = append(result, types.ResolutionNode{
			NodeWithId: startNode,
			NodePropsList: map[int]types.NodeProps{
				0: {WeightTo: 0, PreviousNode: "", Marked: false, Valid: true},
			},
		})
	}

	// Add other nodes (excluding start and end) in natural sorted order
	otherIDs := make([]string, 0, len(nodes)-2)
	for id := range nodes {
		if id != startingNodeID && id != endingNodeID {
			otherIDs = append(otherIDs, id)
		}
	}

	// Sort naturally (node1, node2, node10 instead of node1, node10, node2)
	sort.Slice(otherIDs, func(i, j int) bool {
		return naturalLess(otherIDs[i], otherIDs[j])
	})

	for _, id := range otherIDs {
		if node, ok := nodes[id]; ok {
			result = append(result, types.ResolutionNode{
				NodeWithId:    node,
				NodePropsList: map[int]types.NodeProps{0: {WeightTo: -1, PreviousNode: "", Marked: false, Valid: false}},
			})
		}
	}

	// Always put ending node last
	if endNode, ok := nodes[endingNodeID]; ok {
		result = append(result, types.ResolutionNode{
			NodeWithId: endNode,
			NodePropsList: map[int]types.NodeProps{
				0: {WeightTo: -1, PreviousNode: "", Marked: false, Valid: false},
			},
		})
	}

	return result
}

// naturalLess compares two strings naturally (numeric parts as numbers)
func naturalLess(a, b string) bool {
	ai, aj := 0, 0
	bi, bj := 0, 0

	for ai < len(a) && bi < len(b) {
		// Skip non-digits
		for ai < len(a) && !isDigit(a[ai]) {
			if bi >= len(b) || a[ai] != b[bi] {
				return a[ai] < b[bi]
			}
			ai++
			bi++
		}
		for bi < len(b) && !isDigit(b[bi]) {
			return false // a has more non-digits, so a > b
		}

		if ai >= len(a) {
			return bi < len(b) // a is shorter
		}
		if bi >= len(b) {
			return false // b is shorter
		}

		// Both are at digits, compare numerically
		aj = ai
		for aj < len(a) && isDigit(a[aj]) {
			aj++
		}
		bj = bi
		for bj < len(b) && isDigit(b[bj]) {
			bj++
		}

		aNum := parseInt(a[ai:aj])
		bNum := parseInt(b[bi:bj])

		if aNum != bNum {
			return aNum < bNum
		}

		ai, bi = aj, bj
	}

	return len(a) < len(b)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func parseInt(s string) int {
	n := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
	}
	return n
}

/*
Extrait les noeuds marqués depuis la nodeArray
*/
func extractMarkedNodes(nodeArray []types.ResolutionNode) []string {
	markedSet := make(map[string]bool)
	for _, node := range nodeArray {
		for _, nodeProp := range node.NodePropsList {
			if nodeProp.Marked {
				markedSet[node.ID] = true
				break // on sort dès qu'on trouve une cellule marquée pour ce noeud
			}
		}
	}
	markedNodes := make([]string, 0, len(markedSet))
	for id := range markedSet {
		markedNodes = append(markedNodes, id)
	}
	return markedNodes
}

/*
Crée une copie profonde de nodeArray (slice + maps internes)
*/
func deepCopyNodeArray(original []types.ResolutionNode) []types.ResolutionNode {
	copyArr := make([]types.ResolutionNode, len(original))
	for i, node := range original {
		copyArr[i] = types.ResolutionNode{
			NodeWithId:    node.NodeWithId, // Assuming this is a simple struct/value type
			NodePropsList: make(map[int]types.NodeProps, len(node.NodePropsList)),
		}
		// Deep copy the map
		for k, v := range node.NodePropsList {
			copyArr[i].NodePropsList[k] = v // NodeProps should be a value type
		}
	}
	return copyArr
}

/*
Cherche l'indice d'un noeud dans nodeArray par son ID, ou -1 si absent
*/
func findNodeIndexByID(nodeArray []types.ResolutionNode, id string) int {
	for i, node := range nodeArray {
		if node.ID == id {
			return i
		}
	}
	return -1
}

/*
Met à jour les distances dans nodeArray pour les voisins de currentNode
*/
func updateNeighborDistances(
	nodeArray []types.ResolutionNode,
	currentNode string,
	edges types.EdgesWithId,
	minDist int,
	nextRowIndex int, // The row index where we want to write new values
) error {

	for _, edge := range edges {
		var neighborId string

		// Find neighbor
		if edge.Source == currentNode {
			neighborId = edge.Target
		} else if edge.Target == currentNode {
			neighborId = edge.Source
		} else {
			continue
		}

		neighborIndex := findNodeIndexByID(nodeArray, neighborId)
		if neighborIndex == -1 {
			continue
		}

		// Check if this neighbor column is already marked (any step)
		// Once a column is marked, it should not receive any new values AT ALL
		colMarked := false
		for _, np := range nodeArray[neighborIndex].NodePropsList {
			if np.Marked {
				colMarked = true
				break
			}
		}
		if colMarked {
			continue // Skip this neighbor completely - no new entries in marked columns
		}

		weight, err := strconv.Atoi(edge.Label)
		if err != nil {
			return fmt.Errorf("poids invalide sur edge %s: %v", edge.ID, err)
		}

		newDist := minDist + weight

		// Only update if we can improve the distance or if no entry exists
		prevNodeProp, exists := nodeArray[neighborIndex].NodePropsList[nextRowIndex]
		if !exists || prevNodeProp.WeightTo == -1 || newDist < prevNodeProp.WeightTo {
			nodeArray[neighborIndex].NodePropsList[nextRowIndex] = types.NodeProps{
				WeightTo:     newDist,
				PreviousNode: currentNode,
				Marked:       false,
				Valid:        true,
			}
		}
	}
	return nil
}

/*
Cherche la valeur minimale dans les cellules valides et non marquées
Retourne la PREMIÈRE occurrence en cas d'égalité (déterminisme)
*/
func minimalEdgeSum(nodes []types.ResolutionNode) (minimalNodeIndex int, minimalNodePropIndex int, minimalWeight int) {
	minimalWeight = math.MaxInt64
	minimalNodeIndex = -1
	minimalNodePropIndex = -1

	// Parcourir les noeuds dans l'ordre (pour assurer le déterminisme)
	for nodeIndex, node := range nodes {
		// Skip if this node column is already completely marked
		colMarked := false
		for _, nodeProp := range node.NodePropsList {
			if nodeProp.Marked {
				colMarked = true
				break
			}
		}
		if colMarked {
			continue
		}

		// Find minimum valid weight in this node's column
		// Sort steps to ensure deterministic order
		steps := make([]int, 0, len(node.NodePropsList))
		for step := range node.NodePropsList {
			steps = append(steps, step)
		}
		sort.Ints(steps) // Process steps in ascending order

		for _, stepIndex := range steps {
			nodeProp := node.NodePropsList[stepIndex]
			// Only consider valid, non-marked cells with non-negative weight
			if !nodeProp.Valid || nodeProp.Marked || nodeProp.WeightTo < 0 {
				continue
			}
			if nodeProp.WeightTo < minimalWeight {
				minimalWeight = nodeProp.WeightTo
				minimalNodeIndex = nodeIndex
				minimalNodePropIndex = stepIndex
			}
			// Important: Don't break here if equal, we want the FIRST node in case of ties
		}
	}

	return
}

func (d *Dijkstra) ReconstructPath(nodeArray []types.ResolutionNode) ([]string, error) {
	if len(nodeArray) == 0 {
		return nil, fmt.Errorf("nodeArray vide")
	}

	// The ending node is the last node in nodeArray
	endNode := nodeArray[len(nodeArray)-1]

	// Find the marked step in the ending node
	var markedNodeProp types.NodeProps
	var foundMarked bool
	for _, nodeProp := range endNode.NodePropsList {
		if nodeProp.Marked {
			markedNodeProp = nodeProp
			foundMarked = true
			break
		}
	}

	if !foundMarked {
		return nil, fmt.Errorf("noeud d'arrivée non marqué")
	}

	// Build index map for quick lookup
	idToIndex := make(map[string]int, len(nodeArray))
	for i, node := range nodeArray {
		idToIndex[node.ID] = i
	}

	// Reconstruct path backwards
	path := []string{endNode.ID}
	currentNodeProp := markedNodeProp

	// Follow the chain of previous nodes
	for currentNodeProp.PreviousNode != "" {
		prevNodeId := currentNodeProp.PreviousNode
		path = append(path, prevNodeId)

		// Find the previous node in the array
		nodeIndex, exists := idToIndex[prevNodeId]
		if !exists {
			return nil, fmt.Errorf("noeud précédent %s introuvable dans nodeArray", prevNodeId)
		}

		// Find the marked property for this previous node
		foundMarkedProp := false
		for _, nodeProp := range nodeArray[nodeIndex].NodePropsList {
			if nodeProp.Marked {
				currentNodeProp = nodeProp
				foundMarkedProp = true
				break
			}
		}

		if !foundMarkedProp {
			return nil, fmt.Errorf("aucune propriété marquée trouvée pour le noeud %s", prevNodeId)
		}
	}

	// Reverse path to get correct order (start -> end)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, nil
}

/*
Résolution de l'algorithme de Dijkstra étape par étape
*/
func (d *Dijkstra) Step(
	nodes types.NodesWithId,
	edges types.EdgesWithId,
	nodeArray []types.ResolutionNode,
	currentStep int, // Algorithm iteration counter
) (types.StepResult, error) {

	// Find the minimum unmarked, valid cell
	nodeIndex, stepIndex, minDist := minimalEdgeSum(nodeArray)

	if nodeIndex == -1 {
		// No valid unmarked nodes found - check if we reached destination
		lastNode := nodeArray[len(nodeArray)-1]
		for _, nodeProp := range lastNode.NodePropsList {
			if nodeProp.Marked {
				return types.StepResult{
					NodeArray:   nodeArray,
					MarkedNodes: extractMarkedNodes(nodeArray),
					CurrentNode: "",
					Finished:    true,
				}, nil
			}
		}
		// No solution found
		return types.StepResult{
			NodeArray:   nodeArray,
			MarkedNodes: extractMarkedNodes(nodeArray),
			CurrentNode: "",
			Finished:    false,
		}, fmt.Errorf("aucune solution trouvée")
	}

	// Create deep copy
	newNodeArray := deepCopyNodeArray(nodeArray)

	// Mark the minimum cell
	nodeProp := newNodeArray[nodeIndex].NodePropsList[stepIndex]
	nodeProp.Marked = true
	newNodeArray[nodeIndex].NodePropsList[stepIndex] = nodeProp

	currentNode := newNodeArray[nodeIndex].ID

	// Check if destination is reached
	lastNode := newNodeArray[len(newNodeArray)-1]
	if lastNode.ID == currentNode {
		return types.StepResult{
			NodeArray:   newNodeArray,
			MarkedNodes: extractMarkedNodes(newNodeArray),
			CurrentNode: currentNode,
			Finished:    true,
		}, nil
	}

	// ALWAYS create a new row: find the highest existing step and add 1
	maxStep := 0
	for _, node := range newNodeArray {
		for step := range node.NodePropsList {
			if step > maxStep {
				maxStep = step
			}
		}
	}
	nextRowIndex := maxStep + 1

	// Update neighbor distances
	if err := updateNeighborDistances(newNodeArray, currentNode, edges, minDist, nextRowIndex); err != nil {
		return types.StepResult{}, err
	}

	return types.StepResult{
		NodeArray:   newNodeArray,
		MarkedNodes: extractMarkedNodes(newNodeArray),
		CurrentNode: currentNode,
		Finished:    false,
	}, nil
}
