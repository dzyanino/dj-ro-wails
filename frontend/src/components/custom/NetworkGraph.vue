<script lang="ts" setup>
import { inject, ref, shallowRef, watch, type ComputedRef, type Ref, type ShallowRef } from 'vue';
import { useNodeStore } from '@/stores/nodes';
import { useEdgeStore } from '@/stores/edges';
import { VNetworkGraph, VEdgeLabel, type EventHandlers, type Instance, type Layouts, type UserConfigs } from 'v-network-graph';
import { toast } from 'vue-sonner';
import ModalDialog from '@/components/custom/Dialogs/ModalDialog.vue';
import NumberField from './Inputs/NumberField.vue';

const { getNodes, removeNode } = useNodeStore();
const { getEdges, removeEdge } = useEdgeStore();

const configs = inject<ComputedRef<UserConfigs>>('configs');

const isAddingNode = inject<ShallowRef<boolean>>('isAddingNode', shallowRef<boolean>(false));
const isAddingEdge = inject<ShallowRef<boolean>>('isAddingEdge', shallowRef<boolean>(false));
const isAddingEdgeDialogOpen = inject<ShallowRef<boolean>>('isAddingEdgeDialogOpen', shallowRef<boolean>(false));

const selectedNodes = inject<Ref<string[]>>('selectedNodes', ref<string[]>([]));
const selectedEdges = inject<Ref<string[]>>('selectedEdges', ref<string[]>([]));

const graph = ref<Instance>();
const nodes = ref(getNodes);
const edges = ref(getEdges);
const layouts = (inject<Ref<Layouts>>('layouts', ref<Layouts>({ nodes: {} })));

const nextNodeIndex = shallowRef<number>(Object.keys(nodes.value).length + 1);
const nextEdgeIndex = shallowRef<number>(Object.keys(edges.value).length + 1);

const eventHandlers: EventHandlers = {
  "view:click": ({ event }) => {
    if (!graph.value || !isAddingNode.value) return

    const point = { x: event.offsetX, y: event.offsetY };
    const svgPoint = graph.value.translateFromDomToSvgCoordinates(point);

    const nodeId = `node${nextNodeIndex.value}`;
    const name = `Node ${nextNodeIndex.value}`;

    if (!!layouts.value.nodes) { layouts.value.nodes[nodeId] = svgPoint };

    nodes.value[nodeId] = { id: nodeId, name };
    nextNodeIndex.value++
  }
}
const zoomLevel = shallowRef<number>(3);

const edgeWeight = shallowRef<number>(Math.floor(Math.random() * 50) + 1);

function cancelEdgeAddition() {
  selectedNodes.value = [];
  edgeWeight.value = Math.floor(Math.random() * 50) + 1;
}

function addEdge() {
  const [sourceId, targetId] = selectedNodes.value;
  const duplicationExists = Object.values(edges.value).some(edge =>
    (edge.source === sourceId && edge.target === targetId) ||
    (edge.source === targetId && edge.target === sourceId)
  );

  if (duplicationExists) {
    toast('Attention', {
      description: 'Un arc existe déjà pour ces sommets',
    });
  }
  else {
    const edgeId = `edge${nextEdgeIndex.value}`;
    edges.value[edgeId] = {
      id: edgeId,
      source: `${selectedNodes.value[0]}`,
      target: `${selectedNodes.value[1]}`,
      label: edgeWeight.value.toString(),
    };

    nextEdgeIndex.value++;
  }


  selectedNodes.value = [];
  edgeWeight.value = Math.floor(Math.random() * 50) + 1;
}

function deleteSelectedElements() {
  if (selectedNodes.value.length > 0) {
    selectedNodes.value.forEach(nodeId => {
      removeNode(nodeId);
      delete layouts.value.nodes[nodeId];
    });

    selectedNodes.value = [];
  }

  else if (selectedEdges.value.length > 0) {
    selectedEdges.value.forEach(edgeId => {
      removeEdge(edgeId)
    });

    selectedEdges.value = [];
  }
}

watch(() => Object.keys(nodes.value).length, (newLength) => {
  nextNodeIndex.value = newLength + 1;
}, { immediate: true });

watch(() => Object.keys(edges.value).length, (newLength) => {
  nextEdgeIndex.value = newLength + 1;
}, { immediate: true });

// watch(isAddingEdge, (isIt: boolean) => {
//   if (configs?.value && configs.value.node) configs.value.node.selectable = isIt ? 2 : 1;
// });

watch(selectedNodes, () => {
  isAddingEdgeDialogOpen.value = isAddingEdge.value && selectedNodes.value.length == 2;
});
</script>

<template>
  <v-network-graph ref="graph" v-model:selected-nodes="selectedNodes" v-model:selected-edges="selectedEdges"
    :nodes="nodes" :edges="edges" :layouts="layouts" :configs="configs" :eventHandlers="eventHandlers"
    :zoom-level="zoomLevel" tabindex="0" @keyup.delete="deleteSelectedElements"
    class="h-full graph border rounded-lg bg-background dark:bg-background">
    <template #edge-label="{ edge, ...slotProps }">
      <v-edge-label :text="edge.label" align="center" vertical-align="above" v-bind="slotProps"></v-edge-label>
    </template>
  </v-network-graph>

  <ModalDialog v-model:open="isAddingEdgeDialogOpen" :onCancel="cancelEdgeAddition" :onConfirm="addEdge">
    <template #dialog-title>Ajouter un arc</template>
    <template #dialog-description>Veuillez saisir le poids</template>
    <template #dialog-content>
      <NumberField v-model:model-value="edgeWeight" id="weight" :min="1" />
    </template>
  </ModalDialog>
</template>

<style>
.graph {
  min-width: 100px;
  min-height: 100px;
}
</style>