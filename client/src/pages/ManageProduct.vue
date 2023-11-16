<script setup lang="ts">
	import {ref} from "vue";
	import {useRoute} from "vue-router";
	const route = useRoute()
	const file = ref(null);
	const file_name = ref(null);

	async function handleSubmit(event) {
		if (file.value.files.length > 0) {
			const formData = new FormData();
			formData.append('file', file.value.files[0]);
			formData.append("file_name", file_name.value.value)

			console.log("Posting ", formData)
			try {
				const response = await fetch(`http://localhost:9090/images/${route.params.id}`, {
					body: formData,
					method: "POST"
				})

				console.log('File uploaded successfully', response);
			} catch (error) {
				console.error('Error during file upload', error);
			}
		}
	}
</script>

<template>
	<div>
		<h1 class="text-center text-4xl mb-6">Manage product</h1>
		<form @submit.prevent="handleSubmit" class="max-w-md mx-auto space-y-4">

			<div>
				<label for="file_name" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">File name</label>
				<input type="text" ref="file_name" name="file_name" id="file_name" aria-describedby="helper-text-explanation" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
			</div>

			<div>
				<label class="block mb-2 text-sm font-medium text-gray-900 dark:text-white" for="user_avatar">Upload file</label>
				<input ref="file" class="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400" name="image" aria-describedby="user_avatar_help" id="image" type="file">
				<div class="mt-1 text-sm text-gray-500 dark:text-gray-300" id="user_avatar_help">A profile picture is useful to confirm your are logged into your account</div>
			</div>

			<div class="flex justify-end">
				<button type="submit" class="bg-blue-700 rounded-md py-1.5 px-6 text-white ml-auto">Upload</button>
			</div>
		</form>

	</div>
</template>

<style scoped>

</style>