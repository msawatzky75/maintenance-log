import {
	InMemoryCache,
	IntrospectionFragmentMatcher,
} from 'apollo-cache-inmemory'
import debug from 'debug'
import introspectionQueryResultData from './fragments.json'

const d = debug('ml.apollo.client-configs.default')

export default async function () {
	if (process.client) {
		const httpEndpoint = 'http://localhost:4000/graphql'

		return {
			httpEndpoint,
			cache: new InMemoryCache({
				fragmentMatcher: new IntrospectionFragmentMatcher({
					introspectionQueryResultData,
				}),
			}),
		}
	}
}
