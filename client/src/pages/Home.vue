<script setup lang="ts">
	import CoffeeList from "/src/components/CoffeeList.vue";
	import {inject, ref} from "vue";
	const products = inject("products", ref([]))
</script>

<template>
	<h1 class="text-center text-4xl mb-6">Out Coffee products</h1>
	<CoffeeList :products="products" />
</template>

<style scoped>

</style>