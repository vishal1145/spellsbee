<template>
  <div class="verify-email">
    <h1>Email Verification</h1>
    <div v-if="loading">Verifying your email...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <div v-else class="success">Your email has been verified successfully!</div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'VerifyEmail',
  data() {
    return {
      loading: true,
      error: null,
    };
  },
  async created() {
    try {
      const token = this.$route.query.token;
      if (!token) {
        this.error = 'Verification token is missing';
        this.loading = false;
        return;
      }

      const apiUrl = import.meta.env.VITE_API_URL || process.env.VUE_APP_API_URL;
      if (!apiUrl) {
        this.error = 'API URL is not defined';
        this.loading = false;
        return;
      }

      console.log('API URL:', apiUrl);
      console.log('Verification Token:', token);

      const response = await axios.put(`${apiUrl}/api/users/verify-email`, { token });

      if (response.data) {
        this.loading = false;
        setTimeout(() => {
          this.$router.push('/');
        }, 1500);
      } else {
        this.error = response.data.message || 'Failed to verify email';
        this.loading = false;
      }
      
    } catch (err) {
      console.error('Error verifying email:', err);
      this.error = err.response?.data?.message || 'Failed to verify email. Please try again later.';
      this.loading = false;
    }
  },
};
</script>

<style scoped>
.verify-email {
  max-width: 600px;
  margin: 2rem auto;
  padding: 1rem;
  text-align: center;
}

.error {
  color: red;
  margin-top: 1rem;
}

.success {
  color: green;
  margin-top: 1rem;
}
</style>
