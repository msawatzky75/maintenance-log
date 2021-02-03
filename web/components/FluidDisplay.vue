<template>
	<p v-if="value">
		{{ formatNumber(value) }} {{ short ? $store.state.fluidTypes.find((v) => v.value === type).short : type
		}}{{ (value > 1 || value === 0) && !short ? 's' : '' }}
	</p>
	<p v-else>&mdash;</p>
</template>

<script lang="ts">
import numeral from 'numeral'
import Vue from 'vue'

export default Vue.extend({
	name: 'FluidDisplay',
	props: {
		value: {
			type: Number,
			default: () => 0,
		},
		type: {
			type: String,
			required: true,
		},
		short: Boolean,
	},
	methods: {
		formatNumber(n: string) {
			return numeral(n).format('0,0')
		},
	},
})
</script>
