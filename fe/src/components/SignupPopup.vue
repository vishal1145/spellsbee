<template>
  <div class="popup-overlay" v-if="isVisible" @click.self="close">
    <div class="popup-content">
      <template v-if="!signupSuccess">
        <h2>Create your account</h2>
        <form @submit.prevent="handleSignup">
          <div class="form-group">
            <label>Username</label>
            <input 
              type="text" 
              v-model="username" 
              required
              minlength="4"
              maxlength="50"
              @input="validateUsername"
              @keyup="removeSpaces('username')">
            <span v-if="usernameError" class="error-message">{{ usernameError }}</span>
          </div>
          <div class="form-group">
            <label>Email</label>
            <input 
              type="email" 
              v-model="email" 
              required 
              maxlength="50"
              @keyup="removeSpaces('email')">
            <span v-if="emailError" class="error-message">{{ emailError }}</span>
          </div>
          <div class="form-group">
            <label>Password</label>
            <div class="password-input-container">
              <input 
                :type="showPassword ? 'text' : 'password'" 
                v-model="password" 
                required
                minlength="4"
                maxlength="25"
                @input="validatePassword"
                @keyup="removeSpaces('password')">
              <button type="button" class="toggle-password" @click="togglePassword">
                <i :class="showPassword ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
              </button>
            </div>
            <span v-if="passwordError" class="error-message">{{ passwordError }}</span>
          </div>
          <div class="form-actions">
            <div class="links">
              <a href="#" @click.prevent="$emit('showLogin')">Already have an account? Log in</a>
            </div>
            <button type="submit" class="signup-button">Create Account →</button>
          </div>
        </form>
      </template>
      <div v-else class="success-message">
        <div class="success-icon">✓</div>
        <h2>Account Created Successfully!</h2>
        <p>Please check your email to verify your account. We've sent you a verification link.</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'
import Cookies from 'js-cookie'

const props = defineProps({
  isVisible: Boolean
})

const emit = defineEmits(['close', 'showLogin'])

const username = ref('')
const email = ref('')
const password = ref('')
const signupSuccess = ref(false)
const showPassword = ref(false)
const usernameError = ref('')
const passwordError = ref('')
const emailError = ref('')

const close = () => {
  emit('close')
}

const validateUsername = () => {
  if (username.value.length < 4) {
    usernameError.value = 'Username must be at least 4 characters'
  } else if (username.value.length > 50) {
    usernameError.value = 'Username must not exceed 50 characters'
  } else {
    usernameError.value = ''
  }
}

const validatePassword = () => {
  if (password.value.length < 4) {
    passwordError.value = 'Password must be at least 4 characters long'
  } else if (password.value.length > 25) {
    passwordError.value = 'Password must not exceed 25 characters'
  } else {
    passwordError.value = ''
  }
}

const removeSpaces = (field) => {
  if (field === 'username') {
    username.value = username.value.replace(/\s/g, '')
  } else if (field === 'email') {
    email.value = email.value.replace(/\s/g, '')
  } else if (field === 'password') {
    password.value = password.value.replace(/\s/g, '')
  }
}

const handleSignup = async () => {
  // Validate before submitting
  validateUsername()
  validatePassword()
  
  if (usernameError.value || passwordError.value) {
    return
  }
  
  try {
    const response = await fetch(`${import.meta.env.VITE_API_URL}/api/users/register`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: Cookies.get('username'),
        newusername: username.value,
        email: email.value,
        password: password.value
      })
    })
    
    const data = await response.json()
    
    if (!response.ok) {
      // Reset both error messages first
      emailError.value = ''
      usernameError.value = ''
      
      // Set specific error based on the message
      if (data.message === 'Email already exists') {
        emailError.value = 'This email is already registered. Please use a different email.'
      }
      if (data.message === 'Username already exists') {
        usernameError.value = 'This username is already taken. Please choose a different username.'
      }
      return
    }

    Cookies.set('username', username.value, { expires: 365 });

    username.value = ''
    email.value = ''
    password.value = ''
    signupSuccess.value = true

  } catch (error) {
    console.error('Signup failed:', error.response?.data || error.message)
  }
}

const togglePassword = () => {
  showPassword.value = !showPassword.value
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
  font-size: 15px;
  text-align: center;
  margin-bottom: 2rem;
  color: #2c3e50;
  font-weight: 600;
}

.form-group {
  font-size: 14px;
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
  transition: all 0.2s;
}

input:focus {
  outline: none;
  border-color: #4CAF50;
  box-shadow: 0 0 0 3px rgba(76, 175, 80, 0.1);
}

.form-actions {
  margin-top: 2rem;
}

.links {
  display: flex;
  justify-content: center;
  margin-bottom: 1.25rem;
}

.links a {
  color: #4CAF50;
  text-decoration: none;
  font-size: 14px;
  font-weight: 500;
  transition: color 0.2s;
}

.links a:hover {
  color: #388E3C;
}

.signup-button {
  width: 100%;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.signup-button:hover {
  background-color: #388E3C;
  transform: translateY(-1px);
}

.signup-button:active {
  transform: translateY(0);
}

.success-message {
  text-align: center;
  padding: 1rem;
}

.success-icon {
  background-color: #4CAF50;
  color: white;
  width: 60px;
  height: 60px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32px;
  margin: 0 auto 1.5rem;
}

.success-message h2 {
  color: #4CAF50;
  margin-bottom: 1rem;
}

.success-message p {
  color: #4a5568;
  font-size: 14px;
  line-height: 1.5;
}

.password-input-container {
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
  color: #4CAF50;
}

.error-message {
  color: #dc3545;
  font-size: 12px;
  margin-top: 4px;
  display: block;
}
</style> 