<template>
	<section>
		<div class="columns">
			<div class="column">
				<BField grouped>
					<InputWithValidation
						ref="year"
						v-model.number="innerValue.year"
						label="Year"
						placeholder="Year"
						:schema="yup.number().nullable().label('Year').required()"
					/>
					<InputWithValidation
						ref="make"
						v-model.trim="innerValue.make"
						label="Make"
						placeholder="Make"
						:schema="yup.string().nullable().label('Make').required()"
					/>
					<InputWithValidation
						ref="model"
						v-model.trim="innerValue.model"
						label="Model"
						placeholder="Model"
						:schema="yup.string().nullable().label('Model').required()"
					/>
				</BField>
			</div>
		</div>

		<div class="columns">
			<div class="column">
				<TypeAmountInput
					ref="odometer"
					v-model="innerValue.odometer"
					:types="$store.state.distanceTypes"
					label="Odometer"
					value-placeholder="Distance Value"
					type-placeholder="Distance Type"
				/>
			</div>
		</div>
	</section>
</template>

<script>
import * as yup from 'yup'

export default {
	name: 'VehicleInput',
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
				year: null,
				make: null,
				model: null,
				odometer: null,
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
