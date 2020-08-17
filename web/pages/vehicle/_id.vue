<template>
	<section class="section">
		<h1 class="is-size-1">
			{{ vehicle.year }} {{ vehicle.make }} {{ vehicle.model }}
		</h1>
		<p class="is-size-5"><odometer v-bind="vehicle.odometer" /></p>
		{{ vehicle }}
	</section>
</template>

<script>
import VehicleQuery from '@/queries/vehicle.graphql'
import Odometer from '@/components/odometer'

export default {
	components: {
		Odometer,
	},
	async asyncData({ app, params }) {
		const apollo = app.apolloProvider.defaultClient

		return {
			vehicle: (
				await apollo.query({
					query: VehicleQuery,
					variables: {
						id: params.id,
						logFilter: { recent: 5 },
					},
				})
			).data.getVehicle,
		}
	},
	data() {
		return {
			vehicle: null,
		}
	},
}
</script>
