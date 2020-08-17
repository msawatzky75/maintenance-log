<template>
	<section class="section">
		<div class="container">
			<h1 class="is-size-1 has-text-centered">Login</h1>
			<form @submit.prevent="userLogin">
				<b-field label="Email">
					<b-input v-model="login.email" type="text" placeholder="Email" />
				</b-field>
				<b-field label="Password">
					<b-input
						v-model="login.password"
						type="password"
						placeholder="Password"
						password-reveal
					/>
				</b-field>
				<div>
					<b-button type="is-primary" native-type="submit">
						Login
					</b-button>
				</div>
			</form>
		</div>
	</section>
</template>

<script>
export default {
	data() {
		return {
			login: {
				email: '',
				password: '',
			},
		}
	},
	methods: {
		userLogin() {
			return this.$auth
				.loginWith('local', {
					data: this.login,
				})
				.then(() => {
					console.log('Success')
				})
		},
	},
}
</script>
