<template>
  <div class="view-note">
    <h1>View Note</h1>
    <div v-if="content">
      <p>{{ content }}</p>
    </div>

    <div v-else style="color: red">
      <p>Note not found or has expired!</p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'

export default defineComponent({
  setup() {
    const content = ref<string | null>(null)
    const route = useRoute()

    onMounted(async () => {
      const noteID = route.params.key as string
      try {
        const response = await fetch(`${import.meta.env.VITE_BASE_URL}/note/${noteID}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json'
          }
        })
        const data = await response.json()

        if (response.ok) {
          content.value = data.content
        } else {
          content.value = null
        }
      } catch (error) {
        console.error('Error', error)
      }
    })

    return {
      content
    }
  }
})
</script>
