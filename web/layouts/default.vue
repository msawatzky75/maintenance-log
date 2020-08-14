<template>
	<div>
		<nav
			class="navbar header has-shadow is-primary"
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
export default {
	middleware: 'auth',
	data() {
		return {
			items: [
				{
					title: 'Home',
					icon: 'home',
					to: { name: 'index' },
				},
				{
					title: 'Inspire',
					icon: 'lightbulb',
					to: { name: 'inspire' },
				},
			],
		}
	},
}
</script>
