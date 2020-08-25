<template>
	<div>
		<section class="section">
			Select a vehicle to work with:
			<b-select v-model="vehicleId" placeholder="Select a vehicle">
				<option v-for="v in vehicles" :key="v.id" :value="v.id">
					{{ v.year }} {{ v.make }} {{ v.model }}
				</option>
			</b-select>
			<b-button class="is-primary" @click="selectVehicle(vehicleId)">
				Select
			</b-button>
		</section>
	</div>
</template>

<script lang="ts">
import Vue from 'vue'
import type { Context } from '@nuxt/types'
import debug from 'debug'

const d = debug('ml.pages.vehicle.index')

export default Vue.extend({
	validate({ store }: Context) {
		d(store.state.user.vehicles)
		return store.state.user.vehicles && store.state.user.vehicles.length > 0
	},
	data() {
		return {
			vehicleId: this.$store.state.selcetedVehicle,
		}
	},
	computed: {
		vehicles() {
			return this.$store.state.user.vehicles
		},
	},
	methods: {
		selectVehicle(vid: string) {
			this.$store.commit('user/setDefaultVehicle', vid)
			this.$router.push({ path: '/vehicle/' + vid })
		},
	},
})
</script>
