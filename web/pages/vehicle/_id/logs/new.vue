<template>
	<section class="section">
		<h1 class="title">Create A New Log</h1>
		<section>
			<BField label="Log Type">
				<BSelect v-model="selectedLogType" required placeholder="Select a log type">
					<option v-for="t in LogTypes" :key="t.type" :value="t">
						{{ t.name }}
					</option>
				</BSelect>
			</BField>
		</section>

		<form @submit.prevent="submitLog">
			<section v-if="selectedLogType" class="mt-4">
				<template v-if="selectedLogType === LogTypes.FuelLog">
					<FuelLogInput :ref="LogTypes.FuelLog" v-model="log" />
				</template>

				<template v-else-if="selectedLogType === LogTypes.MaintenenceLog">
					<MaintenenceLogInput :ref="LogTypes.MaintenenceLog" v-model="log" />
				</template>

				<template v-else-if="selectedLogType === LogTypes.OilChangeLog">
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
import CreateFuelLog from '~/apollo/mutations/create-fuel-log.graphql'
import CreateMaintenenceLog from '~/apollo/mutations/create-maintenence-log.graphql'
import CreateOilChangeLog from '~/apollo/mutations/create-oil-change-log.graphql'

const d = debug('ml.pages.vheicle._id.logs.new')

const LogTypes = {
	FuelLog: {
		type: 'FuelLog',
		name: 'Fuel Log',
		_mutation: CreateFuelLog,
	},
	MaintenenceLog: {
		type: 'MaintenenceLog',
		name: 'Maintenence Log',
		_mutation: CreateMaintenenceLog,
	},
	OilChangeLog: {
		type: 'OilChangeLog',
		name: 'Oil Change Log',
		_mutation: CreateOilChangeLog,
	},
}

export default {
	data() {
		return {
			LogTypes,
			selectedLogType: this.$route.query.type ? LogTypes[this.$route.query.type] : null,
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
				await this.$refs[this.selectedLogType].validate()
				try {
					await this.$apollo.mutate({
						mutation: this.selectedLogType._mutation,
						variables: {
							data: { ...this.log, vehicleId: this.$route.params.id },
						},
					})
					this.$router.push({ name: 'vehicle-id-logs' })
				} catch (e) {
					d('Submition Error', Object.keys(e), e.graphqlErrors, e.message)
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
