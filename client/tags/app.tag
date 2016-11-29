app
  .ui.container
    h1 Hello world
    h2 The current time is {currentTime}
    .ui.red.button(if="{timer}" onclick="{stop}") Stop Time
    .ui.blue.button(if="{!timer}" onclick="{start}") Start Time

  script(type='text/javascript').
    // update the current time
    console.log('ok')

  style(scoped).
    .ui.container {
      margin-top: 50px;
    }

    h2 {
      color: white;
      background-color: black;
    }


