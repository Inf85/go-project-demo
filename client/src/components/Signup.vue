<template>
  <div class="col-sm-4 col-sm-offset-4">
    <h2>Sign Up</h2>
    <div class="alert alert-danger" v-if="error">
      <p>{{ error }}</p>
    </div>
    <div class="form-group">
      <input
          type="text"
          class="form-control"
          placeholder="Enter your username"
          v-model="credentials.username"
          v-on:keyup.enter="submit"
      >
    </div>
    <div class="form-group">
      <input
          type="password"
          class="form-control"
          placeholder="Enter your password"
          v-model="credentials.password"
          v-on:keyup.enter="submit"
      >
    </div>
    <button class="btn btn-primary" @click="submit">Signup</button>
  </div>
</template>

<script>
import auth from '../auth'

export default {
  name: "Signup",
  data () {
    return {
      credentials: {
        username: '',
        password: ''
      },
      error: ''
    }
  },
  methods: {
    submit () {
      var credentials = {
        username: this.credentials.username,
        password: this.credentials.password
      }
      // console.log('signup: credentials=', credentials)
      auth.signup(this, credentials, 'secretquote')
    }
  }
}
</script>

<style scoped>

</style>