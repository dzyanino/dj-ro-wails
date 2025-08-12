<script lang="ts" setup>
import { defineProps, computed } from "vue";
import { types } from "../../../../wailsjs/go/models";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

const props = defineProps<{
  nodes: types.ResolutionNode[];
  startNodeId: string;
  endNodeId: string;
}>();

const sortedNodes = computed(() => {
  const startNode = props.nodes.find((n) => n.id === props.startNodeId);
  const endNode = props.nodes.find((n) => n.id === props.endNodeId);
  const others = props.nodes.filter(
    (n) => n.id !== props.startNodeId && n.id !== props.endNodeId
  );

  // Sort others naturally (node1, node2, node10 instead of node1, node10, node2)
  others.sort((a, b) => {
    return naturalCompare(a.id, b.id);
  });

  const result = [];
  if (startNode) result.push(startNode);
  result.push(...others);
  if (endNode) result.push(endNode);

  return result;
});

// Natural comparison function for proper numeric sorting
function naturalCompare(a: string, b: string): number {
  const ax: (string | number)[] = [];
  const bx: (string | number)[] = [];

  a.replace(/(\d+)|(\D+)/g, (_, $1, $2) => {
		//@ts-ignore
    ax.push([$1 || Infinity, $2 || ""]);
    return "";
  });
  b.replace(/(\d+)|(\D+)/g, (_, $1, $2) => {
		//@ts-ignore
    bx.push([$1 || Infinity, $2 || ""]);
    return "";
  });

  while (ax.length && bx.length) {
		//@ts-ignore
    const an = ax.shift() as [string | number, string];
		//@ts-ignore
    const bn = bx.shift() as [string | number, string];
    const nn = (an[0] as number) - (bn[0] as number) || an[1].localeCompare(bn[1]);
    if (nn) return nn;
  }

  return ax.length - bx.length;
}

const nodeIds = computed(() => sortedNodes.value.map((n) => n.id));

const stepIndices = computed(() => {
  const stepsSet = new Set<number>();
  for (const node of props.nodes) {
    for (const step of Object.keys(node.nodePropsList)) {
      stepsSet.add(Number(step));
    }
  }
  return Array.from(stepsSet).sort((a, b) => a - b);
});

function getCellData(
  node: types.ResolutionNode,
  step: number
): { weightText: string; previousNode: string; className: string } {
  // Check if this column has been marked at any previous step
  const hasBeenMarkedBefore = Object.entries(node.nodePropsList).some(
    ([stepStr, props]) => props.marked && Number(stepStr) < step
  );

  // Find the exact step data
  const currentStepProps = node.nodePropsList[step];
  
  if (currentStepProps) {
    // We have data for this exact step
    let weightText = currentStepProps.weightTo < 0 ? "∞" : `${currentStepProps.weightTo}`;
    let previousNode = currentStepProps.previousNode || "";
    let className = "";

    if (currentStepProps.marked) {
      // This cell is marked - yellow background
      className = "bg-yellow-300 dark:bg-yellow-500";
    } else if (hasBeenMarkedBefore) {
      // Column was marked before this step - cell should be grayed out
      className = "bg-gray-300 dark:bg-gray-500 text-gray-600 dark:text-gray-400";
    } else if (!currentStepProps.valid && currentStepProps.weightTo < 0) {
      // Invalid cell with negative weight (initial infinity)
      className = "text-gray-400";
    } else if (!currentStepProps.valid) {
      className = "bg-gray-300 dark:bg-gray-500 text-gray-600 dark:text-gray-400";
    }

    return { weightText, previousNode, className };
  } else {
    // No data for this step
    if (hasBeenMarkedBefore) {
      // Column was marked before - show empty grayed cell
      return { weightText: "", previousNode: "", className: "bg-gray-300 dark:bg-gray-500" };
    } else {
      // Find most recent valid data
      const steps = Object.keys(node.nodePropsList)
        .map(Number)
        .filter((s) => s < step)
        .sort((a, b) => b - a); // Sort descending
      
      if (steps.length === 0) {
        return { weightText: "∞", previousNode: "", className: "text-gray-400" };
      }

      const lastStep = steps[0];
      const lastProps = node.nodePropsList[lastStep];
      
      if (!lastProps || lastProps.weightTo < 0) {
        return { weightText: "∞", previousNode: "", className: "text-gray-400" };
      }

      let weightText = `${lastProps.weightTo}`;
      let previousNode = lastProps.previousNode || "";
      
      return { weightText, previousNode, className: "text-gray-400" };
    }
  }
}
</script>

<template>
    <Table
      class="border-collapse border border-slate-400 w-full text-sm text-center"
    >
      <TableHeader class="sticky top-0 bg-background dark:bg-background z-10">
        <TableRow>
          <!-- <TableHead class="border border-slate-300">Étape</TableHead> -->
          <TableHead
            v-for="id in nodeIds"
            :key="id"
            class="border border-slate-300"
          >
            {{ id.toUpperCase() }}
          </TableHead>
        </TableRow>
      </TableHeader>
  
      <TableBody>
        <TableRow v-for="step in stepIndices" :key="step">
          <!-- <TableCell class="border border-slate-300 font-medium">{{
            step
          }}</TableCell> -->
          <TableCell
            v-for="node in sortedNodes"
            :key="node.id"
            :class="
              getCellData(node, step).className + ' border border-slate-300 relative'
            "
          >
            <div class="inline-block">
              {{ getCellData(node, step).weightText }}
              <sub 
                v-if="getCellData(node, step).previousNode" 
                class="text-xs opacity-75"
              >
                {{ getCellData(node, step).previousNode.toUpperCase() }}
              </sub>
            </div>
          </TableCell>
        </TableRow>
      </TableBody>
    </Table>
</template>