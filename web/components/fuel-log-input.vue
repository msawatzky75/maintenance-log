<template>
	<section>
		<div class="columns is-multiline">
			<div class="column is-4">
				<b-field label="Date">
					<b-datepicker
						v-model="innerValue.date"
						placeholder="Click to select..."
						icon="calendar-today"
						trap-focus
						required
					>
					</b-datepicker>
				</b-field>
			</div>

			<div class="column is-4">
				<type-amount-input
					v-model="innerValue.fuelAmount"
					:types="$store.state.fluidTypes"
					label="Fuel Amount"
					amount-placeholder="Fluid Value"
					type-placeholder="Fluid Type"
					positive
					required
				/>
			</div>

			<div class="column is-4">
				<type-amount-input
					v-model="innerValue.fuelPrice"
					:types="$store.state.currencyTypes"
					label="Fuel Price"
					amount-placeholder="Currency Value"
					type-placeholder="Currency Type"
					positive
					required
				/>
			</div>

			<div class="column is-4">
				<b-input-with-validation
					v-model="innerValue.grade"
					rules="required|numeric"
					label="Fuel Grade"
				/>
			</div>

			<div class="column is-4">
				<type-amount-input
					v-model="innerValue.trip"
					:types="$store.state.distanceTypes"
					label="Trip"
					amount-placeholder="Distance Value"
					type-placeholder="Distance Type"
					positive
					required
				/>
			</div>

			<div class="column is-4">
				<type-amount-input
					v-model="innerValue.Odometer"
					:types="$store.state.distanceTypes"
					label="Odometer"
					amount-placeholder="Distance Value"
					type-placeholder="Distance Type"
					positive
					required
				/>
			</div>
		</div>

		<b-field label="Notes">
			<b-input v-model="innerValue.notes" type="textarea" />
		</b-field>
	</section>
</template>

<script>
export default {
	name: 'FuelLogInput',
	props: {
		value: {
			type: Object,
			default: null,
		},
	},
	data() {
		return {
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
}
</script>
