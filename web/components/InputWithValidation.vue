<template>
	<BField v-bind="$attrs" :type="{ 'is-danger': errors[0] }" :message="errors.length ? errors[0] : null">
		<BInput v-model="innerValue" v-bind="$attrs" :type="inputType" />
	</BField>
</template>

<script lang="ts">
import type { StringSchema } from 'yup'
import Vue from 'vue'
import debug from 'debug'

const d = debug('ml.components.InputWithValidation')
export default Vue.extend({
	name: 'InputWithValidation',
	props: {
		// must be included in props
		value: {
			type: [String, Number],
			default: () => null as string | number | null,
		},
		type: {
			type: String,
			default: () => null,
		},
		inputType: {
			type: String,
			default: () => 'text',
		},
		schema: {
			type: Object as () => StringSchema,
			required: true,
		},
	},
	data() {
		return {
			innerValue: null as string | number | null,
			errors: [],
		}
	},
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
})
</script>
