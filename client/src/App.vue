<script setup lang="ts">
	import NavBar from "@/shared/NavBar.vue";
	import { onMounted, provide, ref } from "vue";

	const products = ref([]);

	onMounted(async () => {
		const res = await fetch("http://localhost:8080/products");
		products.value = await res.json();
	});

	provide('products', products);
</script>

<template>
	<NavBar />
    <main class="max-w-5xl mx-auto my-12">
	    <router-view></router-view>
    </main>
</template>

<style scoped>

</style>
