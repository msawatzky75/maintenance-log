<template>
	<p>
		{{ formatNumber(value) }} {{ type
		}}{{ value > 1 || value === 0 ? 's' : '' }}
	</p>
</template>

<script>
import numeral from 'numeral'

export default {
	name: 'DistanceDisplay',
	props: {
		value: {
			type: Number,
			default: () => 0,
		},
		type: {
			type: String,
			required: true,
		},
	},
	methods: {
		formatNumber(n) {
			return numeral(n).format('0,0')
		},
	},
}
</script>
