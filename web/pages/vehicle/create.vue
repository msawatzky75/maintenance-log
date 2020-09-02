<template>
	<section class="section">
		<h1 class="title">Add New Vehicle</h1>

		<pre v-if="errors">{{ errors }}</pre>

		<validation-observer v-slot="{ handleSubmit, reset }">
			<form
				@submit.prevent="handleSubmit(createVehicle)"
				@reset.prevent="reset"
			>
				<b-field grouped>
					<b-input-with-validation
						v-model="vehicle.make"
						type="text"
						placeholder="Make"
						label="Make"
						rules="required"
					/>
					<b-input-with-validation
						v-model="vehicle.model"
						type="text"
						placeholder="Model"
						label="Model"
						rules="required"
					/>
					<b-input-with-validation
						v-model="vehicle.year"
						type="text"
						placeholder="Year"
						label="Year"
						rules="required|numeric"
					/>
				</b-field>

				<OdometerInput v-model="vehicle.odometer" label="Odometer" />

				<b-field>
					<b-button type="is-primary" native-type="submit">
						Create
					</b-button>
				</b-field>
			</form>
		</validation-observer>
	</section>
</template>

<script>
import debug from 'debug'
import { ValidationProvider, ValidationObserver, extend } from 'vee-validate'
import createVehicle from '~/apollo/mutations/createVehicle.graphql'
import OdometerInput from '~/components/odometer-input'

const d = debug('ml.pages.vehicle.create')

export default {
	components: {
		OdometerInput,
		ValidationObserver,
	},
	data() {
		return {
			vehicle: {
				make: '',
				model: '',
				year: null,
				odometer: { value: null, type: null },
			},
			errors: null,
		}
	},
	methods: {
		validVehicle(v) {
			d.extend('validateVehicle')('validate: ', v)
			return (
				v.make.length &&
				v.model.length &&
				v.year &&
				v.odometer.value &&
				v.odometer.type
			)
		},
		async createVehicle() {
			try {
				if (!this.validVehicle(this.vehicle)) {
					throw new Error('Invalid vehicle')
				}

				const response = await this.$apollo.mutate({
					mutation: createVehicle,
					variables: { data: this.vehicle },
				})

				d('redirect')
				this.$router.push(`/vehicle/${response.data.createVehicle.id}`)
			} catch (e) {
				this.errors = [
					'let me know if you ever see this message',
					'also please share the error:',
					e.graphqlErrors,
				]
				d('%o', e)
			}
		},
	},
}
</script>
