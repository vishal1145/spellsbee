<template>
  <div class="popup-overlay" v-if="isVisible" @click.self="close">
    <div class="popup-content">
      <h2>Reset Password</h2>
      <div v-if="statusMessage" :class="['status-message', statusType]">
        {{ statusMessage }}
      </div>
      <form @submit.prevent="handleForgotPassword" v-if="!isSuccess">
        <div class="form-group">
          <label>Email</label>
          <input type="email" v-model="email" required :disabled="isLoading" @keyup="removeSpaces">
        </div>
        <div class="form-actions">
          <div class="links">
            <a href="#" @click.prevent="$emit('showLogin')">Back to Login</a>
          </div>
          <button type="submit" class="forgot-password-button" :disabled="isLoading">
            {{ isLoading ? 'Sending...' : 'Send Reset Link →' }}
          </button>
        </div>
      </form>
      <div v-else class="success-content">
        <p>Please check your email to reset your password.</p>
        <p class="email-sent">We've sent instructions to:<br><strong>{{ email }}</strong></p>
        <button class="forgot-password-button" @click="close">
          Close
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  isVisible: Boolean
})

const emit = defineEmits(['close', 'showLogin'])

const email = ref('')
const statusMessage = ref('')
const statusType = ref('')
const isLoading = ref(false)
const isSuccess = ref(false)

const close = () => {
  emit('close')
}

const handleForgotPassword = async () => {
  isLoading.value = true
  statusMessage.value = ''
  
  try {
    const response = await fetch(`${import.meta.env.VITE_API_URL}/api/users/forgot-password`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email: email.value }),
    });

    if (!response.ok) {
      throw new Error('Password reset request failed');
    }

    isSuccess.value = true;
    statusType.value = 'success'
  } catch (error) {
      statusMessage.value = 'Email not found please register first'
    console.error('Error:', error);
    statusType.value = 'error'
  } finally {
    isLoading.value = false
  }
}

const removeSpaces = () => {
  email.value = email.value.replace(/\s/g, '')
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

.forgot-password-button {
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

.forgot-password-button:hover {
  background-color: #388E3C;
  transform: translateY(-1px);
}

.forgot-password-button:active {
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

.email-sent {
  background-color: #f3f4f6;
  padding: 1rem;
  border-radius: 6px;
  margin: 1.5rem 0;
}

.email-sent strong {
  color: #1f2937;
}

button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
</style> 