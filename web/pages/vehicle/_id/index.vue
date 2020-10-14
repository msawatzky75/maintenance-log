<template>
	<section class="section">
		<h1 class="is-1 title">
			{{ vehicle.year }} {{ vehicle.make }} {{ vehicle.model }}
		</h1>
		<h5 class="is-5 subtitle">
			<distance-display v-bind="vehicle.odometer" />
		</h5>

		<div class="columns">
			<div class="column">
				<h3 class="title is-3">
					Logs
					<small class="is-size-6">
						<nuxt-link :to="`${vehicle.id}/logs`">
							View all
						</nuxt-link>
					</small>
				</h3>

				<template v-if="vehicle && vehicle.logs">
					<b-table :data="vehicle.logs" :mobile-cards="false">
						<template #default="props">
							<b-table-column field="date" label="Date" sortable>
								{{ props.row.date }}
							</b-table-column>

							<b-table-column field="__typename" label="Type" sortable>
								<template v-if="props.row.__typename == 'FuelLog'">
									{{ props.row.__typename }}
								</template>
							</b-table-column>

							<b-table-column field="notes" label="Notes">
								{{ props.row.notes }}
							</b-table-column>
						</template>
					</b-table>
				</template>
				<template v-else>
					<p>No logs found.</p>
					<p>
						Try
						<nuxt-link :to="`${$route.params.id}/logs`">
							adding some.
						</nuxt-link>
					</p>
				</template>
			</div>
		</div>
	</section>
</template>

<script>
import VehicleQuery from '@/apollo/queries/vehicle.graphql'

export default {
	validate({ params, store }) {
		return (
			store.state.user.vehicles &&
			!!store.state.user.vehicles.find((v) => v.id === params.id)
		)
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
			vehicle: { logs: [] },
		}
	},
}
</script>
