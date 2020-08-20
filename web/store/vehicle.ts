import debug from 'debug'
import { MutationTree } from 'vuex'
// https://typescript.nuxtjs.org/cookbook/store.html#vanilla

const d = debug('ml.store.vehicel')

interface Vehicle {
	id: string
}

export function state() {
	return {
		vehicle: null as Vehicle | null,
	}
}

export type RootState = ReturnType<typeof state>

export const mutations: MutationTree<RootState> = {
	set(state, data) {
		d(state, data)
	},
}
