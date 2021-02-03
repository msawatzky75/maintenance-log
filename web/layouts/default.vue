<template>
	<div>
		<nav class="navbar header is-primary" role="navigation" aria-label="main navigation">
			<div class="navbar-brand">
				<NuxtLink class="navbar-item" to="/">
					<!-- <img src="~assets/buefy.png" alt="Buefy" height="28" /> -->
					<h1 class="is-size-4 has-text-weight-bold">Maintenance Manager II; Electric Boogaloo</h1>
				</NuxtLink>

				<div class="navbar-burger">
					<span />
					<span />
					<span />
				</div>
			</div>
			<div class="navbar-menu">
				<div class="navbar-end">
					<div v-if="$auth.loggedIn" class="navbar-item has-dropdown is-hoverable">
						<a class="navbar-link">{{ $store.state.user.email }}</a>

						<div class="navbar-dropdown is-right">
							<NuxtLink class="navbar-item" :to="{ name: 'profile' }">Profile</NuxtLink>
							<a class="navbar-item" @click="$auth.logout()">Logout</a>
						</div>
					</div>
				</div>
			</div>
		</nav>

		<section class="main-content">
			<div class="container">
				<Nuxt />
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
		async function ({ store }) {
			if (!store.user) {
				await store.dispatch('user/fetchUser')
			}
		},
	],
}
</script>
