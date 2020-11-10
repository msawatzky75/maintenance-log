# Maintenance Log Tracker Frontend

## Environment Setup

The debug namespace is `ml`, so if you are trying to view log outputs, this is the namespace to use. Set this in your browser's localstorage under the `DEBUG` key. For more details see: [debug](https://www.npmjs.com/package/debug)

```
API_URL=localhost:4000/
```

`API_URL` is required for the application to communicate with the server.

## Build Setup

```bash
# install dependencies
$ yarn install

# serve with hot reload at localhost:3000
$ yarn dev

# build for production and launch server
$ yarn build
$ yarn start

# generate static project
$ yarn generate
```

For detailed explanation on how things work, check out [Nuxt.js docs](https://nuxtjs.org).
