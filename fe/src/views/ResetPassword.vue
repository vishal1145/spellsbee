<template>
  <div class="reset-password-container">
    <div class="reset-form">
      <h2>Set New Password</h2>
      <div v-if="statusMessage" :class="['status-message', statusType]">
        {{ statusMessage }}
      </div>
      <form @submit.prevent="handleResetPassword" v-if="!isSuccess">
        <div class="form-group">
          <label>New Password</label>
          <input 
            type="password" 
            v-model="password" 
            required 
            :disabled="isLoading"
            placeholder="Enter new password"
          >
        </div>
        <div class="form-group">
          <label>Confirm Password</label>
          <input 
            type="password" 
            v-model="confirmPassword" 
            required 
            :disabled="isLoading"
            placeholder="Confirm new password"
          >
        </div>
        <div class="form-actions">
          <button type="submit" class="reset-button" :disabled="isLoading">
            {{ isLoading ? 'Resetting...' : 'Reset Password →' }}
          </button>
        </div>
      </form>
      <div v-else class="success-content">
        <p>Your password has been successfully reset!</p>
        <button class="reset-button" @click="redirectToLogin">
          Go to Login
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const password = ref('')
const confirmPassword = ref('')
const statusMessage = ref('')
const statusType = ref('')
const isLoading = ref(false)
const isSuccess = ref(false)

const handleResetPassword = async () => {
  // Clear previous status
  statusMessage.value = ''
  
  // Get token from URL query parameters
  const token = new URLSearchParams(window.location.search).get('token')
  
  // Validate token
  if (!token) {
    statusMessage.value = 'Invalid or missing reset token. Please check your reset password link.'
    statusType.value = 'error'
    return
  }

  // Validate passwords
  if (password.value !== confirmPassword.value) {
    statusMessage.value = 'Passwords do not match'
    statusType.value = 'error'
    return
  }

  isLoading.value = true
  
  try {
    const response = await fetch(`${import.meta.env.VITE_API_URL}/api/users/reset-password`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ 
        token,
        password: password.value 
      }),
    })

    if (!response.ok) {
      const data = await response.json()
      throw new Error(data.message || 'Password reset failed')
    }

    isSuccess.value = true
    statusType.value = 'success'
  } catch (error) {
    statusMessage.value = error.message || 'Failed to reset password. Please try again.'
    statusType.value = 'error'
    console.error('Error:', error)
  } finally {
    isLoading.value = false
  }
}

const redirectToLogin = () => {
  router.push('/')
}
</script>

<style scoped>
.reset-password-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f9fafb;
}

.reset-form {
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
  display: block;
  margin-bottom: 0.5rem;
  color: #4a5568;
  font-size: 14px;
  font-weight: 500;
}

input {
  width: 100%;
  padding: 0.875rem;
  border: 1.5px solid #e2e8f0;
  border-radius: 8px;
  font-size: 14px;
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

.reset-button {
  width: 100%;
  padding: 0.875rem;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.reset-button:hover {
  background-color: #388E3C;
  transform: translateY(-1px);
}

.reset-button:active {
  transform: translateY(0);
}

.status-message {
  padding: 0.75rem;
  border-radius: 6px;
  margin-bottom: 1rem;
  font-size: 14px;
  text-align: center;
}

.status-message.error {
  background-color: #fee2e2;
  color: #dc2626;
  border: 1px solid #fecaca;
}

.status-message.success {
  background-color: #dcfce7;
  color: #16a34a;
  border: 1px solid #bbf7d0;
}

.success-content {
  text-align: center;
}

.success-content p {
  margin-bottom: 1rem;
  color: #4a5568;
  font-size: 14px;
  line-height: 1.5;
}

button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
</style> 