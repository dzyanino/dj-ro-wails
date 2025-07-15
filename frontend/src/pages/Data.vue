<script lang="ts" setup>
import { computed, onBeforeMount, onBeforeUnmount, provide, ref, shallowRef } from 'vue';
import { useColorMode } from '@vueuse/core';
import type { Layouts } from 'v-network-graph';
import FloatingNavBar from '@/components/layouts/FloatingNavBar.vue';
import FloatingControls from '@/components/custom/FloatingControls.vue';
import NetworkGraph from '@/components/custom/NetworkGraph.vue';
import isMobile from '@/utils/isMobile';
import { createGraphConfig } from '@/utils/vNetworkGraphConfigs';

const mode = useColorMode();
const theme = computed(() =>
    mode.value === 'auto'
    ? (window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light')
    : mode.value as 'dark' | 'light'
);

const mobile = shallowRef<boolean>(isMobile());

const isAddingNode = shallowRef<boolean>(false);
const isAddingEdge = shallowRef<boolean>(false);
const layouts = ref<Layouts>({ nodes: {} });

const graphConfigs = computed(() => createGraphConfig(theme.value, isAddingEdge.value));

const selectedNodes = ref<string[]>([]);
const selectedEdges = ref<string[]>([]);

const isAddingEdgeDialogOpen = shallowRef<boolean>(false);

provide('configs', graphConfigs);

provide('isAddingNode', isAddingNode);
provide('isAddingEdge', isAddingEdge);
provide('layouts', layouts);

provide('selectedNodes', selectedNodes);
provide('selectedEdges', selectedEdges);

provide('isAddingEdgeDialogOpen', isAddingEdgeDialogOpen);


function onResize() {
    mobile.value = isMobile();
}

onBeforeMount(() => {
    addEventListener('resize', onResize);
});

onBeforeUnmount(() => {
    removeEventListener('resize', onResize);
});
</script>

<template>
    <FloatingNavBar />
    <main>
        <FloatingControls />
        
        <div class="flex w-full h-dvh p-2">
            <div class="flex flex-1 flex-col gap-4">
                <NetworkGraph />
            </div>
        </div>
    </main>

</template>