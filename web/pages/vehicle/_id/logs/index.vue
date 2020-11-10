<template>
	<section class="section">
		<h3 class="title is-3">Logs for {{ vehicle.year }} {{ vehicle.make }} {{ vehicle.model }}</h3>

		<section>
			<BButton tag="nuxt-link" :to="{ name: 'vehicle-id-logs-new', query: { type: 'MaintenanceLog' } }">
				Create New Log
			</BButton>

			<BButton tag="nuxt-link" :to="{ name: 'vehicle-id-logs-new', query: { type: 'FuelLog' } }">
				Create New Fuel Log
			</BButton>

			<BButton tag="nuxt-link" :to="{ name: 'vehicle-id-logs-new', query: { type: 'MaintenanceLog' } }">
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
				<BTableColumn v-slot="props" field="date" label="Date" sortable>
					{{ formatDate(props.row.date) }}
				</BTableColumn>

				<BTableColumn v-slot="props" field="__typename" label="Type" sortable>
					{{ $store.state.logTypes[props.row.__typename] || props.row.__typename }}
				</BTableColumn>

				<BTableColumn v-slot="props" field="odometer.value" label="Odometer" sortable>
					<DistanceDisplay v-bind="props.row.odometer" />
				</BTableColumn>

				<BTableColumn v-slot="props" field="notes" label="Notes" sortable>
					{{ props.row.notes }}
				</BTableColumn>

				<BTableColumn v-slot="props" field="id" class="action-column">
					<div class="is-pulled-right">
						<!-- <BButton type="is-info" icon-right="pencil" /> -->
						<BButton type="is-danger" icon-right="delete" @click="deleteLog(props.row)" />
					</div>
				</BTableColumn>
			</BTable>

			<section v-if="!vehicle.logs">
				<p>No Logs :(</p>
			</section>
		</template>

		<template v-else> loading... </template>
	</section>
</template>

<script lang="ts">
import Vue from 'vue'
import debug from 'debug'
import moment from 'moment'
import VehicleQuery from '~/apollo/queries/vehicle.graphql'
import type { Log } from '~/store/user'

const d = debug('ml.pages.vehicle._id.logs.index')
export default Vue.extend({
	name: 'Logs',
	async asyncData({ app, params }) {
		const apollo = app.apolloProvider.defaultClient

		return {
			vehicle: (
				await apollo.query({
					query: VehicleQuery,
					variables: {
						id: params.id,
						logFilter: { recent: 25 },
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
	methods: {
		formatDate(d: moment.MomentInput): string {
			return moment(d).format('YYYY-MM-DD HH:mm:ss')
		},
		deleteLog(log: Log): void {
			this.$buefy.dialog.confirm({
				title: 'Deleting ' + this.$store.getters.logType(log.__typename).name,
				message:
					"Are you sure you want to delete this log? This action cannot be undone. This will not update your vehicle's odometer.<br/><br/>" +
					`Date: ${this.formatDate(log.date)}<br/>` +
					`Note: ${log.notes}<br/>`,
				confirmText: 'Delete Log',
				type: 'is-danger',
				hasIcon: true,
				onConfirm: async () => {
					try {
						await this.$apollo.mutate({
							mutation: this.$store.getters.logType(log.__typename).deleteMutation,
							variables: {
								id: log.id,
								type: log.__typename,
							},
						})
						this.$buefy.toast.open('Log deleted.')
						// TODO refetch
					} catch (e) {
						d('log deletion error', e.graphqlErrors, e.message)
						this.$buefy.toast.open('Unable to delete, try again later.')
					}
				},
			})
		},
	},
})
</script>

<style lang="scss">
.action-column {
	white-space: nowrap;
	width: auto;
}
</style>
