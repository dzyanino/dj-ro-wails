<script setup lang="ts">
import { onMounted, shallowRef, nextTick } from 'vue'
import { CirclePlusIcon, PanelRightClose, PanelRightOpen, ShuffleIcon, SplineIcon } from 'lucide-vue-next'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'
import Tooltip from './Tooltip.vue'

const sidePanel = shallowRef<HTMLElement | null>(null)
const panelWidth = shallowRef<number>(384)
const isExpanded = shallowRef<boolean>(false)

onMounted(async () => {
  await nextTick()
  if (sidePanel.value) {
    panelWidth.value = sidePanel.value.getBoundingClientRect().width
    console.log(`Side panel width: ${panelWidth.value}px`)
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
        <Tooltip>
          <template #trigger><Button variant="secondary"><CirclePlusIcon /></Button></template>
          <template #content>Ajouter sommet</template>
        </Tooltip>
        <Tooltip>
          <template #trigger><Button variant="secondary"><SplineIcon /></Button></template>
          <template #content>Ajouter arc</template>
        </Tooltip>
        <Tooltip>
          <template #trigger><Button variant="secondary" class="mt-4"><ShuffleIcon /></Button></template>
          <template #content>Aléatoire</template>
        </Tooltip>
      </div>

    </div>

    <div ref="sidePanel" class="w-xs md:w-sm h-full">
      <Card class="h-full bg-background dark:bg-background rounded-md">
        <CardHeader>
          <CardTitle>Card Title</CardTitle>
          <CardDescription>Card Description</CardDescription>
        </CardHeader>
        <CardContent>
          Card Content
        </CardContent>
        <CardFooter>
          Card Footer
        </CardFooter>
      </Card>
    </div>
  </div>
</template>
