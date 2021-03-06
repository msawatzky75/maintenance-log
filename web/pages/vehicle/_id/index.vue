<template>
	<section class="section">
		<h1 class="is-1 title">{{ vehicle.year }} {{ vehicle.make }} {{ vehicle.model }}</h1>
		<h5 class="is-5 subtitle">
			<DistanceDisplay v-bind="vehicle.odometer" />
		</h5>

		<!-- <section>
			<h3 class="title is-3">
				Insights
				<small class="is-size-6">
					<NuxtLink :to="`${vehicle.id}/logs`">View all</NuxtLink>
				</small>
			</h3>

			Insights here. <LineChart :chartdata="[]" />
		</section>

		<br /> -->

		<section>
			<h3 class="title is-3">
				Logs
				<small class="is-size-6">
					<NuxtLink :to="`${vehicle.id}/logs`">View all</NuxtLink>
				</small>
			</h3>

			<template v-if="vehicle && vehicle.logs">
				<BTable :data="vehicle.logs" :mobile-cards="false">
					<BTableColumn v-slot="props" field="date" label="Date" sortable>
						{{ formatDate(props.row.date) }}
					</BTableColumn>

					<BTableColumn v-slot="props" field="__typename" label="Type" sortable>
						{{ $store.getters.logType(props.row.__typename).name }}
					</BTableColumn>

					<BTableColumn v-slot="props" field="odometer.value" label="Odometer" sortable>
						<DistanceDisplay v-bind="props.row.odometer" />
					</BTableColumn>

					<BTableColumn v-slot="props" field="notes" label="Notes" sortable>
						{{ props.row.notes }}
					</BTableColumn>
				</BTable>
			</template>

			<template v-else>
				<p>No logs found.</p>
				<p>
					Try
					<NuxtLink :to="`${$route.params.id}/logs`"> adding some. </NuxtLink>
				</p>
			</template>
		</section>
	</section>
</template>

<script lang="ts">
import Vue from 'vue'
import moment from 'moment'
import VehicleQuery from '~/apollo/queries/vehicle.graphql'
import type { User } from '~/store/user'

export default Vue.extend({
	validate({ params, store }) {
		return !!(store.state.user as User).vehicles?.find((v) => v.id === params.id)
	},
	async asyncData({ app, params }) {
		const apollo = app.apolloProvider?.defaultClient

		if (!apollo) {
			throw new Error('Apollo Error: Unable to find an apollo provider.')
		}
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
	methods: {
		formatDate(d: Date | string): string {
			return moment(d).format('YYYY-MM-DD HH:mm:ss')
		},
	},
})
</script>
