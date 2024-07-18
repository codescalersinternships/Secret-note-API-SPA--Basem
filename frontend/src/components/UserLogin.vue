<template>
  <div class="login">
    <h1>Login</h1>
    <form @submit.prevent="login">
      <div>
        <label for="username">Username:</label>
        <input type="text" id="username" v-model="username" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" required />
      </div>
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useRouter } from 'vue-router'

export default defineComponent({
  setup() {
    const username = ref('')
    const password = ref('')
    const router = useRouter()

    const login = async () => {
      try {
        const response = await fetch(`${import.meta.env.VITE_BASE_URL}/login`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ username: username.value, password: password.value })
        })

        const data = await response.json()
        if (response.ok) {
          sessionStorage.setItem('token', data.token)
          router.push('/dashboard')
        } else {
          alert(data.error)
        }
      } catch (error) {
        console.error('Error: ', error)
      }
    }
    return {
      username,
      password,
      login
    }
  }
})
</script>
