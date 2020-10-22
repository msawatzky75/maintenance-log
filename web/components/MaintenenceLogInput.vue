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

		<InputWithValidation
			ref="notes"
			v-model.trim="innerValue.notes"
			label="Notes"
			input-type="textarea"
			:schema="yup.string().nullable().label('Notes').required()"
		/>
	</section>
</template>

<script>
import * as yup from 'yup'

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
			yup,
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

			return await Promise.all(promises).then(() => (callback instanceof Function ? callback : () => {}))
		},
	},
}
</script>
