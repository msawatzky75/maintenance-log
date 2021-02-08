import { InMemoryCache, IntrospectionFragmentMatcher } from 'apollo-cache-inmemory'
import { onError } from 'apollo-link-error'
import type { ServerError } from 'apollo-link-http-common'
import debug from 'debug'
import introspectionQueryResultData from './fragments.json'

const d = debug('ml.apollo.client-configs.default')

function isServerError(e: any): e is ServerError {
	return e.statusCode && e instanceof Error
}

export default function ({
	$auth,
	env,
	...rest
}: {
	$auth: { refreshTokens: Function }
	env: { API_URL_GRAPHQL: string }
}) {
	const link = onError(({ networkError }) => {
		if (networkError && $auth) {
			if (isServerError(networkError) && networkError.statusCode === 401) {
				d('refresh', $auth, rest)
				$auth.refreshTokens()
			}
		}
	})
	// .concat(
	// 	new HttpLink({
	// 		uri: process.env.API_URL ? process.env.API_URL + '/graphql' : 'http://localhost:4000/graphql',
	// 		credentials: 'include',
	// 	})
	// )

	return {
		httpEndpoint: env.API_URL_GRAPHQL,
		httpLinkOptions: {
			credentials: 'include',
		},
		link,
		cache: new InMemoryCache({
			fragmentMatcher: new IntrospectionFragmentMatcher({
				introspectionQueryResultData,
			}),
		}),
	}
}
