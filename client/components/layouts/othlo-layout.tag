othlo-layout.othlo-layout.mdl-layout__content
  othlo-hero-section#hero
  othlo-feature-section#feature
  othlo-create-section#create

  script.
    const riot = require('riot')
    require('../main/othlo-hero-section.tag')
    require('../main/othlo-feature-section.tag')
    require('../main/othlo-create-section.tag')
    riot.mount('othlo-hero-section')
    riot.mount('othlo-feature-section')
    riot.mount('othlo-create-section')
