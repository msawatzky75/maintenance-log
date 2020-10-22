<template>
	<BField v-bind="$attrs" :type="{ 'is-danger': errors[0] }" :message="errors.length ? errors[0] : null">
		<BDatepicker v-model="innerValue" v-bind="$attrs" />
	</BField>
</template>

<script lang="ts">
import type { StringSchema } from 'yup'
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
			default: () => null,
		},
		schema: {
			type: Object as () => StringSchema,
			required: true,
		},
	},
	data() {
		return {
			innerValue: null as Date | null,
			errors: [] as string[],
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
			this.innerValue = new Date(this.value)
		}
	},
	methods: {
		async validate(): Promise<void> {
			try {
				await this.schema.validate(this.innerValue, { abortEarly: false })
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
