<template>
	<section>
		<b-field label="Date">
			<b-datepicker
				v-model="innerValue.date"
				placeholder="Click to select..."
				icon="calendar-today"
				trap-focus
			>
			</b-datepicker>
		</b-field>

		<b-field label="Notes">
			<b-input v-model="innerValue.notes" type="textarea" />
		</b-field>
	</section>
</template>

<script>
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
			innerValue: {
				date: null,
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
}
</script>
