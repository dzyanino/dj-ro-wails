<script lang="ts" setup>
import { onBeforeMount, onBeforeUnmount, provide, ref, shallowRef } from 'vue';
import type { Layouts } from 'v-network-graph';
import FloatingNavBar from '@/components/layouts/FloatingNavBar.vue';
import FloatingControls from '@/components/custom/FloatingControls.vue';
import NetworkGraph from '@/components/custom/NetworkGraph.vue';
import isMobile from '@/utils/isMobile';

const mobile = shallowRef<boolean>(isMobile());

const isAddingNode = shallowRef<boolean>(false);
const layouts = ref<Layouts>({ nodes: {} });

provide('isAddingNode', isAddingNode);
provide('layouts', layouts);


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
    <FloatingControls />

    <main class="flex w-full h-dvh p-2">
        <div class="flex flex-1 flex-col gap-4">
            <NetworkGraph class="h-full" />
        </div>
    </main>
</template>