<script lang="ts" setup>
import { computed, onMounted, ref, shallowRef } from "vue";
import { useRoute } from "vue-router";
import { useColorMode } from "@vueuse/core";
import { useNodeStore } from "@/stores/nodes";
import { useEdgeStore } from "@/stores/edges";
import { types } from "../../wailsjs/go/models";
import {
  InitializeNodeArray,
  Step,
  ReconstructPath,
} from "../../wailsjs/go/services/LongestPath";
import { VNetworkGraph, VEdgeLabel, type UserConfigs } from "v-network-graph";
import { Button } from "@/components/ui/button";
import ResolutionMaxTable from "@/components/custom/ResolutionMaxPage/ResolutionMaxTable.vue";
import {
  ArrowRight,
  CheckIcon,
  AlertTriangle,
} from "lucide-vue-next";
import { Badge } from "@/components/ui/badge";
import { Separator } from "@/components/ui/separator";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { createGraphConfig } from "@/utils/vNetworkGraphConfigs";
import FloatingNavBar from "@/components/layouts/FloatingNavBar.vue";

const route = useRoute();

const mode = useColorMode();
const theme = computed(() =>
  mode.value === "auto"
    ? window.matchMedia("(prefers-color-scheme: dark)").matches
      ? "dark"
      : "light"
    : (mode.value as "dark" | "light")
);

// Use directed graph config for longest path
const configs = computed<UserConfigs>(() =>
  createGraphConfig(theme.value, true, true, true)
);

const { getNodes } = useNodeStore();
const { getEdges } = useEdgeStore();

const nodes = ref(getNodes);
const edges = ref(getEdges);

const startingNodeID = route.query.start as string;
const endingNodeID = route.query.end as string;

const currentStep = shallowRef<number>(0);
const nodeArray = ref<types.ResolutionNode[]>([]);

const markedNodes = ref<string[]>([]);
const graphMarkedNodes = ref<string[]>([]);
const currentNode = shallowRef<string>("");
const finished = shallowRef<boolean>(false);
const path = ref<string[]>([]);
const networkPath = ref({});
const errorMessage = ref<string>("");

function buildEdgePathFromNodes(nodePath: string[]): string[] {
  const edgeIds: string[] = [];

  for (let i = 0; i < nodePath.length - 1; i++) {
    const from = nodePath[i];
    const to = nodePath[i + 1];

    const edgeEntry = Object.entries(edges.value).find(([_, edge]) => {
      // For directed graphs, only match source->target direction
      return edge.source === from && edge.target === to;
    });

    if (edgeEntry) {
      edgeIds.push(edgeEntry[0]); // key = ID de l'edge
    }
  }

  return edgeIds;
}

// Check if the graph has cycles (basic DFS-based cycle detection)
function hasCycle(): boolean {
  const visited = new Set<string>();
  const recursionStack = new Set<string>();

  function dfsHasCycle(nodeId: string): boolean {
    visited.add(nodeId);
    recursionStack.add(nodeId);

    // Find outgoing edges from this node
    for (const edge of Object.values(edges.value)) {
      if (edge.source === nodeId) {
        const neighbor = edge.target;

        if (!visited.has(neighbor)) {
          if (dfsHasCycle(neighbor)) {
            return true;
          }
        } else if (recursionStack.has(neighbor)) {
          return true; // Back edge found - cycle detected
        }
      }
    }

    recursionStack.delete(nodeId);
    return false;
  }

  // Check all nodes as potential starting points
  for (const nodeId of Object.keys(nodes.value)) {
    if (!visited.has(nodeId)) {
      if (dfsHasCycle(nodeId)) {
        return true;
      }
    }
  }

  return false;
}

async function nextStep() {
  if (finished.value) return;

  try {
    const res = await Step(
      nodes.value,
      edges.value,
      nodeArray.value,
      currentStep.value
    );

    nodeArray.value = res.nodeArray;
    markedNodes.value = res.markedNodes;
    graphMarkedNodes.value = res.markedNodes;
    currentNode.value = res.currentNode;
    finished.value = res.finished;

    currentStep.value++;

    if (finished.value) {
      path.value = await ReconstructPath(nodeArray.value);
      const orderedEdges = buildEdgePathFromNodes(path.value);

      networkPath.value = {
        path1: { edges: orderedEdges },
      };
    }

    errorMessage.value = "";
  } catch (error) {
    //@ts-ignore
    errorMessage.value = error.message || "L'algorithme s'est arrêté";
  }
}

// Calculate total path weight
const totalPathWeight = computed(() => {
  if (path.value.length < 2) return 0;

  let total = 0;
  for (let i = 0; i < path.value.length - 1; i++) {
    const from = path.value[i];
    const to = path.value[i + 1];

    const edge = Object.values(edges.value).find(
      (e) => e.source === from && e.target === to
    );

    if (edge && edge.label) {
      total += parseInt(edge.label) || 0;
    }
  }
  return total;
});

onMounted(async () => {
  // Check for cycles and warn user
  if (hasCycle()) {
    errorMessage.value =
      "Le graphe contient des cycles. Pour éviter les boucles infinies, nous allons interdire les passages vers les sommets déjà visités.";
  }

  nodeArray.value = await InitializeNodeArray(
    nodes.value,
    startingNodeID,
    endingNodeID
  );
});
</script>

<template>
  <FloatingNavBar />
  <main>
    <div class="min-h-screen h-screen flex">
      <div class="flex flex-col w-1/2 h-full overflow-scroll gap-y-4 p-8">
        <div class="flex justify-between items-center">
          <h1 class="text-lg font-bold">Chemin le plus long</h1>
          <Button :disabled="finished" @click="nextStep" class="max-w-fit">
            Prochaine étape
            <ArrowRight />
          </Button>
        </div>

        <Separator />

        <!-- Warning about directed graphs and cycles -->
        <Alert v-if="errorMessage" class="mb-4 text-amber-600">
          <AlertTriangle class="w-4 h-4" />
          <AlertTitle>Attention</AlertTitle>
          <AlertDescription class="text-amber-600">
            {{ errorMessage }}
          </AlertDescription>
        </Alert>

        <div class="flex gap-x-4 items-center justify-start text-xs">
          <p>
            Étape actuelle: <Badge variant="outline">{{ currentStep }}</Badge>
          </p>
          <p>
            Noeud actuel:
            <Badge v-if="!!currentNode" variant="outline">{{
              currentNode
            }}</Badge>
            <Badge v-else variant="outline">-</Badge>
          </p>
        </div>

        <div class="min-h-fit overflow-auto">
          <div>
            <ResolutionMaxTable
              :start-node-id="startingNodeID"
              :end-node-id="endingNodeID"
              :nodes="nodeArray"
              class="w-full h-full"
            />
          </div>
        </div>

        <Separator />

        <div v-if="finished && path.length > 0">
          <Alert>
            <CheckIcon class="w-4 h-4" />
            <AlertTitle
              >Exploration terminée - Chemin le plus long trouvé</AlertTitle
            >
            <AlertDescription>
              <div class="space-y-2">
                <div class="flex flex-wrap gap-x-2 items-center">
                  <template v-for="(node, index) in path" :key="node">
                    <code>{{ node }}</code>
                    <code v-if="index < path.length - 1" class="text-xl"
                      >→</code
                    >
                  </template>
                </div>
                <p class="text-sm font-medium">
                  Poids total:
                  <Badge variant="secondary">{{ totalPathWeight }}</Badge>
                </p>
                <p class="text-xs text-muted-foreground">
                  Tous les chemins possibles ont été explorés.
                </p>
              </div>
            </AlertDescription>
          </Alert>
        </div>

        <div v-else-if="finished && path.length === 0">
          <Alert variant="destructive">
            <AlertTriangle class="w-4 h-4" />
            <AlertTitle>Exploration terminée - Aucun chemin trouvé</AlertTitle>
            <AlertDescription>
              Il n'existe pas de chemin du noeud {{ startingNodeID }} vers
              {{ endingNodeID }}
              dans ce graphe dirigé. Tous les noeuds accessibles ont été
              explorés.
            </AlertDescription>
          </Alert>
        </div>
      </div>

      <div class="grow">
        <v-network-graph
          v-model:selected-nodes="graphMarkedNodes"
          :nodes="nodes"
          :edges="edges"
          :paths="networkPath"
          :configs="configs"
          :zoom-level="3"
          tabindex="0"
          @click="
            graphMarkedNodes.length <= 0
              ? (graphMarkedNodes = markedNodes)
              : null
          "
          class="h-full graph border rounded-lg bg-background dark:bg-background"
        >
          <template #edge-label="{ edge, ...slotProps }">
            <v-edge-label
              :text="edge.label"
              align="center"
              vertical-align="above"
              v-bind="slotProps"
            ></v-edge-label>
          </template>
        </v-network-graph>
      </div>
    </div>
  </main>
</template>
