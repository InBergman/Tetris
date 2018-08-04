<template>
  <div id="app">
    <HelloWorld msg="Tetris Game"/>
  </div>
</template>

<script>
/* eslint-disable */
import HelloWorld from './components/HelloWorld.vue'
export default {
  name: 'app',
  data: function()
  {
    return {
      message: '',
      matrice: [  [0, 0, 0],
                  [0, 1, 0],
                  [1, 1, 1]  ],
      player: {}
    }
  },
  components:
  {
    HelloWorld
  },
  mounted()
  {
    this.init();
  },
  methods: 
  {
    init: function() 
    {
      this.canvas   = document.getElementById('canvastetris');
      this.context  = this.canvas.getContext('2d');
      this.player = 
      {
        pos: {x : 5, y : 5},
        matrice: this.matrice
      };

      this.context.scale(20, 20);
      this.context.fillStyle = '#000';
      this.context.fillRect(0, 0, this.canvas.width, this.canvas.height);
      
      this.drawMatrice(this.player.matrice, this.player.pos);
    },

    drawMatrice: function(matrice, offset)
    {
      this.matrice.forEach((row, x) => {
        row.forEach((col, y) => {
          if(col === 1) 
          {
            this.context.fillStyle = 'red';
            this.context.fillRect(x + offset.x,
                                  y + offset.y,
                                  1, 1);
          }
        })
      })
    }
  }
}
</script>

<style>

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
