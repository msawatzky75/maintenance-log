<template>
	<div>
		<section class="section">
			<h1 class="title">Your Profile</h1>
			<h4 class="subtitle">
				{{ user.email }} ({{ user.vehicles.length }} vehicle{{ user.vehicles.length != 1 ? 's' : '' }})
			</h4>

			<section>
				<div class="columns">
					<div class="column is-narrow">
						<BField label="Distance Type">
							<BSelect v-model="preference.distanceType" type="text">
								<option :value="null" hidden>Select a Distance type</option>
								<option v-for="(t, k) in $store.state.distanceTypes" :key="k" :value="t.value">
									{{ t.name }}
								</option>
							</BSelect>
						</BField>
					</div>

					<div class="column is-narrow">
						<BField label="Fluid Type">
							<BSelect v-model="preference.fluidType" type="text">
								<option :value="null" hidden>Select a Fluid type</option>
								<option v-for="(t, k) in $store.state.fluidTypes" :key="k" :value="t.value">
									{{ t.name }}
								</option>
							</BSelect>
						</BField>
					</div>

					<div class="column is-narrow">
						<BField label="Currency Type">
							<BSelect v-model="preference.currencyType" type="text">
								<option :value="null" hidden>Select a Currency type</option>
								<option v-for="(t, k) in $store.state.currencyTypes" :key="k" :value="t.value">
									{{ t.name }}
								</option>
							</BSelect>
						</BField>
					</div>

					<div class="column is-narrow">
						<BField label="Default Vehicle">
							<BSelect v-model="preference.vehicleId" type="text">
								<option :value="null" hidden>Select a Vehicle</option>
								<option v-for="(v, k) in $store.state.user.vehicles" :key="k" :value="v.id">
									{{ v.year }} {{ v.make }} {{ v.model }} &mdash; <DistanceDisplay v-bind="v.odometer" short />
								</option>
							</BSelect>
						</BField>
					</div>
				</div>

				<div class="columns">
					<div class="column is-full-height">
						<BButton type="is-primary" @click="updatePreference">Update preference</BButton>
					</div>
				</div>

				<pre v-if="error">{{ error }}</pre>
			</section>
		</section>

		<section class="section">
			<h4 class="title is-4">Vehicles</h4>
			<h6 class="subtitle is-6">
				<NuxtLink to="/vehicle/create">Add new Vehicle</NuxtLink>
			</h6>

			<div class="columns is-multiline">
				<div v-for="(v, k) in user.vehicles" :key="k" class="column is-flex">
					<NuxtLink
						class="box is-align-self-center is-full-height is-full-width"
						:to="{ name: 'vehicle-id', params: { id: v.id } }"
					>
						<h4 class="title">{{ v.year }} {{ v.make }} {{ v.model }}</h4>
						<DistanceDisplay class="subtitle is-6" v-bind="v.odometer" />

						<BTable
							v-if="v.logs && v.logs.length"
							class="is-hidden-mobile"
							:data="v.logs"
							default-sort="date"
							default-sort-direction="desc"
						>
							<template #default="props">
								<BTableColumn field="date" label="Date" width="40" sortable>
									{{ new Date(props.row.date).toISOString().split('T')[0] }}
								</BTableColumn>

								<BTableColumn field="odometer.value" label="Odometer" class="has-text-right" sortable>
									<DistanceDisplay v-bind="props.row.odometer" short />
								</BTableColumn>

								<BTableColumn field="notes" label="Notes" sortable>
									{{ props.row.notes }}
								</BTableColumn>

								<BTableColumn field="__typename" label="Log Type" sortable>
									{{ $store.state.logTypes[props.row.__typename] }}
								</BTableColumn>
							</template>
						</BTable>
						<div v-else>No logs yet.</div>
					</NuxtLink>
				</div>
			</div>
		</section>
	</div>
</template>

<script lang="ts">
import Vue from 'vue'
import debug from 'debug'
import updatePreference from '~/apollo/mutations/updatePreference.graphql'

const d = debug('ml.web.pages.profile')

export default Vue.extend({
	name: 'Profile',
	validate({ store }) {
		return !!store.state.user.id
	},
	data() {
		return {
			error: null,
			preference: {
				distanceType: this.$store.state.user.preference.distance,
				fluidType: this.$store.state.user.preference.fluid,
				currencyType: this.$store.state.user.preference.money,
				vehicleId: this.$store.state.user.preference.vehicle?.id,
			},
		}
	},
	computed: {
		user() {
			return this.$store.state.user
		},
	},
	methods: {
		async updatePreference() {
			try {
				await this.$apollo.mutate({
					mutation: updatePreference,
					variables: {
						data: {
							distance: this.preference.distanceType,
							fluid: this.preference.fluidType,
							money: this.preference.currencyType,
							vehicleId: this.preference.vehicleId,
						},
					},
				})
			} catch (e) {
				d('Submition Error', Object.keys(e), e.graphqlErrors, e.message)
				if (e.networkError) {
					this.error = e.networkError.result.errors
				} else {
					this.error = e.message
				}
			}
		},
	},
})
</script>

<style lang="scss" scoped>
.is-full-height {
	height: 100%;
}
.is-full-width {
	width: 100%;
}
</style>
