<template>
  <div class="col-sm-6 col-sm-offset-3">
    <h1>Get a secret quote!</h1>
    <button class="btn btn-warning" v-on:click="getQuote()">Get a Quote</button>
    <div class="quote-area" v-if="quote">
      <h2>
        <blockquote>{{ quote }}</blockquote>
      </h2>
    </div>
  </div>
</template>

<script>
import auth from '../auth'

export default {
  name: "SecretQuote",
  data () {
    return {
      quote: ''
    }
  },
  methods: {
    getQuote () {
      this.$http
          .get('/api/quote/protected/random', { headers: auth.getAuthHeader() })
          .then(
              response => {
                this.quote = response.body
              },
              response => {
                if (response.status === 401) {
                  auth.logout(this)
                }
                console.log(response)
              }
          )
    }
  }
}
</script>

<style scoped>

</style>