<template>
	<section class="section">
		<h1 class="title">Create A New Log</h1>
		<section>
			<b-field label="Log Type">
				<b-select
					v-model="selectedLogType"
					required
					placeholder="Select a log type"
				>
					<option v-for="t in LogTypes" :key="t.type" :value="t">
						{{ t.name }}
					</option>
				</b-select>
			</b-field>
		</section>

		<validation-observer v-slot="{ handleSubmit, reset }">
			<form @submit.prevent="handleSubmit(submitLog)" @reset.prevent="reset">
				<section v-if="selectedLogType" class="mt-4">
					<template v-if="selectedLogType === LogTypes.FuelLog">
						<fuel-log-input v-model="log" />
					</template>

					<template v-else-if="selectedLogType === LogTypes.MaintenenceLog">
						<maintenence-log-input v-model="log" />
					</template>

					<template v-else-if="selectedLogType === LogTypes.OilChangeLog">
						<oil-change-log-input v-model="log" />
					</template>

					<br />

					<b-button type="is-primary" native-type="submit">Create</b-button>
					<b-button type="is-secondary" native-type="reset">Cancel</b-button>
					<pre v-if="log">{{ log }}</pre>
				</section>
			</form>
		</validation-observer>
	</section>
</template>

<script>
import { ValidationObserver } from 'vee-validate'
import CreateFuelLog from '~/apollo/mutations/create-fuel-log.graphql'
import CreateMaintenenceLog from '~/apollo/mutations/create-maintenence-log.graphql'
import CreateOilChangeLog from '~/apollo/mutations/create-oil-change-log.graphql'

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
	components: { ValidationObserver },
	data() {
		return {
			LogTypes,
			selectedLogType: LogTypes.FuelLog,
			log: null,
		}
	},
	watch: {
		// Clear unshared data when log type is changed.
		selectedLogType() {
			if (this.log) this.log = { date: this.log.date, notes: this.log.notes }
		},
	},
	methods: {
		async submitLog() {
			console.dir({ ...this.log, vehicleId: this.$route.params.id })
			// await this.$apollo.mutate({
			// 	mutation: this.selectedLogType._mutation,
			// 	variables: { data: this.log },
			// })
		},
	},
}
</script>

<style></style>
