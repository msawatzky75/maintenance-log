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
					<h1 class="is-size-1 has-text-centered">Create an account</h1>
					<form @submit.prevent="userSignUp">
						<b-field label="Email" label-for="email">
							<b-input
								v-model="user.email"
								type="email"
								placeholder="Email"
								name="email"
							/>
						</b-field>

						<b-field
							label="Verify Email"
							label-for="verifyEmail"
							:type="{
								'is-danger': user.verifyEmail && user.verifyEmail != user.email,
							}"
						>
							<b-input
								v-model="user.verifyEmail"
								placeholder="Verify Email"
								name="verifyEmail"
							/>
						</b-field>
						<b-field label="Password">
							<b-input
								v-model="user.password"
								type="password"
								placeholder="Password"
								password-reveal
							/>
						</b-field>
						<div>
							<b-button type="is-primary" native-type="submit">
								Sign Up
							</b-button>
						</div>
					</form>
				</div>

				<div class="column is-6">
					<h1 class="is-size-1 has-text-centered">Login</h1>

					<form @submit.prevent="userLogin">
						<b-field label="Email">
							<b-input v-model="login.email" type="email" placeholder="Email" />
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
			</div>
		</section>
	</section>
</template>

<script>
export default {
	auth: false,
	name: 'HomePage',
	data() {
		return {
			user: {
				email: '',
				verifyEmail: '',
				password: '',
			},
			login: {
				email: '',
				password: '',
			},
		}
	},
	methods: {
		userSignUp() {
			console.log("we're gonna sign up!")
		},
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
