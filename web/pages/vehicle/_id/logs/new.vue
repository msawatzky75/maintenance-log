<template>
	<section class="section">
		<h1 class="title">Create A New Log</h1>
		<section>
			<BField label="Log Type">
				<BSelect v-model="selectedLogType" required placeholder="Select a log type">
					<option v-for="t in $store.state.logTypes" :key="t.value" :value="t">
						{{ t.name }}
					</option>
				</BSelect>
			</BField>
		</section>

		<form @submit.prevent="submitLog">
			<section v-if="selectedLogType" class="mt-4">
				<template v-if="selectedLogType.value === LogTypes.FuelLog">
					<FuelLogInput :ref="LogTypes.FuelLog" v-model="log" />
				</template>

				<template v-else-if="selectedLogType.value === LogTypes.MaintenanceLog">
					<MaintenenceLogInput :ref="LogTypes.MaintenanceLog" v-model="log" />
				</template>

				<template v-else-if="selectedLogType.value === LogTypes.OilChangeLog">
					<OilChangeLogInput :ref="LogTypes.OilChangeLog" v-model="log" />
				</template>

				<br />

				<BButton type="is-primary" native-type="submit">Create</BButton>
				<BButton type="is-secondary" native-type="reset">Cancel</BButton>
				<pre v-if="error">{{ error }}</pre>
			</section>
		</form>
	</section>
</template>

<script>
import debug from 'debug'

const d = debug('ml.pages.vheicle._id.logs.new')

const LogTypes = {
	FuelLog: 'FuelLog',
	MaintenanceLog: 'MaintenanceLog',
	OilChangeLog: 'OilChangeLog',
}

export default {
	data() {
		return {
			LogTypes,
			selectedLogType: this.$route.query.type ? this.$store.getters.logType(this.$route.query.type) : null,
			log: null,
			error: null,
		}
	},
	watch: {
		// Clear unshared data when log type is changed.
		selectedLogType(newType) {
			if (this.log) this.log = { date: this.log.date, notes: this.log.notes }
			this.$router.replace({
				name: 'vehicle-id-logs-new',
				query: { type: newType.type },
			})
		},
	},
	methods: {
		async submitLog() {
			// d({ ...this.log, vehicleId: this.$route.params.id })
			try {
				await this.$refs[this.selectedLogType.value].validate()
				try {
					await this.$apollo.mutate({
						mutation: this.selectedLogType.createMutation,
						variables: {
							data: { ...this.log, vehicleId: this.$route.params.id },
						},
					})
					this.$router.push({ name: 'vehicle-id-logs' })
				} catch (e) {
					d('Submition Error', e.graphqlErrors, e.message)
					this.error = e.networkError.result.errors
				}
			} catch (e) {
				d('Validation Error', e)
			}
		},
	},
}
</script>

<style></style>
