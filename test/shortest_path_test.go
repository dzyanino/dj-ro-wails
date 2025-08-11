package test

import (
	"dj-ro/internal/services"
	"dj-ro/internal/types"
	"testing"
)

func TestDijkstraStepByStep(t *testing.T) {
	nodes := types.NodesWithId{
		"node1": {
			Node:             types.Node{Name: ptrString("Node 1")},
			IdentifiedObject: types.IdentifiedObject{ID: "node1"},
		},
		"node2": {
			Node:             types.Node{Name: ptrString("Node 2")},
			IdentifiedObject: types.IdentifiedObject{ID: "node2"},
		},
		"node3": {
			Node:             types.Node{Name: ptrString("Node 3")},
			IdentifiedObject: types.IdentifiedObject{ID: "node3"},
		},
	}

	edges := types.EdgesWithId{
		"edge1": {
			Edge:             types.Edge{Source: "node1", Target: "node2", Label: "1"},
			IdentifiedObject: types.IdentifiedObject{ID: "edge1"},
		},
		"edge2": {
			Edge:             types.Edge{Source: "node2", Target: "node3", Label: "2"},
			IdentifiedObject: types.IdentifiedObject{ID: "edge2"},
		},
	}

	d := services.NewDijkstra()

	nodeArray := d.InitializeNodeArray(nodes, "node1", "node3")
	currentStep := 0
	finished := false

	var currentNode string
	var markedNodes []string

	for !finished {
		var err error
		nodeArray, markedNodes, currentNode, finished, err = d.Step(nodes, edges, nodeArray, currentStep)
		if err != nil {
			t.Fatalf("Erreur Step: %v", err)
		}
		t.Logf("Étape %d: noeud courant = %s, marqués = %v", currentStep, currentNode, markedNodes)
		currentStep++
	}

	path, err := d.ReconstructPath(nodeArray)
	if err != nil {
		t.Fatalf("Erreur reconstruction chemin: %v", err)
	}
	t.Logf("Chemin optimal: %v", path)
}

/* Prendre la string depuis un pointer */
func ptrString(s string) *string {
	return &s
}
