<template>
	<BField
		v-bind="$attrs"
		:type="{ 'is-danger': errors[0] }"
		:message="errors.length ? errors[0] : null"
	>
		<BInput v-model="innerValue" v-bind="$attrs" />
	</BField>
</template>

<script lang="ts">
import * as yup from 'yup'
import debug from 'debug'

const d = debug('ml.components.InputWithValidation')
export default {
	name: 'InputWithValidation',
	props: {
		// must be included in props
		value: {
			type: null,
			default: null,
		},
		type: {
			type: String,
			default: null,
		},
		schema: {
			type: Object,
			default() {
				return yup.string().nullable().label(this.$attrs.label).required()
			},
		},
	},
	emits: ['input'],
	data: () => ({
		innerValue: null,
		errors: [],
	}),
	watch: {
		// Handles internal model changes.
		innerValue(newVal) {
			this.$emit('input', newVal === '' ? null : newVal)
		},
		// Handles external model changes.
		value(newVal) {
			this.innerValue = newVal
		},
	},
	created() {
		if (this.value) {
			this.innerValue = this.value
		}
	},
	methods: {
		async validate() {
			try {
				await this.schema.validate(this.innerValue, { abortEarly: false })
				this.errors = []
			} catch (e) {
				this.errors = e.errors
				throw e
			}
		},
		reset() {
			this.innerValue = null
			this.errors = []
		},
	},
}
</script>
