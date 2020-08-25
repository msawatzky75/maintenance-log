<template>
	<div>
		<section class="section">
			<template v-if="loading">{{ this.$store.state.user }}</template>
			<template v-else
				><h1 class="is-size-1">:(</h1>
				No data, try <nuxt-link to="/login">logging in</nuxt-link></template
			>
		</section>
	</div>
</template>

<script>
import Odometer from '@/components/odometer'

export default {
	computed: {
		loading() {
			return !!this.$store.state.user.id
		},
	},
	middleware: [
		({ store, redirect }) => {
			if (!store.getters.hasVehicleSelected) {
				redirect('/vehicle')
			}
		},
	],
}
</script>
