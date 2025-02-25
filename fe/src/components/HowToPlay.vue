<template>
  <div class="how-to-play">
    <h1 class="title animate-fade-in text-[28px] font-semibold">Guide to Playing Spelling Bee</h1>
    <div class="instructions-container">
      <div v-for="(instruction, index) in instructions" 
           :key="index"
           class="instruction-block"
           :class="{ 'active': hoveredIndex === index }"
           @mouseenter="hoveredIndex = index"
           @mouseleave="hoveredIndex = null"
           :style="{ animationDelay: `${index * 0.2}s` }">
        <div class="number">
          <span class="number-text">{{index + 1}}</span>
          <svg class="hexagon" viewBox="0 0 24 24">
            <path d="M12 2l10 6v8l-10 6-10-6V8z" />
          </svg>
        </div>
        <div class="text-content">
          <h3 class="instruction-title">{{ instruction.title }}</h3>
          <p class="content">{{ instruction.content }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HowToPlay',
  data() {
    return {
      hoveredIndex: null,
      instructions: [
        {
          title: 'Create Words',
          content: 'Form words using seven letters, ensuring the center letter appears in each word. Each word must be at least four letters long, and letters can be repeated. A new puzzle is available every day!'
        },
        {
          title: 'Score Points Rule',
          content: 'Words with four letters earn 1 point, while longer words score 1 point per letter. Discover the Pangram, which uses all the letters, to receive a 7-point bonus!'
        },
        {
          title: 'Track Progress',
          content: 'Every puzzle has a highest possible score. Monitor your progress using the rating system displayed at the top of the game.'
        },
        {
          title: 'Game Rules',
          content: 'Hyphens, proper nouns, and offensive words are not allowed. Press the shuffle button to rearrange the letters and spark new ideas!'
        }
      ]
    }
  }
}
</script>

<style scoped>
.how-to-play {
  max-width: 1000px;
  margin: 0 auto;
  padding: 0 20px;
  margin-bottom: 64px;
  font-family: 'Lato', sans-serif;
}

.title {
  font-size: 2rem;
  font-weight: 600;
  color: #1a1a1a;
  text-align: center;
  margin-bottom: 3rem;
  background: linear-gradient(45deg, #333333, #000000);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  font-family: 'Lato', sans-serif;
}

.instructions-container {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.instruction-block {
  position: relative;
  padding: 1.5rem 2rem;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  animation: slideIn 0.8s cubic-bezier(0.4, 0, 0.2, 1) forwards;
  opacity: 0;
  border: 2px solid transparent;
  transform-origin: center;
  background-color: rgb(245, 245, 245);
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.instruction-block.active {
  border-color: #FFE999;
  transform: translateY(-3px) scale(1.01);
  box-shadow: 0 8px 30px rgba(255, 215, 0, 0.15);
}

.number {
  position: relative;
  width: 40px;
  height: 40px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  animation: rotateIn 0.6s cubic-bezier(0.4, 0, 0.2, 1) forwards;
}

.number-text {
  position: relative;
  z-index: 2;
  color: #1a1a1a;
  font-weight: 600;
  font-size: 1.2rem;
  animation: popIn 0.5s cubic-bezier(0.4, 0, 0.2, 1) forwards;
  animation-delay: 0.3s;
  opacity: 0;
  font-family: 'Lato', sans-serif;
}

.text-content {
  flex: 1;
}

.instruction-title {
  font-size: 16px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 0.5rem 0;
  animation: slideInRight 0.6s cubic-bezier(0.4, 0, 0.2, 1) forwards;
  animation-delay: 0.2s;
  opacity: 0;
  font-family: 'Lato', sans-serif;
}

.hexagon {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 100%;
  height: 100%;
  fill: #FFE999;
  stroke: #1a1a1a;
  stroke-width: 1;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.active .hexagon {
  fill: #e6c940;
  transform: translate(-50%, -50%) scale(1.1) rotate(180deg);
}

.content {
  font-size: 14px;
  line-height: 1.6;
  color: #4a4a4a;
  margin: 0;
  animation: slideInRight 0.6s cubic-bezier(0.4, 0, 0.2, 1) forwards;
  animation-delay: 0.4s;
  opacity: 0;
  font-family: 'Lato', sans-serif;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes rotateIn {
  from {
    transform: rotate(-180deg) scale(0.5);
    opacity: 0;
  }
  to {
    transform: rotate(0) scale(1);
    opacity: 1;
  }
}

@keyframes popIn {
  from {
    transform: scale(0.8);
    opacity: 0;
  }
  to {
    transform: scale(1);
    opacity: 1;
  }
}

@keyframes slideInRight {
  from {
    transform: translateX(-15px);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (max-width: 768px) {
  .how-to-play {
    padding: 20px 16px;
  }
  
  .title {
    font-size: 2rem;
    margin-bottom: 2rem;
  }
  
  .instruction-block {
    padding: 1.25rem;
    gap: 1rem;
  }
  
  .text-content {
    flex: 1;
  }
  
  .number {
    width: 32px;
    height: 32px;
  }
  
  .instruction-title {
    font-size: 1.1rem;
  }
  
  .content {
    font-size: 1rem;
    line-height: 1.5;
  }
}
</style>
