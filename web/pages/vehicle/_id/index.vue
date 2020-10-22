<template>
	<section class="section">
		<h1 class="is-1 title">{{ vehicle.year }} {{ vehicle.make }} {{ vehicle.model }}</h1>
		<h5 class="is-5 subtitle">
			<DistanceDisplay v-bind="vehicle.odometer" />
		</h5>

		<div class="columns">
			<div class="column">
				<h3 class="title is-3">
					Logs
					<small class="is-size-6">
						<NuxtLink :to="`${vehicle.id}/logs`"> View all </NuxtLink>
					</small>
				</h3>

				<template v-if="vehicle && vehicle.logs">
					<BTable :data="vehicle.logs" :mobile-cards="false">
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
				</template>

				<template v-else>
					<p>No logs found.</p>
					<p>
						Try
						<NuxtLink :to="`${$route.params.id}/logs`"> adding some. </NuxtLink>
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
		return store.state.user.vehicles && !!store.state.user.vehicles.find((v) => v.id === params.id)
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
