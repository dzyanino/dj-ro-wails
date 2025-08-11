package services

import (
	"dj-ro/internal/types"
	"fmt"
	"math"
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

	if startNode, ok := nodes[startingNodeID]; ok {
		result = append(result, types.ResolutionNode{
			NodeWithId: startNode,
			NodePropsList: map[int]types.NodeProps{
				0: {WeightTo: 0, PreviousNode: "", Marked: false, Valid: true},
			},
		})
	}

	otherIDs := make([]string, 0, len(nodes)-2)
	for id := range nodes {
		if id != startingNodeID && id != endingNodeID {
			otherIDs = append(otherIDs, id)
		}
	}

	for _, id := range otherIDs {
		if node, ok := nodes[id]; ok {
			result = append(result, types.ResolutionNode{
				NodeWithId:    node,
				NodePropsList: map[int]types.NodeProps{0: {WeightTo: -1, PreviousNode: "", Marked: false, Valid: false}},
			})
		}
	}

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
Extrait les noeuds marqués depuis la nodeArray
*/
func extractMarkedNodes(nodeArray []types.ResolutionNode) []string {
	markedSet := make(map[string]bool)
	for _, node := range nodeArray {
		for _, nodeProp := range node.NodePropsList {
			if nodeProp.Marked {
				markedSet[node.ID] = true
				break
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
			NodeWithId:    node.NodeWithId,
			NodePropsList: make(map[int]types.NodeProps, len(node.NodePropsList)),
		}

		for k, v := range node.NodePropsList {
			copyArr[i].NodePropsList[k] = v
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
	nextRowIndex int,
) error {

	for _, edge := range edges {
		var neighborId string

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

		/*
			Check if this neighbor column is already marked (any step)
			Once a column is marked, it should not receive any new values AT ALL
		*/
		colMarked := false
		for _, np := range nodeArray[neighborIndex].NodePropsList {
			if np.Marked {
				colMarked = true
				break
			}
		}
		if colMarked {
			continue
		}

		weight, err := strconv.Atoi(edge.Label)
		if err != nil {
			return fmt.Errorf("poids invalide sur edge %s: %v", edge.ID, err)
		}

		newDist := minDist + weight

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
*/
func minimalEdgeSum(nodes []types.ResolutionNode) (minimalNodeIndex int, minimalNodePropIndex int, minimalWeight int) {
	minimalWeight = math.MaxInt64
	minimalNodeIndex = -1
	minimalNodePropIndex = -1

	for nodeIndex, node := range nodes {
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

		for stepIndex, nodeProp := range node.NodePropsList {
			if !nodeProp.Valid || nodeProp.Marked || nodeProp.WeightTo < 0 {
				continue
			}
			if nodeProp.WeightTo < minimalWeight {
				minimalWeight = nodeProp.WeightTo
				minimalNodeIndex = nodeIndex
				minimalNodePropIndex = stepIndex
			}
		}
	}

	return
}

func (d *Dijkstra) ReconstructPath(nodeArray []types.ResolutionNode) ([]string, error) {
	if len(nodeArray) == 0 {
		return nil, fmt.Errorf("nodeArray vide")
	}

	endNode := nodeArray[len(nodeArray)-1]

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

	idToIndex := make(map[string]int, len(nodeArray))
	for i, node := range nodeArray {
		idToIndex[node.ID] = i
	}

	path := []string{endNode.ID}
	currentNodeProp := markedNodeProp

	for currentNodeProp.PreviousNode != "" {
		prevNodeId := currentNodeProp.PreviousNode
		path = append(path, prevNodeId)

		nodeIndex, exists := idToIndex[prevNodeId]
		if !exists {
			return nil, fmt.Errorf("noeud précédent %s introuvable dans nodeArray", prevNodeId)
		}

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
	currentStep int,
) (types.StepResult, error) {

	nodeIndex, stepIndex, minDist := minimalEdgeSum(nodeArray)

	if nodeIndex == -1 {
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

		return types.StepResult{
			NodeArray:   nodeArray,
			MarkedNodes: extractMarkedNodes(nodeArray),
			CurrentNode: "",
			Finished:    false,
		}, fmt.Errorf("aucune solution trouvée")
	}

	newNodeArray := deepCopyNodeArray(nodeArray)

	nodeProp := newNodeArray[nodeIndex].NodePropsList[stepIndex]
	nodeProp.Marked = true
	newNodeArray[nodeIndex].NodePropsList[stepIndex] = nodeProp

	currentNode := newNodeArray[nodeIndex].ID

	lastNode := newNodeArray[len(newNodeArray)-1]
	if lastNode.ID == currentNode {
		return types.StepResult{
			NodeArray:   newNodeArray,
			MarkedNodes: extractMarkedNodes(newNodeArray),
			CurrentNode: currentNode,
			Finished:    true,
		}, nil
	}

	maxStep := 0
	for _, node := range newNodeArray {
		for step := range node.NodePropsList {
			if step > maxStep {
				maxStep = step
			}
		}
	}
	nextRowIndex := maxStep + 1

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
