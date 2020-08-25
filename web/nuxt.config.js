export default {
	/*
	 ** Nuxt rendering mode
	 ** See https://nuxtjs.org/api/configuration-mode
	 */
	mode: 'spa',
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
		link: [
			// { rel: 'icon', type: 'image/png', href: require('@/assets/hammer.png') },
		],
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
		baseURL: 'http://localhost:4000/',
		browserBaseURL: 'http://localhost:4000/',
		credentials: true,
	},

	auth: {
		cookie: false,
		strategies: {
			local: {
				scheme: 'refresh',
				token: {
					maxAge: 60 * 15,
					property: false,
					type: false,
				},
				refreshToken: {
					maxAge: 60 * 60 * 24 * 7,
					property: false,
					tokenRequired: false,
				},
				endpoints: {
					user: false,
					login: {
						url: '/api/auth/login',
						method: 'post',
						propertyName: false,
					},
					logout: {
						url: '/api/auth/logout',
						method: 'post',
						propertyName: false,
					},
					refresh: {
						url: '/api/auth/refresh',
						method: 'post',
						propertyName: false,
					},
				},
			},
		},
		redirect: {
			login: '/login',
			logout: '/',
			home: '/home',
		},
	},

	apollo: {
		clientConfigs: {
			default: '@/apollo/client-configs/default.ts',
		},
		errorHandler: '@/plugins/apollo-error-handler.ts', // why even bother? its not like it works
	},

	router: {},

	/*
	 ** Build configuration
	 ** See https://nuxtjs.org/api/configuration-build/
	 */
	build: {
		extend(config, ctx) {
			if (ctx.isDev) {
				config.devtool = ctx.isClient ? 'source-map' : 'inline-source-map'
			}
		},
	},
}
