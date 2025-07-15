<script setup lang="ts">
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
</script>

<template>
  <TooltipProvider>
    <Tooltip>
      <TooltipTrigger as-child>
        <slot name="trigger"></slot>
      </TooltipTrigger>
      <TooltipContent>
        <slot name="content"></slot>
      </TooltipContent>
    </Tooltip>
  </TooltipProvider>
</template>
