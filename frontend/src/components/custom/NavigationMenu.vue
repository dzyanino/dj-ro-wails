<script setup lang="ts">
import { Button } from '@/components/ui/button';
import { NavigationMenu, NavigationMenuItem, NavigationMenuLink, NavigationMenuList, navigationMenuTriggerStyle } from '@/components/ui/navigation-menu'
import { Separator } from '@/components/ui/separator';
import ThemeSwitcher from './ThemeSwitcher.vue';

const navLinks = [
    { to: '#/data', title: 'Résoudre' },
    { to: '#/aide', title: 'Aide' },
]
</script>

<template>
    <NavigationMenu>
        <NavigationMenuList class="flex gap-2 h-8">
            <NavigationMenuItem>
                <Button variant="link">
                    <RouterLink to="/">Ro-Dijkstra</RouterLink>
                </Button>
            </NavigationMenuItem>
            <Separator orientation="vertical" class="mx-2" />
            <NavigationMenuItem v-for="navLink in navLinks" :key="navLink.title">
                <NavigationMenuLink :href="navLink.to" :class="navigationMenuTriggerStyle()">
                    {{ navLink.title }}
                </NavigationMenuLink>
            </NavigationMenuItem>
            <Separator orientation="vertical" class="mx-2" />
            <NavigationMenuItem>
                <ThemeSwitcher />
            </NavigationMenuItem>
        </NavigationMenuList>
    </NavigationMenu>
</template>