app
  othlo-header
  main.othlo-layout.mdl-layout__content
    div.page-content
      h1 ここにいろいろ書いていく
  othlo-footer

  script.
    const riot = require('riot')
    require('./components/layouts/othlo-footer.tag')
    require('./components/layouts/othlo-header.tag')
    riot.mount('othlo-header')
    riot.mount('othlo-footer')
