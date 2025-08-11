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
} from "../../wailsjs/go/services/Dijkstra";
import { VNetworkGraph, VEdgeLabel } from "v-network-graph";
import { Button } from "@/components/ui/button";
import ResolutionTable from "@/components/custom/ResolutionPage/ResolutionTable.vue";
import { ArrowRight, CheckIcon } from "lucide-vue-next";
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

const configs = createGraphConfig(theme.value, false, true);

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

function buildEdgePathFromNodes(nodePath: string[]): string[] {
  const edgeIds: string[] = [];

  for (let i = 0; i < nodePath.length - 1; i++) {
    const from = nodePath[i];
    const to = nodePath[i + 1];

    const edgeEntry = Object.entries(edges.value).find(([_, edge]) => {
      return (
        (edge.source === from && edge.target === to) ||
        (edge.source === to && edge.target === from) // si graphe non orienté
      );
    });

    if (edgeEntry) {
      edgeIds.push(edgeEntry[0]); // key = ID de l'edge
    }
  }

  return edgeIds;
}

async function nextStep() {
  // console.log(nodeArray.value)

  if (finished.value) return;

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
}

onMounted(async () => {
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
		<div class="min-h-screen h-screen flex gap-4">
			<div class="flex flex-col w-1/2 gap-y-4 p-8">
				<div class="flex justify-end">
					<Button :disabled="finished" @click="nextStep" class="max-w-fit">
						Prochaine étape
						<ArrowRight />
					</Button>
				</div>
				<Separator />
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
					<!-- <div class="flex gap-x-4 items-center">
								<p>Noeuds marqués:</p>
								<div class="flex gap-x-1 items-center">
									<Badge v-for="markedNode in markedNodes" :key="markedNode">{{
										markedNode
									}}</Badge>
								</div>
							</div> -->
				</div>
	
				<div class="overflow-auto">
					<ResolutionTable
						:start-node-id="startingNodeID"
						:end-node-id="endingNodeID"
						:nodes="nodeArray"
						class="w-full h-full"
					/>
				</div>
	
				<Separator />
	
				<div v-if="finished">
					<Alert>
						<CheckIcon class="w-4 h-4" />
						<AlertTitle>Chemin optimal</AlertTitle>
						<AlertDescription>
							<div class="flex gap-x-2 items-center">
								<template v-for="(node, index) in path" :key="node">
									<code>{{ node }}</code>
									<code v-if="index < path.length - 1" class="text-xl">→</code>
								</template>
							</div>
						</AlertDescription>
					</Alert>
				</div>
			</div>
	
			<div class="grow">
				<v-network-graph
					ref="graph"
					v-model:selected-nodes="graphMarkedNodes"
					:nodes="nodes"
					:edges="edges"
					:paths="networkPath"
					:configs="configs"
					:zoom-level="3"
					tabindex="0"
					@click="
						graphMarkedNodes.length <= 0 ? (graphMarkedNodes = markedNodes) : null
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
