<template>
	<section>
		<div class="columns">
			<div class="column">
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

			<div class="column">
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

<script>
export default {
	name: 'MaintenenceLogInput',
	components: {},
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
				odometer: null,
				notes: null,
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
