<script lang="ts" setup>
import { onBeforeMount, onBeforeUnmount, provide, shallowRef } from 'vue';
import { ResizableHandle, ResizablePanel, ResizablePanelGroup } from '@/components/ui/resizable';
import FloatingNavBar from '@/components/layouts/FloatingNavBar.vue';
import NetworkGraph from '@/components/custom/NetworkGraph.vue';
import RightControls from '@/components/custom/RightControls.vue';
import isMobile from '@/utils/isMobile';

const mobile = shallowRef<boolean>(isMobile());

const isAddingNode = shallowRef<boolean>(false);

provide('isAddingNode', isAddingNode);


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

    <main class="flex flex-col w-full h-screen p-2">

        <div class="flex flex-1 flex-col gap-4">
            <ResizablePanelGroup :direction="mobile ? 'vertical' : 'horizontal'">

                <ResizablePanel :default-size="50">
                    <NetworkGraph class="h-full" />
                </ResizablePanel>

                <ResizableHandle with-handle />

                <ResizablePanel :default-size="50">
                    <RightControls />
                </ResizablePanel>

            </ResizablePanelGroup>
        </div>

    </main>
</template>