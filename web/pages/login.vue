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
					<b-button
						type="is-primary"
						native-type="submit"
						:disabled="loginable"
					>
						Login
					</b-button>
					<b-button type="is-secondary" @click.prevent="$auth.refreshTokens">
						refresh
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
	computed: {
		loginable() {
			return this.email && this.password.length > 12
		},
	},
	methods: {
		async userLogin() {
			try {
				await this.$auth.loginWith('local', {
					data: this.login,
				})
			} catch (err) {
				console.log(err)
			}
		},
	},
}
</script>
