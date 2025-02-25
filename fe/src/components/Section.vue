<template>
    <div>
    <div class="bg-[#FFE999] py-[32px] px-6 text-center my-[64px] relative rounded-2xl shadow-lg">
        <img 
            src="../assets/bee.svg" 
            alt="Honey Bee" 
            class="absolute left-4 top-1/2 transform -translate-y-1/2 w-[100px] hover:scale-110 transition-transform animate-bounce-slow"
        />
        <img 
            src="../assets/bee.svg" 
            alt="Honey Bee" 
            class="absolute right-4 top-1/2 transform -translate-y-1/2 w-[100px] scale-x-[-1] hover:scale-110 transition-transform animate-bounce-slow"
        />
        <div class="relative">
            <img 
                src="../assets/bee.svg" 
                alt="Honeycomb" 
                class="absolute left-0 top-0 w-[150px] opacity-20"
            />
            <h1 class="text-[28px] font-semibold text-[#8B4513] mb-8 font-lato relative inline-block">
                Spelling Bee Game
            </h1>
        </div>
        <p class="max-w-3xl mx-auto text-[14px] text-[#654321] font-lato leading-relaxed font-normal">
            Spelling Bee is a widely loved word game in the US with simple rules. Your goal is to create as many words as possible using a set of seven letters. Each day, you receive a new set consisting of six regular letters and one required letter. To form words of at least four letters, you can tap on the letters on the screen or use your keyboard. You may use any letter multiple times, but every word must include the central letter. The more words you find, the higher your score. Check below for a detailed breakdown of the rules.
        </p>
    </div>
    </div>
</template>

<script>
 
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Lato:wght@400;700;900&display=swap');

@keyframes bounce-slow {
  0%, 100% {
    transform: translateY(-50%) translateY(-5px);
  }
  50% {
    transform: translateY(-50%) translateY(5px);
  }
}

.animate-bounce-slow {
  animation: bounce-slow 3s infinite ease-in-out;
}

.font-lato {
  font-family: 'Lato', sans-serif;
}
</style>