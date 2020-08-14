export default {
	/*
	 ** Nuxt rendering mode
	 ** See https://nuxtjs.org/api/configuration-mode
	 */
	mode: 'universal',
	/*
	 ** Nuxt target
	 ** See https://nuxtjs.org/api/configuration-target
	 */
	target: 'static',
	/*
	 ** Headers of the page
	 ** See https://nuxtjs.org/api/configuration-head
	 */
	head: {
		title: 'Maintence Log Tracker',
		meta: [
			{ charset: 'utf-8' },
			{ name: 'viewport', content: 'width=device-width, initial-scale=1' },
			{
				hid: 'description',
				name: 'description',
				content: process.env.npm_package_description || '',
			},
		],
		link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
	},
	/*
	 ** Global CSS
	 */
	css: ['@/assets/scss/main.scss'],
	/*
	 ** Plugins to load before mounting the App
	 ** https://nuxtjs.org/guide/plugins
	 */
	plugins: [],
	/*
	 ** Auto import components
	 ** See https://nuxtjs.org/api/configuration-components
	 */
	components: true,
	/*
	 ** Nuxt.js dev-modules
	 */
	buildModules: ['@nuxt/typescript-build'],
	/*
	 ** Nuxt.js modules
	 */
	modules: [
		// Doc: https://buefy.github.io/#/documentation
		'nuxt-buefy',
		'@nuxtjs/pwa',

		'@nuxtjs/axios',
		'@nuxtjs/auth-next',
		'@nuxtjs/apollo',
	],

	axios: {
		baseURL: 'http://localhost:4000',
	},
	auth: {
		cookie: false,
		strategies: {
			local: {
				scheme: 'refresh',
				token: {
					property: 'access_token',
					maxAge: 60 * 15, // 15 minutes
					type: 'Bearer',
				},
				refreshToken: {
					property: 'refresh_token',
					data: 'refresh_token',
					maxAge: 60 * 60 * 24 * 7, // 7 days
				},
				endpoints: {
					login: { url: '/api/auth/login', method: 'post' },
					refresh: { url: '/api/auth/refresh', method: 'post' },
					user: false,
				},
			},
		},
	},
	apollo: {
		authenticationType: 'Bearer',
		tokenName: 'auth._token.local',
		clientConfigs: {
			default: '@/apollo/client-configs/default.ts',
		},
	},
	/*
	 ** Build configuration
	 ** See https://nuxtjs.org/api/configuration-build/
	 */
	build: {},
}
