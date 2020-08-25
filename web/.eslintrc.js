module.exports = {
	root: true,
	env: {
		browser: true,
		node: true,
	},
	extends: [
		'@nuxtjs/eslint-config-typescript',
		'prettier',
		'prettier/vue',
		'plugin:prettier/recommended',
		'plugin:nuxt/recommended',
	],
	plugins: ['prettier'],
	rules: {
		indent: ['error', 'tab'],
		'no-unused-vars': 'off',
		'@typescript-eslint/no-unused-vars': [
			'warn',
			{ varsIgnorePattern: '_|d', args: 'after-used' },
		],
	},
}
