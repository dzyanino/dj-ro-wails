<script setup lang="ts">
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { Button } from '@/components/ui/button';

defineProps<{
  open: boolean
  onConfirm?: () => void
  onCancel?: () => void
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
}>()
</script>

<template>
  <Dialog :open="open" @update:open="emit('update:open', $event)">
    <DialogTrigger as-child v-if="$slots.trigger">
      <slot name="trigger" />
    </DialogTrigger>

    <DialogContent class="sm:max-w-[425px]">
      <DialogHeader>
        <DialogTitle>
          <slot name="dialog-title" />
        </DialogTitle>
        <DialogDescription>
          <slot name="dialog-description" />
        </DialogDescription>
      </DialogHeader>

      <div v-if="$slots['dialog-content']" class="grid gap-4 py-4">
        <slot name="dialog-content" />
      </div>

      <DialogFooter>
        <slot name="dialog-footer">
          <Button @click="() => { emit('update:open', false); onCancel?.() }" variant="secondary">
            Annuler
          </Button>
          <Button @click="() => { emit('update:open', false); onConfirm?.() }">
            Confirmer
          </Button>
        </slot>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
