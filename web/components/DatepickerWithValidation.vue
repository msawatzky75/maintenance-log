<template>
	<BField
		v-bind="$attrs"
		:type="{ 'is-danger': errors[0] }"
		:message="errors.length ? errors[0] : null"
	>
		<BDatepicker v-model="innerValue" v-bind="$attrs" />
	</BField>
</template>

<script lang="ts">
import * as yup from 'yup'
import debug from 'debug'
import Vue from 'vue'

const d = debug('ml.components.DatePickerWithValidation')
export default Vue.extend({
	name: 'DatepickerWithValidation',
	props: {
		// must be included in props
		value: {
			type: Date,
			default: null,
		},
		type: {
			type: Object,
			default: null,
		},
		schema: {
			type: Object,
			default() {
				return yup.mixed().nullable().label(this.$attrs.label).required()
			},
		},
	},
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
})
</script>
