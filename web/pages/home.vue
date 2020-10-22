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
						{{ user.preference.distance ? user.preference.distance : 'None selected' }}
						<NuxtLink to="/preference">change</NuxtLink>
					</p>
					<p>
						{{ user.preference.fluid ? user.preference.fluid : 'None selected' }}
						<NuxtLink to="/preference">change</NuxtLink>
					</p>
					<p>
						{{ user.preference.money ? user.preference.money : 'None selected' }}
						<NuxtLink to="/preference">change</NuxtLink>
					</p>
				</div>
			</div>
		</section>

		<section class="section">
			<h4 class="title is-4">Vehicles</h4>
			<h6 class="subtitle is-6">
				<NuxtLink to="/vehicle/create">Add new Vehicle</NuxtLink>
			</h6>

			<BTable :data="user.vehicles" hoverable @click="vehicleClick">
				<template #default="props">
					<BTableColumn field="year" label="Year" width="40" numeric sortable>
						{{ props.row.year }}
					</BTableColumn>

					<BTableColumn field="make" label="Make" sortable>
						{{ props.row.make }}
					</BTableColumn>

					<BTableColumn field="model" label="Model" sortable>
						{{ props.row.model }}
					</BTableColumn>

					<BTableColumn field="odometer.value" label="Odometer" sortable>
						<DistanceDisplay v-bind="props.row.odometer" />
					</BTableColumn>
				</template>
			</BTable>
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
