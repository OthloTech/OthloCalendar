'use strict'

//import './tags/index.js'
const riot = require('riot')
require('./tags/app.tag')

// const DummyStore = require('./stores/dummy-store')
// const dispatcher = require('./stores/dispatcher')

// //let dummyStore = new DummyStore(dispatcher)
// //dispatcher.addStore(dummyStore)

// const context = require.context('./tags')
// context.keys().forEach((key) => {
//   context(key)
// })
console.log(1)

riot.mount('app')
//riot.mount('app', {store: dummyStore})