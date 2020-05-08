import Vue from 'vue'
import App from './App.vue'
import Home from './views/Home'

if (document.querySelector('#app')) {
  new Vue({
    components: {
      App,
      Home
    },
    render: h => h(Home)
  }).$mount('#app')
}

console.log('test')
