<template>
  <div class="my-notes">
    <h1>My Notes</h1>
    <ul>
      <li v-for="note in notes" :key="note.ID">
        <a :href="'/note/' + note.UniqueKey"
          >{{ note.ID }} - Current Views: {{ note.Views }} - Expiration : {{ note.Expiration }} -
          Max Views: {{ note.MaxViews }}
        </a>
        <br />
        {{ note.Content }}
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'

interface Note {
  ID: number
  Content: string
  UniqueKey: string
  Expiration: string
  CreatedAt: string
  MaxViews: number
  Views: number
}

export default defineComponent({
  setup() {
    const notes = ref<Note[]>([])

    onMounted(async () => {
      try {
        const token = sessionStorage.getItem('token')
        const response = await fetch(`${import.meta.env.VITE_BASE_URL}/user/notes`, {
          headers: {
            Authorization: `Bearer ${token}`,
            'Content-Type': 'application/json'
          }
        })

        const data = await response.json()
        if (response.ok) {
          console.log(data)
          notes.value = data.notes
        } else {
          alert(data.error)
        }
      } catch (error) {
        console.error('Error: ', error)
      }
    })

    return {
      notes
    }
  }
})
</script>
