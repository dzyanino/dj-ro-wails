import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { type Nodes, type NodeWithId } from "v-network-graph";

export const useNodeStore = defineStore("nodes", () => {
    const nodes = ref<Record<string, NodeWithId>>({});

    const getNodes = computed<Record<string, NodeWithId>>(() => nodes.value);

    function setNodes(nodesJson: Nodes) {
        if (!nodesJson) {
            console.log("Empty nodes")
        }
        // Directly assign the record
        // nodes.value = nodesJson as Record<string, NodeWithId>;
        Object.keys(nodes.value).forEach(key => delete nodes.value[key]);
        Object.assign(nodes.value, nodesJson);
    }

    function addNode(node: NodeWithId) {
        if (!node || !node.id || !node.name) {
            console.error("Invalid node object. It must have an 'id' and 'name' properties.");
            return;
        }

        if (nodes.value[node.id]) {
            console.warn(`Node with id ${node.id} already exists.`);
            return;
        }

        const duplicateName = Object.values(nodes.value).some(n => n.name === node.name);
        if (duplicateName) {
            console.warn(`Node with name ${node.name} already exists.`);
            return;
        }

        nodes.value[node.id] = node;
    }

    function removeNode(nodeId: string) {
        if (nodes.value[nodeId]) {
            delete nodes.value[nodeId];
        } else {
            console.warn(`Node with id ${nodeId} does not exist.`);
        }
    }

    function updateNode(node: NodeWithId) {
        if (!node || !node.id || !node.name) {
            console.error("Invalid node object. It must have an 'id' and 'name' properties.");
            return;
        }

        if (nodes.value[node.id]) {
            nodes.value[node.id] = node;
        } else {
            console.warn(`Node with id ${node.id} does not exist. Adding it instead.`);
            addNode(node);
        }
    }

    return { getNodes, setNodes, addNode, removeNode, updateNode };
});