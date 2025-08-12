<script setup lang="ts">
import { onMounted, shallowRef, nextTick, inject, ref, type ShallowRef, type Ref, computed, watch } from 'vue'
import { useNodeStore } from '@/stores/nodes'
import { useEdgeStore } from '@/stores/edges'
import type { Layouts } from 'v-network-graph'
import { RandomGraphJS } from '../../../../wailsjs/go/services/Randomizer'
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { CirclePlusIcon, Loader2, PanelRightClose, PanelRightOpen, ShuffleIcon, SplineIcon, StopCircleIcon, BrushCleaningIcon, WaypointsIcon, RouteIcon, InfoIcon } from 'lucide-vue-next'
import NodeTable from './NodeTable.vue'
import ModalDialog from '@/components/custom/Dialogs/ModalDialog.vue'
import Select from '@/components/custom/Selects/Select.vue'

const { getNodes, setNodes, clearNodes } = useNodeStore();
const { getEdges, setEdges, clearEdges } = useEdgeStore();

const sidePanel = shallowRef<HTMLElement | null>(null)
const panelWidth = shallowRef<number>(384)
const isExpanded = shallowRef<boolean>(false)

const nodePrefix = inject<ShallowRef<string>>('nodePrefix', shallowRef<string>('Sommet'));

const isAddingNode = inject<ShallowRef<boolean>>('isAddingNode', shallowRef<boolean>(false));
const isAddingEdge = inject<ShallowRef<boolean>>('isAddingEdge', shallowRef<boolean>(false));
const isCreatingRandomGraph = shallowRef<boolean>(false);
const isClearingGraph = shallowRef<boolean>(false);
const isClearingGraphDialogOpen = shallowRef<boolean>(false);

const selectedNodes = inject<Ref<string[]>>('selectedNodes', ref<string[]>([]));
const selectedEdges = inject<Ref<string[]>>('selectedEdges', ref<string[]>([]));

const layouts = (inject<Ref<Layouts>>('layouts', ref<Layouts>({ nodes: {} })));

const selectedStart = shallowRef<string>('');
const selectedEnd = shallowRef<string>('');

const resolutionValidNodes = computed(() => {
  if (Object.keys(getEdges).length != 0) {
    const connectedNodeIds = new Set<string>()

    Object.values(getEdges).forEach(edge => {
      connectedNodeIds.add(edge.source)
      connectedNodeIds.add(edge.target)
    });

    return Object.values(getNodes).filter(node => connectedNodeIds.has(node.id));
  }

  return [];
});

const filteredStartNodes = computed(() =>
  resolutionValidNodes.value.filter(node => node.id !== selectedEnd.value)
);

const filteredEndNodes = computed(() =>
  resolutionValidNodes.value.filter(node => node.id !== selectedStart.value)
);


function toggleAddingMode(mode: 'node' | 'edge') {
  const togglingNode = mode === 'node';
  const togglingEdge = mode === 'edge';

  isAddingNode.value = togglingNode ? !isAddingNode.value : false;
  isAddingEdge.value = togglingEdge ? !isAddingEdge.value : false;

  selectedNodes.value = [];
  selectedEdges.value = [];
}

function clearGraph() {
  isClearingGraph.value = true;

  clearNodes();
  clearEdges();
  Object.keys(layouts.value.nodes).forEach((key) => delete layouts.value.nodes[key]);

  isClearingGraph.value = false;
}

async function createRandomGraph() {
  isCreatingRandomGraph.value = true;

  const generatedGraph = await RandomGraphJS(nodePrefix.value, 20, 25, 500, 500);
  setNodes(generatedGraph.nodes);
  setEdges(generatedGraph.edges);
  layouts.value = generatedGraph.layouts;

  isCreatingRandomGraph.value = false;
}
console.log(getNodes)
console.log(getEdges)

onMounted(async () => {
  await nextTick()
  if (sidePanel.value) {
    panelWidth.value = sidePanel.value.getBoundingClientRect().width
  }
});

watch([selectedStart, selectedEnd], ([start, end]) => {
  if (start && end && start === end) {
    selectedEnd.value = '';
  }
});
</script>

<template>
  <div class="fixed top-0 right-0 bottom-0 m-4 flex items-start gap-2 transition-all duration-400"
    :style="{ transform: isExpanded ? 'translateX(0)' : `translateX(${panelWidth + 16}px)` }">
    <div class="h-4/7 flex flex-col justify-between gap-2">

      <Button variant="secondary" @click="isExpanded = !isExpanded">
        <PanelRightClose v-if="isExpanded" />
        <PanelRightOpen v-else />
      </Button>

      <div class="flex flex-col gap-2">
        <Button @click="toggleAddingMode('node')" :variant="isAddingNode ? 'default' : 'secondary'">
          <StopCircleIcon v-if="isAddingNode" />
          <CirclePlusIcon v-else />
        </Button>
        <Button @click="toggleAddingMode('edge')" :variant="isAddingEdge ? 'default' : 'secondary'">
          <StopCircleIcon v-if="isAddingEdge" />
          <SplineIcon v-else />
        </Button>
        <Button @click="createRandomGraph()" :disabled="isCreatingRandomGraph" variant="secondary" class="mt-12">
          <Loader2 v-if="isCreatingRandomGraph" class="w-4 h-4 animate-spin" />
          <ShuffleIcon v-else />
        </Button>
        <Button @click="isClearingGraphDialogOpen = true" :disabled="Object.keys(getNodes).length == 0"
          variant="secondary" class="mt-12">
          <Loader2 v-if="isClearingGraph" class="w-4 h-4 animate-spin" />
          <BrushCleaningIcon v-else />
        </Button>
      </div>

    </div>


    <div ref="sidePanel" class="w-xs md:w-md h-full">
      <Tabs default-value="tableau" class="h-full">

        <TabsList class="grid w-full grid-cols-2">
          <TabsTrigger value="tableau">
            Tableau
          </TabsTrigger>
          <TabsTrigger value="resolution">
            Résolution
          </TabsTrigger>
        </TabsList>

        <TabsContent value="tableau">
          <Card class="h-full">
            <CardHeader>
              <CardTitle>Tableau représentatif du graphe</CardTitle>
              <CardDescription>Arcs et sommets</CardDescription>
            </CardHeader>
            <CardContent>
              <div class="max-h-110 relative overflow-auto">
                <NodeTable v-if="Object.keys(getNodes).length != 0" :nodes="getNodes" :edges="getEdges" />
                <div v-else class="flex items-center justify-center h-12 border">
                  <code>vide...</code>
                </div>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="resolution">
          <Card class="h-full">
            <CardHeader>
              <CardTitle>Résolution</CardTitle>
              <CardDescription>Par l'algorithme de Dijkstra</CardDescription>
            </CardHeader>
            <CardContent class="h-1/4">
              <div class="flex flex-col gap-y-4">
                <div class="flex flex-col md:flex-row gap-4 justify-between">
                  <Select v-model:model-value="selectedStart" :items="filteredStartNodes" label="Début"
                    placeholder="Choisir..." />
                  <Select v-model:model-value="selectedEnd" :items="filteredEndNodes" label="Fin"
                    placeholder="Choisir..." />
                </div>
                <div class="flex flex-row w-full gap-4 justify-evenly">
                  <template v-if="!selectedStart || !selectedEnd">
                    <Button disabled>
                      Chemin optimal
                      <WaypointsIcon />
                    </Button>
                    <Button disabled variant="secondary">
                      Chemin maximal
                      <RouteIcon />
                    </Button>
                  </template>
                  <template v-else>
                    <Button as-child>
                      <RouterLink :to="`/resolution?start=${selectedStart}&end=${selectedEnd}`">
                        Chemin optimal
                        <WaypointsIcon />
                      </RouterLink>
                    </Button>
                    <Button variant="secondary" as-child>
                      <RouterLink :to="`/resolution-max?start=${selectedStart}&end=${selectedEnd}`">
                        Chemin maximal
                        <RouteIcon />
                      </RouterLink>
                    </Button>
                  </template>
                </div>
                <Alert class="mb-4">
                  <InfoIcon class="w-4 h-4" />
                  <AlertTitle>NP Problème</AlertTitle>
                  <AlertDescription>
                    Pour trouver le chemin le plus long, l'algorithme explore tous les chemins possibles 
                    depuis le point de départ. Il ne s'arrête que lorsque tous les noeuds accessibles 
                    ont été traités, ce qui est similaire au problème du voyageur de commerce.
                    Ce problème ne peut être résolu que dans un graphe pondéré positivement, acyclique et directionnel.
                  </AlertDescription>
                </Alert>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

      </Tabs>
    </div>

  </div>

  <ModalDialog v-model:open="isClearingGraphDialogOpen" :onConfirm="clearGraph">
    <template #dialog-title>Attention</template>
    <template #dialog-description>Vous allez supprimer tous les éléments du graphe</template>
  </ModalDialog>

</template>
