<template>
	<b-field :label="label">
		<b-field>
			<b-select
				:value="type"
				type="text"
				placeholder="Distance Type"
				@input="updateType"
			>
				<option value="Kilometre" selected>Kilometer</option>
				<option value="Mile">Mile</option>
			</b-select>
			<b-input
				:value="amount"
				type="number"
				placeholder="Distance Value"
				@input="updateAmount"
			/>
		</b-field>
	</b-field>
</template>

<script>
import debug from 'debug'

const d = debug('ml.components.odometer-input')

export default {
	name: 'OdometerInput',
	components: {},
	props: {
		value: {
			type: Object,
			required: true,
		},
		label: {
			type: String,
			default: null,
		},
	},
	data() {
		return {
			amount: this.value.value,
			type: this.value.type,
		}
	},
	methods: {
		updateType(t) {
			this.type = t
			this.update()
		},
		updateAmount(a) {
			this.amount = a
			this.update()
		},
		update() {
			this.$emit('input', { value: this.amount, type: this.type })
		},
	},
}
</script>
