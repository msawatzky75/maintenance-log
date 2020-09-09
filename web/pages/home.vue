<template>
	<div>
		<section class="section">
			<h1 class="title">You</h1>
			<div class="columns">
				<div class="column is-narrow has-text-right">
					<p>Email:</p>
					<p>Vehicles:</p>
					<p>Preffered Distance Unit:</p>
					<p>Preffered Fluid Unit:</p>
					<p>Preffered Money Unit:</p>
				</div>

				<div class="column">
					<p>{{ user.email }}</p>
					<p>{{ user.vehicles.length }}</p>
					<p>
						{{
							user.preference.distance
								? user.preference.distance
								: 'None selected'
						}}
						<nuxt-link to="/preference">change</nuxt-link>
					</p>
					<p>
						{{
							user.preference.fluid ? user.preference.fluid : 'None selected'
						}}
						<nuxt-link to="/preference">change</nuxt-link>
					</p>
					<p>
						{{
							user.preference.money ? user.preference.money : 'None selected'
						}}
						<nuxt-link to="/preference">change</nuxt-link>
					</p>
				</div>
			</div>
		</section>

		<section class="section">
			<h4 class="title is-4">Vehicles</h4>
			<h6 class="subtitle is-6">
				<nuxt-link to="/vehicle/create">Add new Vehicle</nuxt-link>
			</h6>

			<b-table :data="user.vehicles" hoverable @click="vehicleClick">
				<template slot-scope="props">
					<b-table-column field="year" label="Year" width="40" numeric sortable>
						{{ props.row.year }}
					</b-table-column>

					<b-table-column field="make" label="Make" sortable>
						{{ props.row.make }}
					</b-table-column>

					<b-table-column field="model" label="Model" sortable>
						{{ props.row.model }}
					</b-table-column>

					<b-table-column field="odometer.value" label="Odometer" sortable>
						<distance-display v-bind="props.row.odometer" />
					</b-table-column>
				</template>
			</b-table>
		</section>
	</div>
</template>

<script>
export default {
	validate({ store }) {
		return !!store.state.user.id
	},
	computed: {
		user() {
			return this.$store.state.user
		},
	},
	methods: {
		vehicleClick({ id }) {
			this.$router.push('/vehicle/' + id)
		},
	},
}
</script>
