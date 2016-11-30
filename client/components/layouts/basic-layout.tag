require('./basic-layout.scss')
require('material-elements/material-navbar/material-navbar.tag');
require('material-elements/material-pane/material-pane.tag');
require('material-elements/material-button/material-button.tag');
//require('riotmui-elements/riotmui-list/riotmui-list.tag');
//require('riot-router');
basic-layout
  h1 okokok
  material-navbar
    div.row
      div.col.col-lg-3.col-md-6.col-sm-11.col-xs-11
        div.logo
          i.material-icons.menu(onclick="{{parent.toggleMenu}}") menu
          a(href='#' title='Material UI'){{parent.logoText}}
      div.col.col-lg-9.col-md-6.col-sm-1.col-xs-1.gitcol
        a.github GitHub
  
  div.row.content

  div.overlay(name='overlay' if='{{opened}}' onclick='{{close}}')

  script(type='text/javascript').
    console.log(1)
    this.opened = false
    this.logoText = 'material UI'
    this.open = function() {
      this.update({opened: true})
      this.menu.style.left + '0px'
    }
    this.close = function() {
      this.update({opened: false})
      this.menu.style.left = '-100%'
    }
