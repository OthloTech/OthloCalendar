'use strict'

require('./materialize.js')
require('./tags/index.js')
const riot = require('riot')

// const DummyStore = require('./stores/dummy-store')
// const dispatcher = require('./stores/dispatcher')

// //let dummyStore = new DummyStore(dispatcher)
// //dispatcher.addStore(dummyStore)

// const context = require.context('./tags')
// context.keys().forEach((key) => {
//   context(key)
// })

riot.mount('app')
//riot.mount('app', {store: dummyStore})
