<template>
  <div class="create-note">
    <h1>Create a Note</h1>
    <form @submit.prevent="createNote">
      <div>
        <label for="content">Content:</label>
        <textarea v-model="content" required></textarea>
      </div>
      <div>
        <label for="expiration">Expiration (YYYY-MM-DDTHH:MM:SSZ):</label>
        <input type="datetime-local" v-model="expiration" required />
      </div>
      <div>
        <label for="maxViews">Max Views:</label>
        <input type="number" v-model="maxViews" required min="1" />
      </div>
      <button type="submit">Create Note</button>
    </form>

    <div v-if="noteURL">
      <p>
        Note URL: <a :href="noteURL">{{ noteURL }}</a>
      </p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  setup() {
    const content = ref('')
    const expiration = ref('')
    const maxViews = ref(1)
    const noteURL = ref<string | null>(null)

    const createNote = async () => {
      try {
        const token = sessionStorage.getItem('token')

        let headers = new Headers({
          'Content-Type': 'application/json'
        })
        if (token) {
          console.log('Bearer ' + token)
          headers.append('Authorization', 'Bearer ' + token)
        }

        const response = await fetch(`${import.meta.env.VITE_BASE_URL}/note`, {
          method: 'POST',
          headers: headers,
          body: JSON.stringify({
            content: content.value,
            expiration: new Date(expiration.value).toISOString(),
            maxViews: maxViews.value
          })
        })

        const data = await response.json()
        if (response.ok) {
          noteURL.value = `${window.location.origin}${data.url}`
        } else {
          alert(data.error)
        }
      } catch (error) {
        console.error('Error: ', error)
      }
    }

    return {
      content,
      expiration,
      maxViews,
      noteURL,
      createNote
    }
  }
})
</script>
