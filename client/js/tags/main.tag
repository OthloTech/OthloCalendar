main
  .ui.container
    h1 Hello world
    h2 okokok
    article
      ul
        li.list(each="{ name, i in items }")
          {name}

  script(type='text/javascript').
    console.log('ok')
    this.items = [1, 2, 3, 4, 5]

  style(scoped).
    .ui.container {
      margin-top: 50px;
    }

    h2 {
      color: white;
      background-color: black;
    }