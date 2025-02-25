<template>
  <div class="popup-overlay" v-if="isVisible" @click.self="close">
    <div class="popup-content">
      <h2>Reset Password</h2>
      <form @submit.prevent="handleForgotPassword">
        <div class="form-group">
          <label>Email</label>
          <input type="email" v-model="email" required>
        </div>
        <div class="form-actions">
          <div class="links">
            <a href="#" @click.prevent="$emit('showLogin')">Back to Login</a>
          </div>
          <button type="submit" class="forgot-password-button">Send Reset Link →</button>
        </div>
      </form>
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

const close = () => {
  emit('close')
}

const handleForgotPassword = async () => {
  // Implement password reset logic here
  console.log('Password reset attempt:', { email: email.value })
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
</style> 