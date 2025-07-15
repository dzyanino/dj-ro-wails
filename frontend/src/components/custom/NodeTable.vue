<script lang="ts" setup>
import { computed, type PropType } from 'vue';
import type { EdgeWithId, NodeWithId } from 'v-network-graph';
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';

const { nodes, edges } = defineProps({
  nodes: {
    type: Object as PropType<Record<string, NodeWithId>>,
    required: true
  },
  edges: {
    type: Object as PropType<Record<string, EdgeWithId>>,
    required: true
  }
});

// Cache node IDs to prevent reactive loops
// const nodeIds = computed(() => Object.keys(nodes).sort());
const nodeIds = computed(() =>
  Object.keys(nodes).sort((a, b) => {
    const getNumber = (id: string) =>
      Number(nodes[id].name?.replace(/\D+/g, '')) // Extract number from name like "Node 1"
    return getNumber(a) - getNumber(b)
  })
)


// Create edge lookup map (using node IDs)
const edgeMap = computed(() => {
  const map: Record<string, Record<string, string>> = {};
  Object.values(edges).forEach(edge => {
    map[edge.source] = map[edge.source] || {};
    map[edge.source][edge.target] = edge.label || '';
  });
  return map;
});

// Helper function to get edge label
function getEdgeLabel(sourceId: string, targetId: string): string {
  return edgeMap.value[sourceId]?.[targetId] || '';
}
</script>

<template>
  <Table class="border">
    <TableCaption />
    <TableHeader class="sticky top-0 bg-background dark:bg-background">
      <TableRow>
        <TableHead />
        <TableHead v-for="id in nodeIds" :key="id" class="border text-center">
          {{ nodes[id].name }}
        </TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <TableRow v-for="sourceId in nodeIds" :key="sourceId">
        <TableCell class="border font-medium">
          {{ nodes[sourceId].name }}
        </TableCell>
        <TableCell v-for="targetId in nodeIds" :key="targetId" class="border text-center">
          {{ getEdgeLabel(sourceId, targetId) }}
        </TableCell>
      </TableRow>
    </TableBody>
  </Table>
</template>