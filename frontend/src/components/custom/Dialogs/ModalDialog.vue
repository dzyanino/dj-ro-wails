<script setup lang="ts">
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/components/ui/alert-dialog'

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
  <AlertDialog :open="open" @update:open="emit('update:open', $event)">
    <AlertDialogTrigger as-child v-if="$slots.trigger">
      <slot name="trigger" />
    </AlertDialogTrigger>

    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle>
          <slot name="dialog-title" />
        </AlertDialogTitle>
        <AlertDialogDescription>
          <slot name="dialog-description" />
        </AlertDialogDescription>
      </AlertDialogHeader>

      <div v-if="$slots['dialog-content']" class="grid gap-4 py-4">
        <slot name="dialog-content" />
      </div>

      <AlertDialogFooter>
        <slot name="dialog-footer">
          <AlertDialogCancel @click="() => { emit('update:open', false); onCancel?.() }">
            Annuler
          </AlertDialogCancel>

          <AlertDialogAction @click="() => { emit('update:open', false); onConfirm?.() }">
            Confirmer
          </AlertDialogAction>
        </slot>
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>
