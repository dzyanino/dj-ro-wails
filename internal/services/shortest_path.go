package services

import (
	"dj-ro/internal/types"
	"fmt"
	"math"
	"strconv"
)

type dijkstra struct{}

func NewDijkstra() *dijkstra {
	return &dijkstra{}
}

/*
Initialise un tableau de noeuds avec des propriétés
*/
func (d *dijkstra) InitializeNodeArray(nodes types.NodesWithId, startingNodeID string, endingNodeID string) []types.ResolutionNode {
	result := make([]types.ResolutionNode, 0, len(nodes))

	otherIDs := make([]string, 0, len(nodes)-2)
	for id := range nodes {
		if id != startingNodeID && id != endingNodeID {
			otherIDs = append(otherIDs, id)
		}
	}

	if startNode, ok := nodes[startingNodeID]; ok {
		result = append(result, types.ResolutionNode{
			NodeWithId: startNode,
			NodePropsList: map[int]types.NodeProps{
				0: {WeightTo: 0, PreviousNode: "", Marked: false, Valid: true},
			},
		})
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
func extractMarkedNodes(nodeArray []types.ResolutionNode) (markedNodes []string) {
	for _, node := range nodeArray {
		for _, nodeProp := range node.NodePropsList {
			if nodeProp.Marked {
				markedNodes = append(markedNodes, node.ID)
			}
		}
	}
	return markedNodes
}

/*
Crée une copie profonde de nodeArray (slice + maps internes)
*/
func deepCopyNodeArray(original []types.ResolutionNode) []types.ResolutionNode {
	copyArr := make([]types.ResolutionNode, len(original))
	copy(copyArr, original)
	for i := range copyArr {
		newMap := make(map[int]types.NodeProps, len(copyArr[i].NodePropsList))
		for k, v := range copyArr[i].NodePropsList {
			newMap[k] = v
		}
		copyArr[i].NodePropsList = newMap
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
	nextStep int,
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

		weight, err := strconv.Atoi(edge.Label)
		if err != nil {
			return fmt.Errorf("poids invalide sur edge %s: %v", edge.ID, err)
		}

		newDist := minDist + weight

		prevNodeProp, exists := nodeArray[neighborIndex].NodePropsList[nextStep]
		if !exists || prevNodeProp.WeightTo == -1 || newDist < prevNodeProp.WeightTo {
			nodeArray[neighborIndex].NodePropsList[nextStep] = types.NodeProps{
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
Cherche la valeur minimale dans les cellules valides
*/
func minimalEdgeSum(nodes []types.ResolutionNode) (minimalNodeIndex int, minimalNodePropIndex int, minimalWeight int) {
	minimalWeight = math.MaxInt64
	minimalNodeIndex = -1
	minimalNodePropIndex = -1

	for nodeIndex, node := range nodes {
		for stepIndex, nodeProp := range node.NodePropsList {
			// On ne considère que les cellules valides et non marquées
			if !nodeProp.Valid || nodeProp.Marked {
				continue
			}
			if nodeProp.WeightTo >= 0 && nodeProp.WeightTo < minimalWeight {
				minimalWeight = nodeProp.WeightTo
				minimalNodeIndex = nodeIndex
				minimalNodePropIndex = stepIndex
			}
		}
	}

	return
}

func (d *dijkstra) ReconstructPath(nodeArray []types.ResolutionNode) ([]string, error) {
	if len(nodeArray) == 0 {
		return nil, fmt.Errorf("nodeArray vide")
	}

	/* Le dernier noeud dans nodeArray est le noeud d'arrivée */
	endNode := nodeArray[len(nodeArray)-1]

	/* Trouver la dernière étape marquée dans NodePropsList du noeud d'arrivée */
	var lastStep int = -1
	for step := range endNode.NodePropsList {
		if endNode.NodePropsList[step].Marked && step > lastStep {
			lastStep = step
		}
	}

	if lastStep == -1 {
		return nil, fmt.Errorf("noeud d'arrivée non marquée")
	}

	path := []string{endNode.ID}
	currendNodeId := endNode.ID
	currentStep := lastStep

	/* Rechercher les indices des nodes pour accéder facilement */
	idToIndex := make(map[string]int, len(nodeArray))
	for i, node := range nodeArray {
		idToIndex[node.ID] = i
	}

	for currentStep > 0 {
		nodeIndex := idToIndex[currendNodeId]
		nodeProp := nodeArray[nodeIndex].NodePropsList[currentStep]
		prevNodeId := nodeProp.PreviousNode

		if prevNodeId == "" {
			return nil, fmt.Errorf("chemin interrompu à l'étape %d pour le noeud %s", currentStep, currendNodeId)
		}

		path = append(path, prevNodeId)
		currendNodeId = prevNodeId
		currentStep--
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, nil
}

/*
Résolution de l'algorithme de Dijkstra étape par étape
Prends la liste des noeuds, des arcs, la `nodeArray` générée et le `currentStep`
Pour retourner la nouvelle `newNodeArray`, les `markedNodes`, `currentNode`,
si oui l'algorithme est fini avec `finished`, et potentiellement un erreur `err`
*/
func (d *dijkstra) Step(
	nodes types.NodesWithId,
	edges types.EdgesWithId,
	nodeArray []types.ResolutionNode,
	currentStep int,
) (newNodeArray []types.ResolutionNode, markedNodes []string, currentNode string, finished bool, err error) {

	nodeIndex, stepIndex, minDist := minimalEdgeSum(nodeArray)

	if nodeIndex == -1 {
		/* Aucun noeud valide non marqué trouvé → fin ou pas de solution */
		return nodeArray, extractMarkedNodes(nodeArray), "", false, nil
	}

	/* On fait une copie profonde du nodeArray et de ses NodePropsList */
	newNodeArray = deepCopyNodeArray(nodeArray)

	/* Marquer la cellule minimale trouvée */
	nodeProp := newNodeArray[nodeIndex].NodePropsList[stepIndex]
	nodeProp.Marked = true
	newNodeArray[nodeIndex].NodePropsList[stepIndex] = nodeProp

	currentNode = newNodeArray[nodeIndex].ID

	/* Vérifier si le dernier noeud (endingNode) est marqué à la même étape stepIndex */
	lastNode := newNodeArray[len(newNodeArray)-1]
	nodeProps, exists := lastNode.NodePropsList[stepIndex]

	if exists && nodeProps.Marked {
		return newNodeArray, extractMarkedNodes(newNodeArray), currentNode, true, nil
	}

	nextStep := stepIndex + 1

	/* Mettre à jour les voisins sur la ligne suivante nextStep */
	err = updateNeighborDistances(newNodeArray, currentNode, edges, minDist, nextStep)
	if err != nil {
		return nodeArray, nil, "", false, err
	}

	markedNodes = extractMarkedNodes(newNodeArray)

	return newNodeArray, markedNodes, currentNode, false, nil
}
