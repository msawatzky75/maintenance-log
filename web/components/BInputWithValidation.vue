<template>
	<validation-provider
		v-slot="{ errors, valid }"
		:vid="vid"
		:name="$attrs.name || $attrs.label"
		:rules="rules"
		slim
	>
		<b-field
			v-bind="$attrs"
			:type="{ 'is-danger': errors[0], 'is-success': valid }"
			:message="errors"
		>
			<b-input v-model="innerValue" v-bind="$attrs"></b-input>
		</b-field>
	</validation-provider>
</template>

<script>
import { ValidationProvider, extend } from 'vee-validate'

extend('required', {
	validate: (v) => ({
		required: true,
		valid: !['', null, undefined].includes(v),
	}),
	message: '{_field_} is required.',
	computesRequired: true,
})
extend('numeric', {
	validate: (v) => ({
		required: true,
		valid: /\d+/.test(v),
	}),
	message: '{_field_} must be a number.',
})

export default {
	components: {
		ValidationProvider,
	},
	props: {
		vid: {
			type: String,
			default: null,
		},
		rules: {
			type: [Object, String],
			default: '',
		},
		// must be included in props
		value: {
			type: null,
			default: null,
		},
	},
	data: () => ({
		innerValue: '',
	}),
	watch: {
		// Handles internal model changes.
		innerValue(newVal) {
			this.$emit('input', newVal)
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
}
</script>
