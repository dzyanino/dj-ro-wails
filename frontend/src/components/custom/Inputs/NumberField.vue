<script setup lang="ts">
import { Label } from "@/components/ui/label";
import {
  NumberField,
  NumberFieldContent,
  NumberFieldDecrement,
  NumberFieldIncrement,
  NumberFieldInput,
} from "@/components/ui/number-field";

defineProps<{
  id?: string;
  label?: string;
  modelValue: number;
  min?: number;
  max?: number;
  step?: number;
}>();

const emit = defineEmits<{
  (e: "update:modelValue", value: number): void;
}>();
</script>

<template>
  <NumberField
    :id="id"
    :model-value="modelValue"
    :min="min"
    :max="max"
    :step="step"
    @update:model-value="emit('update:modelValue', $event)"
  >
    <Label v-if="label" :for="id">{{ label }}</Label>

    <NumberFieldContent>
      <NumberFieldDecrement />
      <NumberFieldInput />
      <NumberFieldIncrement />
    </NumberFieldContent>
  </NumberField>
</template>
