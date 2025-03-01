<template>
  <div class="max-w-[1000px] mx-auto px-4   mb-[64px] font-lato">
    <!-- Hero Section -->

        <!-- Bottom CTA Section -->
    <div class="  text-center bg-gradient-to-r from-yellow-50 to-orange-50 rounded-3xl p-8 md:p-12">
      <h2 class="text-[28px] font-black font-semibold mb-4 " id="are-you-a-fan-of-spelling-bee">Are you a fan of Spelling Bee?</h2>
      <p class="text-[14px] text-gray-600 max-w-2xl mx-auto mb-8">
       Test your skills with our daily word puzzle! Form words from seven letters, score extra points for pangrams, and challenge players from around the world. Ideal for both word lovers and casual gamers.
      </p>
      <button id="start-playing-now"   class="bg-[#FFE999] hover:bg-[#FFD666] text-[14px] font-semibold py-3 px-8 rounded-full 
                     transition-colors duration-300 shadow-lg hover:shadow-xl">
        Start Playing Now
      </button>
    </div>

    <div class="mt-[64px]  text-center mb-4">
      <h1 class="text-[28px] font-semibold mb-4">
        Frequently Asked Questions
      </h1>
      <div class="w-20 h-1 bg-yellow-400 mx-auto rounded-full"></div>
    </div>
    
    <!-- FAQ Accordion -->
    <div class="space-y-4">
      <div v-for="(faq, index) in faqs" 
           :key="index"
           class="border border-gray-100 rounded-2xl overflow-hidden shadow-sm hover:shadow-md transition-shadow duration-300"
           :class="{ 'bg-white': activeIndex === index, 'bg-gray-50': activeIndex !== index }">
        <button 
          @click="toggleFaq(index)"
          class="w-full px-6 py-5 flex items-center justify-between text-left"
        >
          <h2 class="text-[16px] font-semibold text-gray-900">{{ faq.question }}</h2>
          <span class="transform transition-transform duration-300"
                :class="{ 'rotate-180': activeIndex === index }">
            <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
          </span>
        </button>
        
        <div 
          v-show="activeIndex === index"
          class="px-6 pb-5"
        >
          <p class="text-gray-600 leading-relaxed text-[14px]" v-html="faq.answer"></p>
        </div>
      </div>
    </div>


  </div>
</template>

<script setup>
import { ref } from 'vue'

const activeIndex = ref(null)

const faqs = [
  
  {
    question: "Why is the word not in the word list?",
    answer: "Our dictionary isn't exhaustive, so if you encounter this error, it means the word you entered isn't recognized. As a result, it won't be accepted in the game, and you won't score any points for it. If you believe this word should be included, feel free to email us at spellingbeemail@gmail.com."
  },
  {
    question: "Where can I see the answers to the game?",
    answer: "You can check the answers to today's game only after midnight the following day. Simply click on the \"Yesterday\" button to view the full list of words from the previous game, with your submitted answers highlighted in bold."
  },
  {
    question: "At what time does the daily puzzle change?",
    answer: "Every day, a new set of letters is available to players. New letters appear every day at 12:00 am, according to the time on the device you are playing on."
  }
]

const toggleFaq = (index) => {
  activeIndex.value = activeIndex.value === index ? null : index
}
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
