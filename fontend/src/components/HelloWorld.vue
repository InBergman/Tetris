<template>
  <div class="hello">
    <img src="../assets/logoTetris.png">
    <h1>{{ msg }}</h1>
    <div id="score"></div>
    <canvas id="canvastetris" width="240" height="400"></canvas>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
img {
      height: 200px;
      width: 200px;
}

 html {
   border : 4px solid Fuchsia;
    background-image: linear-gradient(MediumSpringGreen, Aqua);
}
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>