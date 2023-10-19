<template>
<table class="table">
    <thead>
        <tr>
            <th scope="col">ID</th>
            <th scope="col">Nazwa Przedmiotu</th>
        </tr>
    </thead>
    <tbody>
        <tr v-for="subject in subjects" :key="subject.id">
            <th scope="row">{{ subject.id }}</th>
            <td>{{ subject.name }}</td>
        </tr>
    </tbody>
</table>
</template>
<script setup lang="ts">
import type { Subject } from '@/types/Subject'
defineProps<{ subjects: Subject[] }>()
</script>