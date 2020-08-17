<template>
	<div>
		<section class="section">
			<h3 class="is-size-3">Logs</h3>
			<b-table :data="user.logs" :mobile-cards="false">
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
		</section>

		<section class="section">
			<h3 class="is-size-3">Vehicles</h3>
			<b-table :data="user.vehicles" :mobile-cards="false">
				<template slot-scope="props">
					<b-table-column field="id" label="Link" sortable>
						<nuxt-link :to="`/vehicle/${props.row.id}`">
							Link
						</nuxt-link>
					</b-table-column>

					<b-table-column field="year" label="Year" sortable>
						{{ props.row.year }}
					</b-table-column>

					<b-table-column field="make" label="Make" sortable>
						{{ props.row.make }}
					</b-table-column>

					<b-table-column field="model" label="Model" sortable>
						{{ props.row.model }}
					</b-table-column>

					<b-table-column field="odometer" label="Odometer" numeric>
						<odometer v-bind="props.row.odometer" />
					</b-table-column>
				</template>
			</b-table>
		</section>
	</div>
</template>

<script>
import UserQuery from '@/queries/user.graphql'
import Odometer from '@/components/odometer'

export default {
	components: {
		Odometer,
	},
	data() {
		return {
			user: {
				logs: [],
				vehicles: [],
			},
		}
	},
	async mounted() {
		this.user = (
			await this.$apollo.query({
				query: UserQuery,
				variables: {},
			})
		).data.getUser
	},
}
</script>

<style></style>
