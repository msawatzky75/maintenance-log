<template>
	<section class="section">
		<h1 class="is-1 title">
			{{ vehicle.year }} {{ vehicle.make }} {{ vehicle.model }}
		</h1>
		<h5 class="is-5 subtitle"><odometer v-bind="vehicle.odometer" /></h5>

		<div class="columns">
			<div class="column">
				<h3 class="is-3 title">
					Logs
					<small class="is-size-6">
						<nuxt-link :to="`${vehicle.id}/logs`">
							View all
						</nuxt-link>
					</small>
				</h3>
				<b-table :data="vehicle.logs" :mobile-cards="false">
					<template slot-scope="props">
						<b-table-column field="date" label="Date" sortable>
							{{ props.row.date }}
						</b-table-column>

						<b-table-column field="__typename" label="Type" sortable>
							{{ props.row.__typename }}
						</b-table-column>

						<b-table-column field="notes" label="Notes">
							{{ props.row.notes }}
						</b-table-column>
					</template>
				</b-table>
			</div>
		</div>
	</section>
</template>

<script>
import VehicleQuery from '@/apollo/queries/vehicle.graphql'
import Odometer from '@/components/odometer'

export default {
	validate({ params, store }) {
		return (
			store.state.user.vehicles &&
			!!store.state.user.vehicles.find((v) => v.id === params.id)
		)
	},
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
						logFilter: { recent: 10 },
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
