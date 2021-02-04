<template>
	<section class="section">
		<h1 class="title">Add New Vehicle</h1>

		<pre v-if="errors">{{ errors }}</pre>

		<form @submit.prevent="createVehicle">
			<section class="section pt-0 px-0">
				<VehicleInput ref="vehicle" v-model="vehicle" />
			</section>

			<BField>
				<BButton type="is-primary" native-type="submit"> Create </BButton>
			</BField>
		</form>
	</section>
</template>

<script>
import * as yup from 'yup'
import debug from 'debug'
import VehicleInput from '../../components/VehicleInput.vue'
import createVehicle from '~/apollo/mutations/createVehicle.graphql'

const d = debug('ml.pages.vehicle.create')

export default {
	components: { VehicleInput },
	data() {
		return {
			yup,
			vehicle: {
				make: null,
				model: null,
				year: null,
				odometer: { value: null, type: null },
			},
			errors: null,
		}
	},
	methods: {
		async createVehicle() {
			try {
				await this.$refs.vehicle.validate()
				try {
					d.extend('createVehicle')('vaildated successfully')
					const response = await this.$apollo.mutate({
						mutation: createVehicle,
						variables: { data: this.vehicle },
					})

					d('redirect')
					this.$router.push(`/vehicle/${response.data.createVehicle.id}`)
				} catch (e) {
					this.errors = ['let me know if you ever see this message', 'also please share the error:', e.graphqlErrors]
					d('%o', e)
				}
			} catch (e) {
				d.extend('createVehicle')('validation error', e)
			}
		},
	},
}
</script>
