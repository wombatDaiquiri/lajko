import './assets/base.css'
import { ApolloClient, createHttpLink, InMemoryCache } from '@apollo/client/core'

// HTTP connection to the API
const httpLink = createHttpLink({
    // You should use an absolute URL here
    uri: 'http://localhost:8080/graphql',
})

// Cache implementation
const cache = new InMemoryCache()

// Create the apollo client
const apolloClient = new ApolloClient({
    link: httpLink,
    cache,
})

// below should be options-specific

import { createApolloProvider } from '@vue/apollo-option'

const apolloProvider = createApolloProvider({
    defaultClient: apolloClient,
})

import { createApp, h } from 'vue'
import App from './App.vue'

const app = createApp({
    render: () => h(App),
})

app.use(apolloProvider)

app.mount('#app');