<template>
  <div id="app">
    <froala
      id="edit"
      :tag="'textarea'"
      :config="config"
      v-model="model"
    ></froala>
  </div>
</template>

<script>
// import VueFroala from 'vue-froala-wysiwyg';

export default {
  name: "app",
  data() {
    return {
      config: {
        events: {
          initialized: function () {
            console.log("initialized");
          },
        },
        // Set the save param.
        saveParam: "content",

        // Set the save URL.
        saveURL: "http://localhost:8080/editor/save",

        // HTTP request type.
        saveMethod: "POST",
      },
      model: "Edit Your Content Here!",
    };
  },
};
</script>