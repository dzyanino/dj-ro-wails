package services

import (
	"dj-ro/internal/types"
	"fmt"
	"math"
	"sort"
	"strconv"
)

type LongestPath struct{}

func NewLongestPath() *LongestPath {
	return &LongestPath{}
}

/*
Initialise un tableau de noeuds avec des propriétés pour le chemin le plus long
*/
func (lp *LongestPath) InitializeNodeArray(nodes types.NodesWithId, startingNodeID string, endingNodeID string) []types.ResolutionNode {
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

/*
Met à jour les distances dans nodeArray pour les voisins de currentNode (version chemin le plus long)
*/
func updateNeighborDistancesMax(
	nodeArray []types.ResolutionNode,
	currentNode string,
	edges types.EdgesWithId,
	maxDist int,
	nextRowIndex int, // The row index where we want to write new values
) error {

	for _, edge := range edges {
		var neighborId string

		// For longest path, we only follow directed edges FROM current node
		if edge.Source == currentNode {
			neighborId = edge.Target
		} else {
			continue // Skip reverse edges - this is for directed graphs only
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

		newDist := maxDist + weight

		// Only update if we can improve the distance (find MAXIMUM) or if no entry exists
		prevNodeProp, exists := nodeArray[neighborIndex].NodePropsList[nextRowIndex]
		if !exists || prevNodeProp.WeightTo == -1 || newDist > prevNodeProp.WeightTo {
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
Cherche la valeur maximale dans les cellules valides et non marquées
Retourne la PREMIÈRE occurrence en cas d'égalité (déterminisme)
*/
func maximalEdgeSum(nodes []types.ResolutionNode) (maximalNodeIndex int, maximalNodePropIndex int, maximalWeight int) {
	maximalWeight = math.MinInt64
	maximalNodeIndex = -1
	maximalNodePropIndex = -1

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

		// Find maximum valid weight in this node's column
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
			if nodeProp.WeightTo > maximalWeight {
				maximalWeight = nodeProp.WeightTo
				maximalNodeIndex = nodeIndex
				maximalNodePropIndex = stepIndex
			}
			// Important: Don't break here if equal, we want the FIRST node in case of ties
		}
	}

	return
}

func (lp *LongestPath) ReconstructPath(nodeArray []types.ResolutionNode) ([]string, error) {
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
Résolution de l'algorithme du chemin le plus long étape par étape
*/
func (lp *LongestPath) Step(
	nodes types.NodesWithId,
	edges types.EdgesWithId,
	nodeArray []types.ResolutionNode,
	currentStep int, // Algorithm iteration counter
) (types.StepResult, error) {

	// Find the maximum unmarked, valid cell
	nodeIndex, stepIndex, maxDist := maximalEdgeSum(nodeArray)

	if nodeIndex == -1 {
		// No more valid unmarked nodes found - algorithm finished
		return types.StepResult{
			NodeArray:   nodeArray,
			MarkedNodes: extractMarkedNodes(nodeArray),
			CurrentNode: "",
			Finished:    true,
		}, nil
	}

	// Create deep copy
	newNodeArray := deepCopyNodeArray(nodeArray)

	// Mark the maximum cell
	nodeProp := newNodeArray[nodeIndex].NodePropsList[stepIndex]
	nodeProp.Marked = true
	newNodeArray[nodeIndex].NodePropsList[stepIndex] = nodeProp

	currentNode := newNodeArray[nodeIndex].ID

	// Don't stop when reaching destination - continue exploring all paths
	// The algorithm only stops when no more valid nodes can be processed

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

	// Update neighbor distances (looking for maximum path)
	if err := updateNeighborDistancesMax(newNodeArray, currentNode, edges, maxDist, nextRowIndex); err != nil {
		return types.StepResult{}, err
	}

	return types.StepResult{
		NodeArray:   newNodeArray,
		MarkedNodes: extractMarkedNodes(newNodeArray),
		CurrentNode: currentNode,
		Finished:    false,
	}, nil
}
