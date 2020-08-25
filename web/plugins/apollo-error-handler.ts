import type { Context } from '@nuxt/types'
import debug from 'debug'

const d = debug('ml.plugins.apollo-error-handler')

export default function (
	{ graphQLErrors, networkError, operation, forward }: any,
	nuxtContext: Context
) {
	// this is bullshit, it doesnt run at all...?
	d('Global error handler')
	// console.log(graphQLErrors, networkError, operation, forward)
}
