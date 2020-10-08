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
		'plugin:vue/vue3-recommended',
	],
	plugins: ['prettier'],
	rules: {
		indent: ['error', 'tab'],
		'vue/html-indent': ['error', 'tab'],
		'vue/html-self-closing': [
			'warn',
			{
				html: {
					void: 'always',
				},
			},
		],
		'no-unused-vars': 'off',
		'@typescript-eslint/no-unused-vars': [
			'warn',
			{ varsIgnorePattern: '_|d', args: 'after-used' },
		],
		'sort-imports': ['error', { ignoreCase: true }],
	},
}
