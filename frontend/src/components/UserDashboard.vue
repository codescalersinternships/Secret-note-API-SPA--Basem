<template>
  <div class="user-dashboard">
    <h1>Dashboard</h1>
    <nav>
      <ul>
        <li><router-link to="/note"> Create Note </router-link></li>
        <li><router-link to="/notes"> View My Notes </router-link></li>
      </ul>
    </nav>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'UserDashboard'
})
</script>
