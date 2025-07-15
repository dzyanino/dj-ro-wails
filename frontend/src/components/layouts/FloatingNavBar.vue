<script setup lang="ts">
import { provide, shallowRef } from 'vue'
import { Button } from '@/components/ui/button'
import { ChevronDown } from 'lucide-vue-next'
import { Card, CardContent } from '../ui/card';
import NavigationMenu from '@/components/custom/NavigationMenu.vue';

const isVisible = shallowRef<boolean>(false)
const isThemeSwitcherShown = shallowRef<boolean>(false)
provide('isVisible', isVisible)
provide('isThemeSwitcherShown', isThemeSwitcherShown)

let hideTimeout: number | null = null

const showNav = (show: boolean) => {
    if (hideTimeout) {
        clearTimeout(hideTimeout)
        hideTimeout = null
    }
    isVisible.value = show
}

const scheduleHide = () => {
    if (!isThemeSwitcherShown.value) {
        hideTimeout = window.setTimeout(() => {
            isVisible.value = false
        }, 300)
    }
}
</script>

<template>
    <nav>
        <div class="fixed top-0 left-1/2 -translate-x-1/2 z-50">
            <Button variant="outline" @mouseenter="showNav(true)" @mouseleave="scheduleHide()"
                class="mt-2 transition-transform">
                <slot name="icon">
                    <ChevronDown class="transition-transform duration-300" :class="{ 'rotate-180': isVisible }" />
                </slot>
            </Button>
    
        </div>
        <div class="fixed top-0 left-1/2 -translate-x-1/2 z-50 mt-2">
            <Transition name="slide">
                <Card v-show="isVisible" @mouseenter="showNav(true)" @mouseleave="scheduleHide()" class="py-2">
                    <CardContent class="px-2">
                        <slot>
                            <NavigationMenu />
                        </slot>
                    </CardContent>
                </Card>
            </Transition>
        </div>
    </nav>
</template>

<style scoped>
.slide-enter-active,
.slide-leave-active {
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.slide-enter-from,
.slide-leave-to {
    transform: translateY(-20px);
    opacity: 0;
}

.rotate-180 {
    transform: rotate(180deg);
}
</style>