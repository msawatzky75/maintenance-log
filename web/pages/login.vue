<template>
	<section class="section">
		<div class="container">
			<h1 class="is-size-1 has-text-centered">Login</h1>
			<form @submit.prevent="userLogin">
				<BField label="Email">
					<BInput v-model="login.email" type="text" placeholder="Email" />
				</BField>
				<BField label="Password">
					<BInput v-model="login.password" type="password" placeholder="Password" password-reveal />
				</BField>
				<div>
					<BButton type="is-primary" native-type="submit"> Login </BButton>
				</div>
			</form>
		</div>
	</section>
</template>

<script>
import Vue from 'vue'

export default Vue.extend({
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
					this.$store.dispatch('user/fetchUser')
				})
				.catch((err) => {
					console.error(err)
				})
		},
	},
})
</script>
