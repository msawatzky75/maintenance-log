<template>
	<section class="section">
		<h1 class="is-size-1 has-text-centered">Create an account</h1>

		<form @submit.prevent="signUp">
			<BField
				label="Email"
				label-for="email"
				:message="emailError ? emailError.message : null"
				:type="{ 'is-danger': emailError }"
			>
				<BInput id="email" v-model="user.email" name="email" placeholder="Email" />
			</BField>

			<BField
				label="Verify Email"
				label-for="verifyEmail"
				:message="verifyEmailError ? verifyEmailError.message : null"
				:type="{ 'is-danger': verifyEmailError }"
			>
				<BInput id="verifyEmail" v-model="user.verifyEmail" name="verifyEmail" placeholder="Verify Email" />
			</BField>

			<BField
				label="Password"
				label-for="password"
				:message="passwordError ? passwordError.message : null"
				:type="{ 'is-danger': passwordError }"
			>
				<BInput id="password" v-model="user.password" type="password" placeholder="Password" password-reveal />
			</BField>

			<div>
				<BButton type="is-primary" native-type="submit">Sign Up</BButton>
			</div>
		</form>
		<pre v-if="networkError">{{ networkError }}</pre>
	</section>
</template>

<script>
import debug from 'debug'
import * as yup from 'yup'

const d = debug('ml.web.pages.signup')
const schema = yup.object().shape({
	email: yup.string().label('Email').email().required(),
	verifyEmail: yup.string().test('equalToEmail', 'Verify Email must be the same as Email', function () {
		return this.parent.email === this.parent.verifyEmail
	}),
	password: yup.string().min(8).label('Password').required(),
})

export default {
	auth: false,
	data() {
		return {
			user: {
				email: '',
				verifyEmail: '',
				password: '',
			},
			error: null,
			networkError: null,
		}
	},
	computed: {
		emailError() {
			return this.error && this.error.inner.find((e) => e.path === 'email')
		},
		verifyEmailError() {
			return this.error && this.error.inner.find((e) => e.path === 'verifyEmail')
		},
		passwordError() {
			return this.error && this.error.inner.find((e) => e.path === 'password')
		},
	},
	methods: {
		async signUp() {
			this.error = null
			try {
				schema.validateSync(this.user, { abortEarly: false })
			} catch (e) {
				this.error = e
				return
			}

			try {
				const response = await this.$axios.put('/api/signup', {
					email: this.user.email,
					password: this.user.password,
					// credentials: 'omit',
				})
				d('this should not be undefined', response)
				// 200 OK
				if (response.status === 200) {
					this.$auth
						.loginWith('local', {
							data: {
								email: this.user.email,
								password: this.user.password,
							},
						})
						.then(() => {
							this.$store.dispatch('user/fetchUser')
						})
						.catch((err) => {
							// eslint-disable-next-line
							console.error(err)
						})
				}
			} catch (e) {
				// 409 Conflict
				if (e.response.status === 409) {
					this.error = {
						inner: [
							{
								path: 'email',
								message: e.response.data.error,
							},
						],
					}
				}
				// 400 Bad Request
				else if (e.response.status === 400) {
					this.error = {
						inner: [
							{
								path: 'password',
								message: e.response.data.error,
							},
						],
					}
				}
				// literally anything else
				else {
					throw e
				}
				this.networkError = e.response.data.error
			}
		},
	},
}
</script>

<style></style>
