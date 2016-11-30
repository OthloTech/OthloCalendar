app
  othlo-header
  othlo-layout
  othlo-footer

  script.
    const riot = require('riot')
    require('./components/layouts/othlo-header.tag')
    require('./components/layouts/othlo-layout.tag')
    require('./components/layouts/othlo-footer.tag')
    riot.mount('othlo-header')
    riot.mount('othlo-layout')
    riot.mount('othlo-footer')
