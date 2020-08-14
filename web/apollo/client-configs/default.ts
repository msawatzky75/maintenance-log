export default (context: any) => {
	return {
		httpEndpoint: 'http://localhost:4000/graphql',
		getAuth() {
			console.log(context)
			return localStorage.getItem('auth._token.local')
		},
	}
}
