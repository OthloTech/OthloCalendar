require('./components/layouts/basic-layout.tag')

// require('./components/layouts/home-layout.tag');
// <app>
//     <basic-layout if="{{isHome}}" />
//     <home-layout if="{{!isHome}}" />
//     <script type="es6">
//         this.isHome = !!window.location.hash.replace('#', '');
//         riot.route((collection, id, action)=> {
//             this.update({isHome:!!collection})
//         });
//     </script>
// </app>

app
  div.mdl-layout.mdl-js-layout
    header.mdl-layout__header.mdl-layout__header--scroll
      div.mdl-layout__header-row
        span.mdl-layout-title OthloCalendar
        div.mdl-layout-spacer
        nav.mdl-navigation
          a.mdl-navigation__link(href='#about') about
          a.mdl-navigation__link(href='#feature') 特徴
          a.mdl-navigation__link(href='#function') 機能
    div.mdl-layout__drawer
      span.mdl-layout-title OthloCalendar
      nav.mdl-navigation
        a.mdl-navigation__link(href='#about') about
        a.mdl-navigation__link(href='#feature') 特徴
        a.mdl-navigation__link(href='#function') 機能
    main.mdl-layout__content
      div.page-content
        h1 ここにいろいろ書いていく
