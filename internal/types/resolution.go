package types

type NodeProps struct {
	WeightTo     int    `json:"weightTo"`
	PreviousNode string `json:"previousNode"`
	Marked       bool   `json:"marked"`
	Valid        bool   `json:"valid"`
}

type ResolutionNode struct {
	NodeWithId
	NodePropsList map[int]NodeProps `json:"nodePropsList"`
}

type StepResult struct {
	NodeArray   []ResolutionNode `json:"nodeArray"`
	MarkedNodes []string         `json:"markedNodes"`
	CurrentNode string           `json:"currentNode"`
	Finished    bool             `json:"finished"`
}
