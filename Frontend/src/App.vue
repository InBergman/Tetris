<template>
  <div id="app">
    <img src="./assets/logoVue.png">
    <HelloWorld msg="Sync our head"/>
      <p>All Account: "{{ message }}"</p>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.use(VueAxios, axios)

export default {
  name: 'app',
  data: function () {
    return {
      message: ''
    }
  },
  components: {
    HelloWorld
  },
  created () {
    this.getAccount()
  },
  methods: {
    getAccount: function () {
      this.message = 'dude'
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