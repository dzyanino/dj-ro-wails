<script setup>
import { ref } from 'vue';
import { RouterLink } from 'vue-router';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { ArrowRight, ListChecks, Presentation, Rocket } from 'lucide-vue-next';

const features = ref([
    {
        icon: ListChecks,
        iconBg: 'bg-gradient-to-br from-blue-500 to-indigo-600',
        title: 'Étape par étape',
        desc: 'Chaque étape de l\'algorithme en détail.',
    },
    {
        icon: Presentation,
        iconBg: 'bg-gradient-to-br from-purple-500 to-pink-500',
        title: 'Visuels explicatifs',
        desc: 'Animations !<br />Quoi de mieux.',
    },
    {
        icon: Rocket,
        iconBg: 'bg-gradient-to-br from-yellow-400 to-orange-500',
        title: 'Propulsé par Go',
        desc: 'La puissance du langage Go couplée à Vue.',
    },
]);
</script>

<template>
    <div class="min-h-screen flex flex-col items-center bg-background">
        <header class="text-center mt-24">
            <img src="/src/assets/ro-wails.svg" alt="Ro-Wails Logo" class="w-28 h-28 mx-auto mb-6" />
            <h1 class="text-4xl font-extrabold mb-2">
                Bienvenue à
                <span class="bg-gradient-to-r from-primary to-chart-4 bg-clip-text text-transparent">Ro-Dijkstra</span>
            </h1>
            <p class="text-lg text-muted-foreground mb-6">
                Résolution de l'algorithme du chemin le plus court, avec des visuels.
            </p>
            <div class="flex justify-center gap-6">
                <Button>
                    <RouterLink to="/data">Commencer</RouterLink>
                    <ArrowRight />
                </Button>
                <Button variant="secondary">
                    <RouterLink to="/aide">Aide</RouterLink>
                </Button>
            </div>
        </header>
        <section class="flex flex-wrap justify-center gap-8 my-14">
            <Card class="w-64" v-for="(feature, i) in features" :key="i">
                <CardContent>
                    <div class="text-center flex flex-col items-center">
                        <div class="w-14 h-14 rounded-full flex items-center justify-center mb-4 text-2xl"
                            :class="feature.iconBg">
                            <component :is="feature.icon" class="w-8 h-8 text-white" />
                        </div>
                        <h2 class="text-lg font-bold mb-2">{{ feature.title }}</h2>
                        <p class="text-muted-foreground" v-html="feature.desc"></p>
                    </div>
                </CardContent>
            </Card>
        </section>
    </div>
</template>