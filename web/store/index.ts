export function state() {
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
		logTypes: {
			FuelLog: 'Fuel Log',
			MaintenceLog: 'Maintence Log',
			OilChangeLog: 'Oil Change Log',
		},
	}
}
