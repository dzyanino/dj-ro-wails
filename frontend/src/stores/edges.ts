import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { type Edges, type EdgeWithId } from "v-network-graph";

export const useEdgeStore = defineStore("edges", () => {
    // Use Record<string, EdgeWithId> for edges
    const edges = ref<Record<string, EdgeWithId>>({});

    const getEdges = computed<Record<string, EdgeWithId>>(() => edges.value);

    function setEdges(edgesJson: Edges) {
        if (!edgesJson) {
            console.log("Empty edges");
        }
        // Directly assign the record
        // edges.value = edgesJson as Record<string, EdgeWithId>;
        Object.keys(edges.value).forEach(key => delete edges.value[key]);
        Object.assign(edges.value, edgesJson);
    }

    function addEdge(edge: EdgeWithId) {
        if (!edge || !edge.id || !edge.source || !edge.target) {
            console.error("Invalid edge object. It must have an 'id', 'source', and 'target' properties.");
            return;
        }

        if (edges.value[edge.id]) {
            console.warn(`Edge with id ${edge.id} already exists. Updating the edge.`);
        }

        const duplicateEdge = Object.values(edges.value).some(e => e.source === edge.source && e.target === edge.target);
        if (duplicateEdge) {
            console.warn(`Edge with source ${edge.source} and target ${edge.target} already exists.`);
        }

        edges.value[edge.id] = edge;
    }

    function removeEdge(edgeId: string) {
        if (edges.value[edgeId]) {
            delete edges.value[edgeId];
        } else {
            console.warn(`Edge with id ${edgeId} does not exist.`);
        }
    }

    function updateEdge(edge: EdgeWithId) {
        if (!edge || !edge.id || !edge.source || !edge.target) {
            console.error("Invalid edge object. It must have an 'id', 'source', and 'target' properties.");
            return;
        }

        if (edges.value[edge.id]) {
            edges.value[edge.id] = edge;
        } else {
            console.warn(`Edge with id ${edge.id} does not exist. Adding it instead.`);
            addEdge(edge);
        }
    }

    function getEdgeById(edgeId: string): EdgeWithId | undefined {
        return edges.value[edgeId];
    }

    function clearEdges() {
        edges.value = {};
    }

    return { getEdges, setEdges, addEdge, removeEdge, updateEdge, getEdgeById, clearEdges };
});