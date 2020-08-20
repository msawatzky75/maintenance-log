import {
	InMemoryCache,
	IntrospectionFragmentMatcher,
} from 'apollo-cache-inmemory'
import introspectionQueryResultData from './fragments.json'

export default function () {
	return {
		httpEndpoint: 'http://localhost:4000/graphql',
		httpLinkOptions: {
			credentials: 'include',
		},
		cache: new InMemoryCache({
			fragmentMatcher: new IntrospectionFragmentMatcher({
				introspectionQueryResultData,
			}),
		}),
	}
}
