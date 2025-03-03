<template>
  <div class="popup-overlay" v-if="isVisible" @click.self="close">
    <div class="popup-content">
      <h2>Welcome Back</h2>
      <div v-if="errorMessage" class="error-message">
        {{ errorMessage }}
      </div>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label>Email</label>
          <input
            type="email"
            v-model="email"
            required
            minlength="4"
            maxlength="50"
            @keyup="removeSpaces('email')"
          />
        </div>
        <div class="form-group">
          <label>Password</label>
          <div class="password-input-wrapper">
            <input
              :type="showPassword ? 'text' : 'password'"
              v-model="password"
              required
              minlength="4"
              maxlength="25"
              @keyup="removeSpaces('password')"
            />
            <button type="button" class="toggle-password" @click="togglePassword">
              <i :class="showPassword ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
            </button>
          </div>
        </div>
        <div class="form-actions">
          <div class="links">
            <a href="#" @click.prevent="$emit('showForgotPassword')">Forgot Password?</a>
            <a href="#" @click.prevent="$emit('showSignup')">Sign Up</a>
          </div>
          <button type="submit" class="login-button">Login →</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import Cookies from 'js-cookie'
import axios from 'axios'

const props = defineProps({
  isVisible: Boolean,
})

const emit = defineEmits(['close', 'showSignup', 'showForgotPassword', 'loginSuccess'])

const email = ref('')
const password = ref('')
const showPassword = ref(false)
const errorMessage = ref('')

const close = () => {
  emit('close')
}

const handleLogin = async () => {
  try {
    errorMessage.value = '' // Clear any previous error

    // Add validation check
    if (email.value.length < 4 || email.value.length > 50) {
      errorMessage.value = 'Email must be between 4 and 50 characters.'
      return
    }

    if (password.value.length < 4 || password.value.length > 25) {
      errorMessage.value = 'Password must be between 4 and 25 characters.'
      return
    }

    const response = await axios.put(`${import.meta.env.VITE_API_URL}/api/users/login`, {
      email: email.value,
      password: password.value,
    })

    const data = response.data
    if (response.status === 200) {
      Cookies.set('spellsBeeUsername', data.username, { expires: 365 })
      localStorage.setItem('userLoginData', JSON.stringify(data))
      emit('loginSuccess', data)
      close()
      window.location.reload()
    } else {
      if (data.isEmailVerified === false) {
        errorMessage.value = 'Please verify your email account before login.'
      } else {
        errorMessage.value = data.message || 'Login failed. Please try again.'
      }
    }
  } catch (error) {
    console.error('Login error:', error)
    if (error.response) {
      // Handle specific API error responses
      const data = error.response.data
      if (data.isEmailVerified === false) {
        errorMessage.value = 'Please verify your email account before login.'
      } else {
        errorMessage.value = data.message || 'Login failed. Please try again.'
      }
    } else {
      errorMessage.value = 'An error occurred. Please try again later.'
    }
  }
}

const togglePassword = () => {
  showPassword.value = !showPassword.value
}

const removeSpaces = (field) => {
  if (field === 'email') {
    email.value = email.value.replace(/\s/g, '')
  } else if (field === 'password') {
    password.value = password.value.replace(/\s/g, '')
  }
}
</script>

<style scoped>
.popup-overlay {
  background: rgba(0, 0, 0, 0.6);
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.popup-content {
  background: white;
  padding: 2.5rem;
  border-radius: 12px;
  width: 100%;
  max-width: 400px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);
}

h2 {
  text-align: center;
  margin-bottom: 2rem;
  font-size: 15px;
  color: #2c3e50;
  font-weight: 600;
}

.form-group {
  margin-bottom: 1.25rem;
}

label {
  font-size: 14px;
  display: block;
  margin-bottom: 0.5rem;
  color: #4a5568;
  font-weight: 500;
}

input {
  font-size: 14px;
  width: 100%;
  border: 1.5px solid #e2e8f0;
  border-radius: 8px;
  font-size: 14px;
  transition: all 0.2s;
}

input:focus {
  outline: none;
  border-color: #4caf50;
  box-shadow: 0 0 0 3px rgba(76, 175, 80, 0.1);
}

.form-actions {
  margin-top: 2rem;
}

.links {
  display: flex;
  justify-content: space-between;
  margin-bottom: 1.25rem;
}

.links a {
  color: #4caf50;
  text-decoration: none;
  font-size: 14px;
  font-weight: 500;
  transition: color 0.2s;
}

.links a:hover {
  color: #388e3c;
}

.login-button {
  width: 100%;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.login-button:hover {
  background-color: #388e3c;
  transform: translateY(-1px);
}

.login-button:active {
  transform: translateY(0);
}

.password-input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.toggle-password {
  position: absolute;
  right: 10px;
  background: none;
  border: none;
  cursor: pointer;
  color: #4a5568;
  padding: 5px;
}

.toggle-password:hover {
  color: #4caf50;
}

.error-message {
  background-color: #fee2e2;
  color: #dc2626;
  padding: 0.75rem;
  border-radius: 6px;
  font-size: 14px;
  margin-bottom: 1rem;
  text-align: center;
}
</style>
