<template>
    <MainLayout>
        <SubjectList :subjects="subjects"/>
    </MainLayout>
</template>
<script lang="ts" setup>
import MainLayout from '@/layouts/MainLayout.vue'
import SubjectList from '@/components/SubjectList.vue'
import { ref } from 'vue';
import type {Subject}from '@/types/Subject'

const subjects = ref<Subject[]>([])

fetch(import.meta.env.VITE_API_BASE_URL + '/subject/all', {
    headers: {
    'Authorization': 'Bearer '+localStorage.getItem('token')
    }
}).then(res => res.json()).then(({ data }) => {
    subjects.value = data
}).then(console.log).catch(err => {
    alert(err)
})
</script>

<style></style>