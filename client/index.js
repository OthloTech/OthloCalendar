'use strict'
//require('riot-mui')
const riot = require('riot')
require('./app.tag')

riot.settings.brackets = '{{ }}'
riot.mount('*')

// Router
//require('./router.js')
// SASS
require('./stylesheets/normalize.scss')
require('./stylesheets/flex-grid.scss')
require('./stylesheets/default-hljs.scss')
require('material-design-lite/material.min.css')
require('material-design-lite/material.min.js')
require('./stylesheets/style.scss')
