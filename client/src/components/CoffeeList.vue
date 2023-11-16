<template>
	<div class="relative overflow-x-auto">
		<table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
			<thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
				<tr>
					<th scope="col" class="px-6 py-3">
						name
					</th>
					<th scope="col" class="px-6 py-3">
						description
					</th>
					<th scope="col" class="px-6 py-3">
						price
					</th>
				</tr>
			</thead>
			<tbody>
				<tr v-for="product in products" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
					<td scope="row" class="px-6 py-4 font-medium text-green-500 underline hover:opacity-70 whitespace-nowrap dark:text-white">
						<router-link :to="`/products/${product.id}`">
							{{ product.name }}
						</router-link>
					</td>
					<td class="px-6 py-4">
						{{ product.description }}
					</td>
					<td class="px-6 py-4">
						${{ product.price }}
					</td>
				</tr>
			</tbody>
		</table>
	</div>
</template>

<script setup>
	defineProps(['products'])
</script>