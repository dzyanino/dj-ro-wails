<script setup lang="ts">
import { onMounted, shallowRef, nextTick, inject, ref, type ShallowRef, type Ref } from 'vue'
import { useNodeStore } from '@/stores/nodes'
import { useEdgeStore } from '@/stores/edges'
import type { Layouts } from 'v-network-graph'
import { RandomGraphJS } from '../../../wailsjs/go/services/Randomizer'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'
import { CirclePlusIcon, Loader2, PanelRightClose, PanelRightOpen, ShuffleIcon, SplineIcon, StopCircleIcon } from 'lucide-vue-next'
import NodeTable from './NodeTable.vue'

const { getNodes, setNodes } = useNodeStore();
const { getEdges, setEdges } = useEdgeStore();

const sidePanel = shallowRef<HTMLElement | null>(null)
const panelWidth = shallowRef<number>(384)
const isExpanded = shallowRef<boolean>(false)

const isAddingNode = inject<ShallowRef<boolean>>('isAddingNode', shallowRef<boolean>(false));
const isAddingEdge = inject<ShallowRef<boolean>>('isAddingEdge', shallowRef<boolean>(false));
const isCreatingRandomGraph = shallowRef<boolean>(false);

// const selectedNodes = inject<Ref<string[]>>('selectedNodes', ref<string[]>([]));

const layouts = (inject<Ref<Layouts>>('layouts', ref<Layouts>({ nodes: {} })))

async function createRandomGraph() {
  isCreatingRandomGraph.value = true;

  const generatedGraph = await RandomGraphJS("node", 10, 14, 50, 50);
  setNodes(generatedGraph.nodes);
  setEdges(generatedGraph.edges);
  layouts.value = generatedGraph.layouts;

  isCreatingRandomGraph.value = false;
}

onMounted(async () => {
  await nextTick()
  if (sidePanel.value) {
    panelWidth.value = sidePanel.value.getBoundingClientRect().width
  }
});
</script>

<template>
  <div class="fixed top-0 right-0 bottom-0 m-4 flex items-start gap-2 transition-all duration-400"
    :style="{ transform: isExpanded ? 'translateX(0)' : `translateX(${panelWidth + 16}px)` }">
    <div class="h-56 flex flex-col justify-between gap-2">

      <Button variant="secondary" @click="isExpanded = !isExpanded">
        <PanelRightClose v-if="isExpanded" />
        <PanelRightOpen v-else />
      </Button>

      <div class="flex flex-col gap-2">
        <Button @click="isAddingNode = !isAddingNode; isAddingEdge = false"
          :variant="isAddingNode ? 'default' : 'secondary'">
          <StopCircleIcon v-if="isAddingNode" />
          <CirclePlusIcon v-else />
        </Button>
        <Button @click="isAddingEdge = !isAddingEdge; isAddingNode = false"
          :variant="isAddingEdge ? 'default' : 'secondary'">
          <StopCircleIcon v-if="isAddingEdge" />
          <SplineIcon v-else />
        </Button>
        <Button @click="createRandomGraph()" :disabled="isCreatingRandomGraph" variant="secondary" class="mt-4">
          <Loader2 v-if="isCreatingRandomGraph" class="w-4 h-4 animate-spin" />
          <ShuffleIcon v-else />
        </Button>
      </div>

    </div>


    <div ref="sidePanel" class="w-xs md:w-sm h-full">
      <Card class="h-full bg-background dark:bg-background rounded-md">
        <CardHeader>
          <CardTitle>Tableau représentatif du graphe</CardTitle>
          <CardDescription>Arcs et sommets</CardDescription>
        </CardHeader>
        <CardContent>
          <div class="max-h-64 relative overflow-auto">
            <NodeTable v-if="Object.keys(getNodes).length != 0" :nodes="getNodes" :edges="getEdges" />
            <div v-else class="flex items-center justify-center h-12 border">
              <code>vide...</code>
            </div>
          </div>
        </CardContent>
        <CardFooter>
          <p>card footer</p>
        </CardFooter>
      </Card>
    </div>

  </div>
</template>
