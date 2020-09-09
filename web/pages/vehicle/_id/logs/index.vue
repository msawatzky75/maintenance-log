<template>
	<section class="section">
		<h3 class="title is-3">Logs</h3>
		<b-table :data="vehicle.logs || []" :mobile-cards="false">
			<template slot-scope="props">
				<b-table-column field="date" label="Date" sortable>
					{{ props.row.date }}
				</b-table-column>

				<b-table-column field="__typename" label="Type" sortable>
					{{ props.row.__typename }}
				</b-table-column>

				<b-table-column field="vehicle" label="Vehicle">
					<nuxt-link :to="`/vehicle/${props.row.vehicle.id}`">
						{{ props.row.vehicle.year }}
						{{ props.row.vehicle.make }}
						{{ props.row.vehicle.model }}
					</nuxt-link>
				</b-table-column>

				<b-table-column field="notes" label="Notes" sortable>
					{{ props.row.date }}
				</b-table-column>
			</template>
		</b-table>

		<section v-if="!vehicle.logs">
			<p>No Logs :(</p>
		</section>

		<hr />

		<section>
			<b-button tag="nuxt-link" to="logs/new"> Create New Log</b-button>

			<b-button tag="nuxt-link" to="logs/new/fuel">
				Create New Fuel Log
			</b-button>

			<b-button tag="nuxt-link" to="logs/new/maintenence">
				Create New Maintenence Log
			</b-button>

			<b-button tag="nuxt-link" to="logs/new/oil">
				Create New Oil-Change Log
			</b-button>
		</section>
	</section>
</template>

<script>
import VehicleQuery from '@/apollo/queries/vehicle.graphql'
export default {
	name: 'Logs',
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
			logs: [],
		}
	},
}
</script>
