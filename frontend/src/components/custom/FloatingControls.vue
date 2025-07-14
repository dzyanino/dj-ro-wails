<script setup lang="ts">
import { onMounted, shallowRef, nextTick, inject, ref, type ShallowRef, type Ref } from 'vue'
import { useNodeStore } from '@/stores/nodes'
import { useEdgeStore } from '@/stores/edges'
import type { Layouts } from 'v-network-graph'
import { RandomGraphJS } from '../../../wailsjs/go/services/Randomizer'

import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'
import { CirclePlusIcon, PanelRightClose, PanelRightOpen, ShuffleIcon, SplineIcon, StopCircleIcon } from 'lucide-vue-next'

import NodeTable from './NodeTable.vue'

const { getNodes, setNodes } = useNodeStore();
const { getEdges, setEdges } = useEdgeStore();

const sidePanel = shallowRef<HTMLElement | null>(null)
const panelWidth = shallowRef<number>(384)
const isExpanded = shallowRef<boolean>(false)

const isAddingNode = inject<ShallowRef<boolean>>('isAddingNode', shallowRef<boolean>(false));

const layouts = (inject<Ref<Layouts>>('layouts', ref<Layouts>({ nodes: {} })))

async function createRandomGraph() {
  const generatedGraph = await RandomGraphJS("node", 10, 14, 500, 300);
  setNodes(generatedGraph.nodes);
  setEdges(generatedGraph.edges);
  layouts.value = generatedGraph.layouts;
  // console.log(getNodes)
  // console.log(getEdges)
}

onMounted(async () => {
  await nextTick()
  if (sidePanel.value) {
    panelWidth.value = sidePanel.value.getBoundingClientRect().width
  }
});
</script>

<template>
  <div
    class="fixed top-0 right-0 bottom-0 m-4 flex items-start gap-2 transition-all duration-400"
    :style="{ transform: isExpanded ? 'translateX(0)' : `translateX(${panelWidth + 16}px)` }"
  >
    <div class="h-56 flex flex-col justify-between gap-2">

      <Button variant="secondary" @click="isExpanded = !isExpanded">
        <PanelRightClose v-if="isExpanded" />
        <PanelRightOpen v-else />
      </Button>

      <div class="flex flex-col gap-2">
          <Button :variant="isAddingNode ? 'default' : 'secondary'" @click="isAddingNode = !isAddingNode">
            <StopCircleIcon v-if="isAddingNode" />
            <CirclePlusIcon v-else />
          </Button>
          <Button variant="secondary"><SplineIcon /></Button>
          <Button @click="createRandomGraph()" variant="secondary" class="mt-4"><ShuffleIcon /></Button>
      </div>

    </div>


    <div ref="sidePanel" class="w-xs md:w-sm h-full">
      <Card class="h-full bg-background dark:bg-background rounded-md">
        <CardHeader>
          <CardTitle>Tableau représentatif du graphe</CardTitle>
          <CardDescription>Arcs et sommets</CardDescription>
        </CardHeader>
        <CardContent>
          <div class="max-h-64 overflow-auto">
            <NodeTable v-if="Object.keys(getNodes).length != 0" :nodes="getNodes" :edges="getEdges" />
            <div v-else class="flex items-center justify-center h-12 border">
              <code>vide...</code>
            </div>
          </div>
        </CardContent>
        <CardFooter>
          Card Footer
        </CardFooter>
      </Card>
    </div>

  </div>
</template>
