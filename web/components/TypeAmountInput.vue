<template>
	<BField :label="label" :type="type" :message="message">
		<BField :type="type">
			<BSelect v-model="innerValue.type" type="text" :placeholder="typePlaceholder">
				<option v-for="(t, k) in types" :key="k" :value="t.value">
					{{ t.name }}
				</option>
			</BSelect>

			<BInput v-model.number="innerValue.value" type="number" :placeholder="valuePlaceholder" />
		</BField>
	</BField>
</template>

<script lang="ts">
import * as yup from 'yup'
import debug from 'debug'
import Vue from 'vue'

const d = debug('ml.components.TypeValueInput')

interface TypeValue {
	type: string | null
	value: number | null
}

export default Vue.extend({
	name: 'TypeValueInput',
	props: {
		// options in Type dropdown
		types: { type: Array as () => Array<string>, required: true },

		// selected value and type
		value: { type: Object as () => TypeValue, default: () => ({ type: null, value: null }) },
		label: { type: String, required: false, default: () => null },
		valuePlaceholder: { type: String, default: () => null },
		typePlaceholder: { type: String, default: () => null },
		schema: {
			type: Object as () => yup.ObjectSchema<TypeValue>,
			default: () =>
				yup.object().shape({
					type: yup.string().nullable().label('Type').required(),
					value: yup.number().nullable().label('Value').required(),
				}),
		},
	},
	data() {
		return {
			innerValue: {
				value: null,
				type: null,
			} as TypeValue,
			errors: [] as string[],
		}
	},
	computed: {
		message(): string | null {
			return this.errors.length > 0 ? this.errors[0] : null
		},
		type(): object {
			return { 'is-danger': this.errors[0] }
		},
	},
	watch: {
		innerValue: {
			handler(newVal) {
				this.$emit('input', newVal.value === '' ? { ...newVal, value: null } : newVal)
			},
			deep: true,
		},
		value: {
			handler(newVal) {
				this.innerValue = newVal
			},
			deep: true,
		},
	},
	created() {
		if (this.value) {
			this.innerValue = this.value
		}
		this.$emit('input', this.innerValue)
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
			this.innerValue = {
				value: null,
				type: null,
			}
			this.errors = []
		},
	},
})
</script>

<style lang="scss" scoped>
// removes space between input and help message on addon fields
.field .field.has-addons {
	margin-bottom: 0;
}
</style>
