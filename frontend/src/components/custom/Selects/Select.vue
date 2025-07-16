<script setup lang="ts">
import type { PropType } from 'vue';
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import type { NodeWithId } from 'v-network-graph';

defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  placeholder: {
    type: String,
    default: 'Choisir une option'
  },
  label: {
    type: String,
    default: ''
  },
  items: {
    type: Array as PropType<NodeWithId[]>,
    required: true
  },
  width: {
    type: String,
    default: 'w-full'
  }
})

const emit = defineEmits<{
  (e: "update:modelValue", value: string | null): void;
}>();
</script>

<template>
  <Select :modelValue="modelValue" @update:modelValue="emit('update:modelValue', $event as string | null)">
    <SelectTrigger :disabled="items.length == 0" :class="width">
      <SelectValue :placeholder="placeholder" />
    </SelectTrigger>
    <SelectContent>
      <SelectGroup>
        <SelectLabel v-if="label">{{ label }}</SelectLabel>
        <SelectItem v-for="item in items" :key="item.id" :value="item.id">
          {{ item.name }}
        </SelectItem>
      </SelectGroup>
    </SelectContent>
  </Select>
</template>
