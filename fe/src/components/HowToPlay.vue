<template>
  <div class="how-to-play">
    <div class="instructions-container">
       <div class="mt-[64px]">
        <h1 class="text-[32px] font-semibold text-[#8B4513] font-lato text-center mb-4">
            Spelling Bee Game
        </h1>
    </div>
    
      <p class="description mb-8">
        Spelling Bee is a widely loved word game in the US with simple rules. Your goal is to create as many words as possible using a set of seven letters. Each day, you receive a new set consisting of six regular letters and one required letter. To form words of at least four letters, you can tap on the letters on the screen or use your keyboard. You may use any letter multiple times, but every word must include the central letter. The more words you find, the higher your score.
      </p>
      
      <h2 class="subtitle mb-4">Game Rules</h2>
      <ul class="rules-list">
        <li v-for="(instruction, index) in instructions" 
            :key="index"
            class="rule-item">
          <h3 class="rule-title">{{ instruction.title }}</h3>
          <p class="rule-content">{{ instruction.content }}</p>
        </li>
      </ul>
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

.description {
  font-size: 14px; 
  line-height: 1.6;
  color: #4a4a4a;
}

.subtitle {
  font-size: 16px;
  font-weight: 600;
  color: #1a1a1a;
}

.rules-list {
  list-style-type: disc;
  padding-left: 20px;
}

.rule-item {
  margin-bottom: 24px;
}

.rule-title {
  font-size: 14px;
  font-weight: 600;
  color: #1a1a1a;
  margin-bottom: 8px;
}

.rule-content {
  font-size: 14px;
  line-height: 1.6;
  color: #4a4a4a;
}

@media (max-width: 768px) {
  .how-to-play {
    padding: 20px 16px;
  }
  
  .description {
    font-size: 14px;
  }
  
  .subtitle {
    font-size: 20px;
  }
}
</style>
