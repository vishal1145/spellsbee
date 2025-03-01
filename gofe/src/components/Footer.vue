<template>
  <footer class="footer">
    <div class="footer-content">
      <div>
        <div class="copyright">© 2025 SpellsBee.com</div>
        <div class="disclaimer" style="margin-bottom: 1rem;">
           All content and features are intended to support learning and practice.
        </div>
        <div class="links">
          <router-link to="/terms">Terms</router-link>
          <span class="separator">·</span>
          <router-link to="/privacy">Privacy</router-link>
          <span class="separator">·</span>
          <router-link to="/disclaimer">Disclaimer</router-link>
        </div>
      </div>
    </div>
  </footer>
</template>

<style scoped>
.footer {
  background-color: #f5f5f7;
  padding: 1.5rem;
  color: #6b7280;
  border-top: 1px solid #e5e7eb;
}

.footer-content {
  max-width: 1000px;
  margin: 0 auto;
  text-align: center;
}
 

.copyright {
  margin-bottom: 1rem;
  font-weight: 600;
  color: #374151;
}

.links {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 0.5rem;
}

.links a {
  color: #2563eb;
  text-decoration: none;
  transition: all 0.2s ease;
  font-weight: 400;
  padding: 0.25rem 0.5rem;
}

.links a:hover {
  color: #1d4ed8;
  background-color: transparent;
  text-decoration: underline;
}

.separator {
  color: #9ca3af;
  font-weight: bold;
}

.disclaimer {
  font-size: 0.875rem;
  color: #9ca3af;
  max-width: 500px;
  margin: 0 auto;
  line-height: 1.5;
}

@media (max-width: 640px) {
  .footer {
    padding: 1rem;
  }
  
  .links {
    flex-direction: column;
    gap: 0.75rem;
  }
  
  .separator {
    display: none;
  }
}
</style>
