<template>
  <div class="col-sm-6 col-sm-offset-3">
    <h1>Get a public quote!</h1>
    <button class="btn btn-primary" v-on:click="getQuote()">Get a Quote</button>
    <div class="quote-area" v-if="quote">
      <h2><blockquote>{{ quote }}</blockquote></h2>
    </div>
  </div>
</template>

<script>
export default {
  name: "Home",
  data(){
    return{
      quote:''
    }
  },
  methods: {
    getQuote () {
      console.log("asd");
    }
      /*this.$http.get('/api/quote/random').then(response => {
        this.quote = response.body
      }, response => {
        console.log(response)
      })
    }*/
  }
}
</script>

<style scoped>

</style>