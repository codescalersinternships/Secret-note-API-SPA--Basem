<template>
  <div class="home">
    <h1>Welcome to the Secret Note Sharing Application</h1>
    <router-link to="/login">Login</router-link> or
    <router-link to="/register">Register</router-link>
    <br />
    Or Simply
    <router-link to="/note">Create a Note!</router-link>
  </div>
</template>

<script lang="ts">
export default {
  name: 'HomePage'
}
</script>
