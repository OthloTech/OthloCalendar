'use strict'
const riot = require('riot')

require('./app.tag')
riot.mount('app')

// Router
//require('./router.js')
// SASS
require('./stylesheets/normalize.scss')
require('./stylesheets/flex-grid.scss')
require('./stylesheets/default-hljs.scss')
