<script setup lang="ts">
import { inject, shallowRef, useTemplateRef, type ShallowRef } from 'vue';
import { onClickOutside } from '@vueuse/core';
import { useColorMode } from '@vueuse/core'
import { Button } from '@/components/ui/button'
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from '@/components/ui/dropdown-menu'
import { Icon } from '@iconify/vue'

const mode = useColorMode({ disableTransition: false })

const isVisible = inject<ShallowRef<boolean>>('isVisible', shallowRef<boolean>(false))
const isThemeSwitcherShown = inject<ShallowRef<boolean>>('isThemeSwitcherShown', shallowRef<boolean>(false))

const themeSwitcher = useTemplateRef<HTMLElement>('target')
onClickOutside(themeSwitcher, () => {
    isThemeSwitcherShown.value = false;
    isVisible.value = false;
})
</script>

<template>
    <DropdownMenu>
        <DropdownMenuTrigger as-child ref="target">
            <Button variant="ghost" @click="isThemeSwitcherShown = true">
                <Icon icon="radix-icons:moon"
                    class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
                <Icon icon="radix-icons:sun"
                    class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
                <span class="sr-only">Changer thème</span>
            </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end" class="bg-card mt-1">
            <DropdownMenuItem @click="mode = 'light'">
                Clair
            </DropdownMenuItem>
            <DropdownMenuItem @click="mode = 'dark'">
                Sombre
            </DropdownMenuItem>
            <DropdownMenuItem @click="mode = 'auto'">
                Système
            </DropdownMenuItem>
        </DropdownMenuContent>
    </DropdownMenu>
</template>