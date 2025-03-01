<template>
  <div class="fixed inset-0 bg-black/50 flex justify-center items-center z-[1000]" @click="closePopup">
    <div class="bg-white dark:bg-gray-800 rounded-2xl w-[90%] max-w-[480px] max-h-[90vh] overflow-y-auto p-6" @click.stop>
      <div class="flex justify-between items-center">
        <button class="bg-transparent border-0 text-lg cursor-pointer text-gray-600 dark:text-gray-400" @click="closePopup">
          <i class="fas fa-arrow-left"></i>
        </button>
      </div>

      <div class="rank-stats-content">
        <div class="stats-tab">
          <div class="text-center mb-5">
            <span>D E K O R W Y</span>
            <div class="text-sm text-gray-600 dark:text-gray-400 mt-2">0 of 27 words | 0 of 117 points</div>
          </div>

          <div class="grid grid-cols-[repeat(auto-fill,minmax(100px,1fr))] gap-3 p-4">
            <div v-for="word in words" 
                 :key="word" 
                 class="p-2 bg-gray-100 dark:bg-gray-700 rounded text-center text-sm text-gray-600 dark:text-gray-400"
                 :class="{ 'bg-green-600 dark:bg-green-600 text-white dark:text-white': foundWords.includes(word) }">
              {{ word }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  isVisible: {
    type: Boolean,
    required: true
  },
  words: {
    type: Array,
    default: () => [
      'Dewy', 'Doddery', 'Dorky', 'Dory', 'Dowdy',
      'Dowry', 'Dryer', 'Dyed', 'Dyer', 'Eddy',
      'Eyed', 'Keyed', 'Keyword', 'Kooky',
      'Redye', 'Redyed', 'Reedy', 'Rookery',
      'Rowdy', 'Weedy', 'Woody', 'Wordy',
      'Worry', 'Wryer', 'Yoke', 'Yoked', 'Yore'
    ]
  },
  foundWords: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['close'])

const closePopup = () => {
  emit('close')
}
</script> 