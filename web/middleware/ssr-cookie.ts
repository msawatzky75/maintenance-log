import { Context } from '@nuxt/types'
import debug from 'debug'

const d = debug('ml.middleware.ssr-cookie')

export default function ({ req, app }: Context) {
	if (process.server) {
		d('Forwarding cookies')
		app.$axios.defaults.headers.common.cookie = req.headers.cookie
	}
}
