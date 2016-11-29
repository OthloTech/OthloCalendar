app
  .ui.container
    h1 Hello world
    article
      ul
        li.list(each="{ name, i in items }")
          {name}
    sample

  script(type='text/javascript').
    // update the current time
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
