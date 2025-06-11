<template>
  <div class="card">
    <h1>{{ title }}</h1>
    <p>{{ message }}</p>
    <button @click="fetchHello">Ambil Data dari Backend</button>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const title = '🎉 Vue + Go berjalan!'
const message = ref('Klik tombol untuk ambil pesan dari Go API.')

async function fetchHello() {
  const res = await fetch('/api/hello')
  const data = await res.json()
  message.value = data.message
}
</script>

<style scoped>
.card {
  max-width: 600px;
  margin: 2rem auto;
  padding: 2rem;
  border-radius: 12px;
  background: #42b883;
  color: white;
  text-align: center;
  box-shadow: 0 10px 25px rgba(0,0,0,0.2);
}
button {
  margin-top: 1rem;
  padding: 0.6rem 1.2rem;
  font-size: 1rem;
  border: none;
  border-radius: 8px;
  background: #35495e;
  color: white;
  cursor: pointer;
}
button:hover {
  background: #2c3e50;
}
</style>
