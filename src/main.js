import Vue from 'vue'
import Home from './views/Home'

Vue.config.productionTip = false

if (document.querySelector('#app')) {
  new Vue({
    components: {
      Home
    }
  }).$mount('#app')
}

console.log('test');
