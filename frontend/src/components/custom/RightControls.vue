<script lang="ts" setup>
import { inject, shallowRef, type ShallowRef } from "vue";
// import { useEdgeStore } from '@/stores/edges';
// import { useNodeStore } from '@/stores/nodes';
import { ShuffleIcon, PlusIcon, StopCircleIcon } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Separator } from "@/components/ui/separator";
import { Toggle } from "@/components/ui/toggle";

const isAddingNode = inject<ShallowRef<boolean>>(
  "isAddingNode",
  shallowRef<boolean>(false)
);

// const { getNodes } = useNodeStore();
// const { getEdges } = useEdgeStore();
</script>

<template>
  <div class="flex flex-1 flex-col gap-4 p-4">
    <div class="flex items-center justify-between gap-4">
      <span class="text-md font-bold"></span>
      <div class="flex justify-start gap-4">
        <Toggle v-model="isAddingNode" variant="outline">
          <StopCircleIcon v-if="isAddingNode" class="size-4" />
          <PlusIcon v-else class="size-4" />
          {{ isAddingNode ? "Stop" : "Sommets" }}
        </Toggle>
        <div class="flex items-center space-x-2">
          <Checkbox id="arcs-aleatoires" />
          <Label for="arcs-aleatoires">Arcs aléatoires</Label>
        </div>
        <Button>
          <ShuffleIcon />
          Aléatoire
        </Button>
      </div>
    </div>

    <!-- <div class="max-h-48 overflow-auto">
            <NodeTable v-if="Object.keys(getNodes).length != 0" :nodes="getNodes" :edges="getEdges" />
            <p v-else>Vide</p>
        </div> -->

    <Separator orientation="horizontal" class="my-4" />

    <div class="flex items-center justify-between gap-4">
      <span class="text-md font-bold"></span>
      <div class="flex justify-start gap-4">
        <Button>
          <PlusIcon class="size-4" />
          Arcs
        </Button>
        <Button>Retirer</Button>
      </div>
    </div>

    <Button>Test</Button>
    <Input type="text" placeholder="Ecrire ici..." />
  </div>
</template>
