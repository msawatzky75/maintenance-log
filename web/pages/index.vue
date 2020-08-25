<template>
	<section>
		<section class="section">
			<h1 class="is-size-1 has-text-centered">Maintenence Log Tracker</h1>
			<p class="is-size-4 has-text-centered">
				A simple application that allows you to track various maintenence logs
				of your vehicle.
			</p>
		</section>

		<section class="section">
			<div class="columns">
				<div class="column is-6">
					<p v-if="!$auth.loggedIn" class="is-size-4 has-text-centered">
						You need to
						<nuxt-link class="button is-link" to="/login">Login</nuxt-link>
						or
						<nuxt-link class="button is-link" to="/signup">Sign Up</nuxt-link>
						to use the app.
					</p>
					<p v-else>
						Welcome {{ $store.state.user.email }}
						<nuxt-link class="button is-link" to="/home">
							Home
						</nuxt-link>
					</p>
				</div>
			</div>
		</section>
	</section>
</template>

<script>
export default {
	auth: false,
	name: 'HomePage',
}
</script>
