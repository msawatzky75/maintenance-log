<template>
	<p>
		{{ formatNumber(value) }} {{ type }}{{ value > 1 || value == 0 ? 's' : '' }}
	</p>
</template>

<script>
import numeral from 'numeral'

export default {
	name: 'DistanceDisplay',
	props: {
		value: {
			type: Number,
			required: true,
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
