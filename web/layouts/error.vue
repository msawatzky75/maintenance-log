<template>
	<section class="section">
		<h1 class="title">Error Code: {{ error.statusCode }}</h1>
		<div class="content">
			<pre>{{ error.message }}</pre>
		</div>
	</section>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
	name: 'Error',
	props: {
		error: {
			type: Object as () => { statusCode: number; message: string },
			required: true,
		},
	},
})
</script>
