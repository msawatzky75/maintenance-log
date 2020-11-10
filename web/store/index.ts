import type { GetterTree } from 'vuex'
import CreateFuelLog from '~/apollo/mutations/create-fuel-log.graphql'
import CreateMaintenanceLog from '~/apollo/mutations/create-maintenance-log.graphql'
import CreateOilChangeLog from '~/apollo/mutations/create-oil-change-log.graphql'
import DeleteFuelLog from '~/apollo/mutations/deleteFuelLog.graphql'
import DeleteMaintenanceLog from '~/apollo/mutations/deleteMaintenanceLog.graphql'
import DeleteOilChangeLog from '~/apollo/mutations/deleteOilChangeLog.graphql'

function state() {
	return {
		distanceTypes: [
			{ value: 'Kilometre', name: 'Kilometre', short: 'km' },
			{ value: 'Mile', name: 'Mile', short: 'mi' },
		],
		currencyTypes: [
			{ value: 'CAD', name: 'CAD' },
			{ value: 'USD', name: 'USD' },
		],
		fluidTypes: [
			{ value: 'Litre', name: 'Litre' },
			{ value: 'Gallon', name: 'Gallon' },
		],
		logTypes: [
			{
				value: 'FuelLog',
				name: 'Fuel Log',
				deleteMutation: DeleteFuelLog,
				createMutation: CreateFuelLog,
			},
			{
				value: 'MaintenanceLog',
				name: 'Maintenance Log',
				deleteMutation: DeleteMaintenanceLog,
				createMutation: CreateMaintenanceLog,
			},
			{
				value: 'OilChangeLog',
				name: 'Oil Change Log',
				deleteMutation: DeleteOilChangeLog,
				createMutation: CreateOilChangeLog,
			},
		],
	}
}

export type RootState = ReturnType<typeof state>

export default {
	state,
	getters: {
		logType: (state) => (type: string) => state.logTypes.find((l) => l.value === type),
	} as GetterTree<RootState, RootState>,
}
