<script lang="ts" setup>
import { inject, isReactive, ref, shallowRef, type Ref, type ShallowRef } from 'vue';
import { useNodeStore } from '@/stores/nodes';
import { useEdgeStore } from '@/stores/edges';
import configs from '@/utils/vNetworkGraphConfigs';
import { VNetworkGraph, VEdgeLabel, type EventHandlers , type Instance, type Layouts } from 'v-network-graph'

const { getNodes } = useNodeStore();
const { getEdges } = useEdgeStore();

const isAddingNode = inject<ShallowRef<boolean>>('isAddingNode', shallowRef<boolean>(false));

const graph = ref<Instance>();
const nodes = ref(getNodes);
const edges = ref(getEdges);
const layouts = (inject<Ref<Layouts>>('layouts', ref<Layouts>({ nodes: {} })))

let nextNodeIndex = Object.keys(nodes.value).length + 1;

const eventHandlers: EventHandlers = {
    "view:click": ({ event }) => {
        if (!graph.value || !isAddingNode.value) return

        const point = { x: event.offsetX, y: event.offsetY };
        const svgPoint = graph.value.translateFromDomToSvgCoordinates(point);

        const nodeId = `node${nextNodeIndex}`;
        const name = `Node ${nextNodeIndex}`;

        if (!!layouts.value.nodes) { layouts.value.nodes[nodeId] = svgPoint };

        nodes.value[nodeId] = { id: nodeId, name };
        nextNodeIndex++
        
        console.log(isReactive(nodes.value));
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
        class="graph border rounded-lg bg-background dark:bg-background"
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