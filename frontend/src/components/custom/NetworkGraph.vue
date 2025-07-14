<script lang="ts" setup>
import { inject, ref, shallowRef, type ShallowRef } from 'vue';
import { useNodeStore } from '@/stores/nodes';
import { useEdgeStore } from '@/stores/edges';
import configs from '@/utils/vNetworkGraphConfigs';
import { VNetworkGraph, VEdgeLabel, type EventHandlers , type Instance } from 'v-network-graph'

const { getNodes } = useNodeStore();
const { getEdges } = useEdgeStore();

const isAddingNode = inject<ShallowRef<boolean>>('isAddingNode', shallowRef<boolean>(false));

const graph = ref<Instance>();
const nodes = ref(getNodes);
const edges = ref(getEdges);
const layouts = ref();

let nextNodeIndex = Object.keys(nodes.value).length + 1;

const eventHandlers: EventHandlers = {
    "view:click": ({ event }) => {
        if (!graph.value || !isAddingNode.value) return

        const point = { x: event.offsetX, y: event.offsetY };
        const svgPoint = graph.value.translateFromDomToSvgCoordinates(point);

        const nodeId = `node${nextNodeIndex}`;
        const name = `Node ${nextNodeIndex}`;

        layouts.value.nodes[nodeId] = svgPoint;

        nodes.value[nodeId] = { id: nodeId, name };
        nextNodeIndex++
    }
}
const zoomLevel = shallowRef<number>(2);
</script>

<template>
    <v-network-graph
        ref="graph"
        :nodes="nodes"
        :edges="edges"
        :layouts="layouts"
        :configs="configs"
        :eventHandlers="eventHandlers"
        :zoom-level="zoomLevel"
        class="graph border rounded-tl-lg rounded-tr-lg md:rounded-tr-none md:rounded-bl-lg bg-background dark:bg-background"
    >
        <template #edge-label="{ edge, ...slotProps }">
            <v-edge-label :text="edge.label" align="center" vertical-align="above" v-bind="slotProps"></v-edge-label>
        </template>
    </v-network-graph>
</template>

<style>
.graph {
    min-width: 100px;
    min-height: 100px;
}
</style>