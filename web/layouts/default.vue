<template>
	<div>
		<nav
			class="navbar header is-primary"
			role="navigation"
			aria-label="main navigation"
		>
			<div class="navbar-brand">
				<a class="navbar-item" href="/">
					<img src="~assets/buefy.png" alt="Buefy" height="28" />
				</a>

				<div class="navbar-burger">
					<span />
					<span />
					<span />
				</div>
			</div>
		</nav>

		<section class="main-content">
			<div class="container">
				<nuxt />
			</div>
		</section>
	</div>
</template>

<script>
import debug from 'debug'

const d = debug('ml.layouts.default')

export default {
	middleware: [
		'auth',
		async function ({ redirect, store }) {
			await store.dispatch('user/fetchUser')
		},
	],
}
</script>
