<template>
	<div>
		<b-form @submit="userLogin">
			<div>
				<label>Username</label>
				<input v-model="login.username" type="text" />
			</div>
			<div>
				<label>Password</label>
				<input v-model="login.password" type="text" />
			</div>
			<div>
				<b-button type="submit">Submit</b-button>
			</div>
		</b-form>
	</div>
</template>

<script>
export default {
	data() {
		return {
			login: {
				username: '',
				password: '',
			},
		}
	},
	methods: {
		async userLogin() {
			try {
				const response = await this.$auth.loginWith('local', {
					data: this.login,
				})
				console.log(response)
			} catch (err) {
				console.log(err)
			}
		},
	},
}
</script>
