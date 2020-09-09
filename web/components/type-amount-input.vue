<template>
	<b-field :label="label" :type="status" :message="message">
		<b-field :type="status">
			<b-select
				v-model="innerValue.type"
				type="text"
				:placeholder="typePlaceholder"
				@input="validate"
			>
				<option v-for="(t, k) in types" :key="k" :value="t.value">
					{{ t.name }}
				</option>
			</b-select>
			<b-input
				v-model="innerValue.amount"
				type="number"
				:placeholder="amountPlaceholder"
				@input="validate"
			/>
		</b-field>
	</b-field>
</template>

<script>
import debug from 'debug'
import { extend, validate } from 'vee-validate'

const d = debug('ml.components.type-amount-input')

extend('type-required', {
	params: ['type'],
	validate(_, { type }) {
		if (type && type.length) {
			return true
		}
		return 'Type is required.'
	},
})
extend('positive', {
	validate(v) {
		return v >= 0 ? true : '{_field_} must be positive'
	},
})

export default {
	name: 'TypeAmountInput',
	props: {
		value: { type: Object, default: null },
		label: { type: String, default: null },
		required: { type: Boolean, default: false },
		positive: { type: Boolean, default: false },
		types: { type: Array, required: true },
		amountPlaceholder: { type: String, default: null },
		typePlaceholder: { type: String, default: null },
	},
	data() {
		return {
			status: null,
			message: null,
			innerValue: {
				amount: null,
				type: null,
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
		async validate() {
			const rules = {
				type: [this.required ? 'required' : null].join('|'),
				amount: [
					this.required ? 'required' : null,
					'type-required:@type',
					this.positive ? 'positive' : null,
				]
					.filter((v) => v)
					.join('|'),
			}
			const typeValidation = await validate(this.innerValue.type, rules.type, {
				name: (this.label + ' Type').trim(),
			})
			const amountValidation = await validate(
				this.innerValue.amount,
				rules.amount,
				{
					name: (this.label + ' Amount').trim(),
					values: { type: this.innerValue.type },
				}
			)

			if (!typeValidation.valid) {
				this.message = typeValidation.errors[0]
				this.status = 'is-danger'
			} else if (!amountValidation.valid) {
				this.message = amountValidation.errors[0]
				this.status = 'is-danger'
			} else {
				this.message = null
				this.status = this.innerValue.amount != null ? 'is-success' : null
			}
		},
	},
}
</script>

<style lang="scss" scoped>
// removes space between input and help message on addon fields
.field .field.has-addons {
	margin-bottom: 0;
}
</style>
