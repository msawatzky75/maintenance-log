<template>
	<section class="section">
		<h1 class="is-size-1 has-text-centered">Create an account</h1>
		<form @submit.prevent="userSignUp">
			<b-field label="Email" label-for="email">
				<b-input
					id="email"
					v-model="user.email"
					type="email"
					name="email"
					placeholder="Email"
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
					id="verifyEmail"
					v-model="user.verifyEmail"
					name="verifyEmail"
					placeholder="Verify Email"
				/>
			</b-field>
			<b-field label="Password" label-for="password">
				<b-input
					id="password"
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
	</section>
</template>

<script>
export default {
	auth: false,
	data() {
		return {
			user: {
				email: '',
				password: '',
			},
		}
	},
}
</script>

<style></style>
