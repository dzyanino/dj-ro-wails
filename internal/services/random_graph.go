package services

import (
	"fmt"
	"math/rand"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"dj-ro/internal/types"
)

type Randomizer struct{}

func NewRandomizer() *Randomizer {
	return &Randomizer{}
}

/*
 * Function to create a key for each edge by comparing both sides to check for duplication
 */
func makeEdgeKey(a, b string) string {
	if a > b {
		a, b = b, a
	}

	return a + "|" + b
}

/*
 * Function to create a random graph from counts and node's name
 */
func (r *Randomizer) RandomGraph(
	nodeName string,
	nodeCount, edgeCount, width, height int,
) (types.Nodes, types.Edges, types.Layouts) {

	nodes := make(types.Nodes)

	loweredName := cases.Lower(language.English).String(nodeName)   /* used for key/id */
	titledName := cases.Title(language.English).String(loweredName) /* used for title/display */

	for i := 1; i <= nodeCount; i++ {
		name := fmt.Sprintf("%s %d", titledName, i)

		/* creates nodes with corresponding name and key */
		nodes[fmt.Sprintf("%s%d", loweredName, i)] = types.Node{
			Name: &name,
		}
	}

	edges := make(types.Edges)
	edgePairs := make(map[string]struct{})
	edgeId := 1

	nodeKeys := make([]string, 0, len(nodes))

	/* go through all nodes and retrieve each key */
	for k := range nodes {
		nodeKeys = append(nodeKeys, k)
	}

	for i := 0; i < len(nodeKeys)-1 && edgeId <= edgeCount; i++ {
		source := nodeKeys[i]
		target := nodeKeys[i+1]
		key := makeEdgeKey(source, target)

		edgePairs[key] = struct{}{}
		label := fmt.Sprintf("%d", rand.Intn(100))

		edges[fmt.Sprintf("edge%d", edgeId)] = types.Edge{
			Source: source,
			Target: target,
			Label:  label,
		}

		edgeId++
	}

	for edgeId <= edgeCount {
		source := nodeKeys[rand.Intn(nodeCount)]
		target := nodeKeys[rand.Intn(nodeCount)]

		if source == target {
			continue
		}

		key := makeEdgeKey(source, target)

		if _, exists := edgePairs[key]; exists {
			continue
		}

		edgePairs[key] = struct{}{}
		label := fmt.Sprintf("%d", rand.Intn(100))
		edges[fmt.Sprintf("edge%d", edgeId)] = types.Edge{
			Source: source,
			Target: target,
			Label:  label,
		}

		edgeId++
	}

	layouts := types.Layouts{
		Nodes: make(types.NodePositions),
	}

	for nodeKey := range nodes {
		layouts.Nodes[nodeKey] = types.FixablePosition{
			Position: types.Position{
				X: float64(rand.Intn(width)),
				Y: float64(rand.Intn(height)),
			},
		}
	}

	return nodes, edges, layouts
}

/*
* Needs another function that will combine the data(nodes/edges/layouts) into a single map
* As the JS part is expecting a single object regrouping the three of them
* And thus making it serializable through JSON
 */
func (r *Randomizer) RandomGraphJS(
	nodeName string,
	nodeCount, edgeCount, width, height int,
) map[string]interface{} {
	nodes, edges, layouts := r.RandomGraph(nodeName, nodeCount, edgeCount, width, height)

	return map[string]interface{}{
		"nodes":   nodes,
		"edges":   edges,
		"layouts": layouts,
	}
}
