<template>
	<section>
		<div class="columns is-multiline">
			<div class="column is-4">
				<DatepickerWithValidation
					ref="date"
					v-model="innerValue.date"
					label="Date"
					placeholder="Click to select..."
					icon="calendar-today"
					trap-focus
					required
				/>
			</div>

			<div class="column is-4">
				<TypeAmountInput
					ref="fuelAmount"
					v-model="innerValue.fuelAmount"
					:types="$store.state.fluidTypes"
					label="Fuel Amount"
					amount-placeholder="Fluid Value"
					type-placeholder="Fluid Type"
				/>
			</div>

			<div class="column is-4">
				<TypeAmountInput
					ref="fuelPrice"
					v-model="innerValue.fuelPrice"
					:types="$store.state.currencyTypes"
					label="Fuel Price"
					amount-placeholder="Currency Value"
					type-placeholder="Currency Type"
				/>
			</div>

			<div class="column is-4">
				<InputWithValidation
					ref="fuelGrade"
					v-model.number="innerValue.grade"
					label="Fuel Grade"
					type="number"
					:schema="
						yup
							.number()
							.typeError('Fuel Grade must be a number')
							.nullable()
							.label('Fuel Grade')
							.required()
					"
				/>
			</div>

			<div class="column is-4">
				<TypeAmountInput
					ref="trip"
					v-model="innerValue.trip"
					:types="$store.state.distanceTypes"
					label="Trip"
					amount-placeholder="Distance Value"
					type-placeholder="Distance Type"
				/>
			</div>

			<div class="column is-4">
				<TypeAmountInput
					ref="odometer"
					v-model="innerValue.odometer"
					:types="$store.state.distanceTypes"
					label="Odometer"
					amount-placeholder="Distance Value"
					type-placeholder="Distance Type"
				/>
			</div>
		</div>

		<BField label="Notes">
			<BInput v-model="innerValue.notes" type="textarea" />
		</BField>
	</section>
</template>

<script lang="ts">
import * as yup from 'yup'
import debug from 'debug'

const d = debug('ml.components.FuelLogInput')

export default {
	name: 'FuelLogInput',
	props: {
		value: {
			type: Object,
			default: null,
		},
	},
	emits: ['validate', 'input'],
	data() {
		return {
			yup,
			innerValue: {
				date: null,
				notes: null,
				trip: null,
				odometer: null,
				grade: null,
				fuelAmount: null,
				fuelPrice: null,
			},
		}
	},
	watch: {
		innerValue: {
			handler(newVal) {
				this.$emit('input', newVal)
			},
			deep: true,
		},
		value(newVal) {
			this.innerValue = newVal
		},
	},
	created() {
		if (this.value) {
			this.innerValue = this.value || {}
		}
	},
	methods: {
		async validate(callback) {
			const promises = Object.keys(this.$refs).map(async (r) => {
				if (this.$refs[r].validate instanceof Function) {
					await this.$refs[r].validate()
				}
			})

			return await Promise.all(promises).then(() =>
				callback instanceof Function ? callback : () => {}
			)
		},
	},
}
</script>
