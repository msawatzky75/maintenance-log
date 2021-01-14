<template>
	<BField v-bind="$attrs" :type="{ 'is-danger': errors[0] }" :message="errors.length ? errors[0] : null">
		<BDatepicker v-model="innerValue" v-bind="$attrs" locale="en-CA" />
	</BField>
</template>

<script lang="ts">
import type { StringSchema } from 'yup'
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
			default: () => null as Date | null,
		},
		type: {
			type: Object,
			default: () => null as Object | null,
		},
		schema: {
			type: Object as () => StringSchema | null,
			default: () => yup.string().nullable().label('Date'),
		},
		required: {
			type: Boolean,
			default: () => false,
		},
	},
	data() {
		return {
			innerValue: null as Date | null,
			errors: [] as string[],
		}
	},
	computed: {
		// Must specify return type here or typescript gets upset
		validationSchema(): StringSchema {
			const schema = (this.schema as StringSchema) || yup.string().nullable().label('Date')

			if (this.required) {
				return schema.required()
			}
			return schema
		},
	},
	watch: {
		// Handles internal model changes.
		innerValue(newVal) {
			this.$emit('input', newVal === '' ? (null as string | null) : newVal)
		},
		// Handles external model changes.
		value(newVal) {
			this.innerValue = newVal
		},
	},
	created() {
		if (this.value) {
			this.innerValue = new Date(this.value)
		}
	},
	methods: {
		async validate(): Promise<void> {
			try {
				await this.validationSchema.validate(this.innerValue, { abortEarly: false })
				this.errors = []
			} catch (e) {
				this.errors = e.errors
				throw e
			}
		},

		reset(): void {
			this.innerValue = null
			this.errors = []
		},
	},
})
</script>
