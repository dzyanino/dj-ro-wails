package types

type IdentifiedObject struct {
	ID string `json:"id"`
}

/*
* Different types of nodes
 */
type Node struct {
	Name       *string        `json:"name,omitempty"`
	Properties map[string]any `json:"properties,omitempty"`
}
type NodeWithId struct {
	Node
	IdentifiedObject
}
type Nodes map[string]Node

/*
* Different types of edges
 */
type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Label  string `json:"label,omitempty"`
}
type EdgeWithId struct {
	Edge
	IdentifiedObject
}
type Edges map[string]Edge

/*
* Layouting
 */
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
type FixablePosition struct {
	Position
	Fixed *bool `json:"fixed,omitempty"`
}
type NodePositions map[string]FixablePosition
type Layouts struct {
	Nodes NodePositions `json:"nodes"`
}
