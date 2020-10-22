<template>
	<section class="section">
		<h3 class="title is-3">Logs</h3>

		<section>
			<BButton tag="nuxt-link" :to="{ name: 'vehicle-id-logs-new', query: { type: 'MaintenenceLog' } }">
				Create New Log
			</BButton>

			<BButton tag="nuxt-link" :to="{ name: 'vehicle-id-logs-new', query: { type: 'FuelLog' } }">
				Create New Fuel Log
			</BButton>

			<BButton tag="nuxt-link" :to="{ name: 'vehicle-id-logs-new', query: { type: 'MaintenenceLog' } }">
				Create New Maintenence Log
			</BButton>

			<BButton tag="nuxt-link" :to="{ name: 'vehicle-id-logs-new', query: { type: 'OilChangeLog' } }">
				Create New Oil Change Log
			</BButton>
		</section>

		<hr />

		<template v-if="vehicle">
			<BTable
				:data="vehicle.logs || []"
				:mobile-cards="false"
				default-sort="date"
				default-sort-direction="desc"
				sort-icon="chevron-up"
			>
				<template #default="props">
					<BTableColumn field="date" label="Date" sortable>
						{{ props.row.date }}
					</BTableColumn>

					<BTableColumn field="__typename" label="Type" sortable>
						{{ props.row.__typename }}
					</BTableColumn>

					<BTableColumn field="odometer.value" label="Odometer" sortable>
						<DistanceDisplay v-bind="props.row.odometer" />
					</BTableColumn>

					<BTableColumn field="notes" label="Notes" sortable>
						{{ props.row.notes }}
					</BTableColumn>
				</template>
			</BTable>

			<section v-if="!vehicle.logs">
				<p>No Logs :(</p>
			</section>
		</template>

		<template v-else> loading... </template>
	</section>
</template>

<script>
import debug from 'debug'
import VehicleQuery from '@/apollo/queries/vehicle.graphql'

const d = debug('ml.pages.vehicle._id.logs.index')
export default {
	name: 'Logs',
	apollo: {
		getVehicle: {
			query: VehicleQuery,
			variables() {
				return {
					id: this.$route.params.id,
					logFilter: { recent: 25 }, // TODO pagination here
				}
			},
			result({ data, loading }) {
				if (!loading) {
					this.vehicle = data.getVehicle
				}
			},
		},
	},
	data() {
		return {
			vehicle: null,
		}
	},
	async mounted() {
		await this.$apollo.queries.getVehicle.refetch()
	},
}
</script>
