<template>
  <div class="card">
    <h1>Vue + Fiber</h1>
    <p>{{ message }}</p>
    <button @click="fetchData">Fetch from Go API</button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
const message = ref("Klik tombol untuk ambil data")

const fetchData = async () => {
  const res = await fetch("/api/hello")
  const data = await res.json()
  message.value = data.message
}
</script>

<style scoped>
.card {
  margin: 3rem auto;
  padding: 2rem;
  max-width: 600px;
  text-align: center;
  background: #42b883;
  color: white;
  border-radius: 10px;
  box-shadow: 0 8px 20px rgba(0,0,0,0.2);
}
button {
  margin-top: 1rem;
  padding: 0.6rem 1.2rem;
  background: #35495e;
  border: none;
  color: white;
  border-radius: 8px;
  cursor: pointer;
}
button:hover {
  background: #2c3e50;
}
</style>
