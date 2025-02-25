<template>
  <div id="app" :class="{ 'dark-mode': isDarkMode }">
    <!-- Navigation icons -->
    <div class="nav-icons hidden lg:block">
      <div class="dark-mode-toggle">
        <button @click="toggleDarkMode">
          <i class="fas" :class="isDarkMode ? 'fa-sun' : 'fa-moon'"></i>
        </button>
      </div>
    </div>
    
    <router-view></router-view>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Cookies from 'js-cookie'
import axios from 'axios'

const isDarkMode = ref(false)
const userName = ref(null)

const toggleDarkMode = () => {
  isDarkMode.value = !isDarkMode.value
  localStorage.setItem('darkMode', isDarkMode.value)
}

const generateUniqueName = () => {
  const timestamp = Date.now()
  const email = `user_${timestamp}@example.com`
  const password = `pass_${timestamp}`
  return { username: `user_${timestamp}`, email, password }
}

const checkAndSetUserName = async () => {
  let username = Cookies.get('username')
  
  if (!username) {
    try {
      const credentials = generateUniqueName()
 
      const response = await axios.post(`${import.meta.env.VITE_API_URL}/api/users`, {
        username: credentials.username,
        email: credentials.email,
        password: credentials.password
      })
      
      if (response.status == 201) {
        Cookies.set('username', credentials.username, { expires: 365 }) 
        console.log('New user registered:', credentials.username)
      } else {
        console.error('Failed to register user')
      }
    } catch (error) {
      console.error('Error registering new user:', error)
    }
  }
  
  return username
}

onMounted(async () => {
  const username = await checkAndSetUserName()

  const savedDarkMode = localStorage.getItem('darkMode')
  if (savedDarkMode !== null) {
    isDarkMode.value = savedDarkMode === 'true'
  }
})
</script>

<style>
#app {
  font-family: Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  min-height: 100vh;
  transition: background-color 0.3s, color 0.3s;
}


/* Dark mode styles */
.dark-mode {
  background-color: #000000;
  color: #ffffff;
}

/* Add style for the spelling bee text */
#are-you-a-fan-of-spelling-bee {
  color: inherit;
}

.dark-mode #are-you-a-fan-of-spelling-bee {
  color: #000000;
}

/* Add style for start playing now text */
.start-playing-now {
  color: inherit;
}

.dark-mode .start-playing-now {
  color: #000000;
}

.nav-icons {
  width: 100px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: fixed;
  top: 1rem;
  left: 1rem;
  right: 1rem;
  z-index: 1000;
}

.dark-mode-toggle {
  font-size: 15px;
  display: flex;
  align-items: center;
}

.dark-mode-toggle button {
  background: none;
  border: none;
  cursor: pointer;
  padding: 8px;
  border-radius: 50%;
  color: #666;
  font-size: 1.2rem;
}

.dark-mode .dark-mode-toggle button {
  color: #fff;
}

/* Add this to your index.html or install via npm */
@import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css');
</style>
