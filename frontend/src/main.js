import Vue from 'vue'
import App from './App.vue'


//Import Froala Editor 
import 'froala-editor/js/plugins.pkgd.min.js';
//Import third party plugins
import 'froala-editor/js/third_party/embedly.min';
import 'froala-editor/js/third_party/font_awesome.min';
import 'froala-editor/js/third_party/spell_checker.min';
import 'froala-editor/js/third_party/image_tui.min';
// Import Froala Editor css files.
import 'froala-editor/css/froala_editor.pkgd.min.css';

// Import and use Vue Froala lib.
import VueFroala from 'vue-froala-wysiwyg'

import router from './router'
Vue.use(VueFroala)
Vue.config.productionTip = false


Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
