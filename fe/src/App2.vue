<template>
  <div>
    <div class="game-container">
      <LoginPopup 
        :is-visible="activePopup === 'login'"
        @close="closePopup"
        @show-signup="activePopup = 'signup'"
        @show-forgot-password="activePopup = 'forgotPassword'"
      />
      <SignupPopup 
        :is-visible="activePopup === 'signup'"
        @close="closePopup"
        @show-login="activePopup = 'login'"
      />
      <ForgotPasswordPopup 
        :is-visible="activePopup === 'forgotPassword'"
        @close="closePopup"
        @show-login="activePopup = 'login'"
      />

      <!-- Top navigation bar -->
      <div class="nav-bar">
        <div class="left-section">
          <div class="help-section">
            <img src="./assets/bee2.svg" alt="Flower" class="flower-icon cursor-pointer" />
            <a
              class="text-[14px] font-[500] text-[#666] underline cursor-pointer flex items-center gap-2"
              @click="showHowToPlay = true"
              title="Learn how to play the game"
            >
              How to play
            </a>
          </div>
        </div>

        <div class="center-section">
          <div class="mode-toggles">
            <span
              class="mode-btn"
              :class="{ active: gameMode === 'today' }"
              @click="setGameMode('today')"
              >Now</span
            >
            <span
              class="mode-btn"
              :class="{ active: gameMode === 'unlimited' }"
              @click="setGameMode('unlimited')"
              >Full access</span
            >
          </div>
        </div>

        <div class="right-section">
          <button class="nav-icon-btn" title="Reset Game" @click="resetGame">
            <i class="fas fa-redo"></i>
          </button>
          <button class="nav-icon-btn" title="Contact Us">
            <i class="fas fa-envelope"></i>
          </button>

          <!-- User profile section -->
          <div v-if="userProfile" class="user-profile-section">
            <button class="profile-btn" @click="toggleProfileMenu" ref="profileBtn">
              <div class="avatar">{{ userProfile.username[0].toUpperCase() }}</div>
            </button>
            
            <!-- Profile dropdown menu -->
            <div v-if="showProfileMenu" class="profile-dropdown" ref="profileMenu">
              <div class="profile-header">
                <div class="avatar-large">{{ userProfile.username[0].toUpperCase() }}</div>
                <div class="user-info">
                  <div class="username">{{ userProfile.username }}</div>
                  <div class="email">{{ userProfile.email }}</div>
                </div>
              </div>
              <div class="dropdown-divider"></div>
              <button class="dropdown-item" @click="logout">
                <i class="fas fa-sign-out-alt"></i>
                Logout
              </button>
            </div>
          </div>

          <!-- Show login button if user is not logged in -->
          <button v-else class="nav-icon-btn ml-[8px] font-weight-[600] text-[14px]" title="Login">
            <a class="nav-link cursor-pointer" @click="showLoginPopup">Login</a>
          </button>
        </div>
      </div>

      <!-- Modified layout section -->
      <div class="game-layout">
        <div class="left-section2">
          <div class="right-content">
            <div class="level-info">
              <div class="level-header">
                <div class="level-title">Rank</div>
                <span class="subtext">( 0 points | Just 20 more to reach N level )</span>
              </div>

              <div class="progress-dots">
                <span 
                  v-for="(letter, index) in ['B', 'N', 'A', 'V', 'S', 'F', 'E', 'G']" 
                  :key="letter"
                  class="dot w-[20px] h-[20px]"
                  :style="{ backgroundColor: getRankBackground(letter, index) }"
                  :title="getRankTooltip(letter)"
                  @click="showRankPopup(letter)"
                >
                  {{ letter }}
                </span>
              </div>
            </div>

            <hr class="divider" />

            <div class="words-section">
              <div class="level-title mb-[8px]">Words</div>
              <div class="found-words-container">
                <div v-if="foundWords.length === 0" class="text-[14px] text-[#666]">
                  No correct words entered by you
                </div>
                <div v-else class="word-chips text-[14px]">
                  <span v-for="word in foundWords" 
                        :key="word" 
                        class="word-chip">
                    {{ word }}
                  </span>
                </div>
              </div>
            </div>
            <hr class="divider" />

            <!-- New tabbed section -->
            <div class="tabs-section">
              <div class="tabs-header">
                <button
                  class="tab-button"
                  :class="{ active: activeTab === 'stats' }"
                  @click="activeTab = 'stats'"
                >
                  Statistics
                </button>
                <button
                  class="tab-button"
                  :class="{ active: activeTab === 'games' }"
                  @click="activeTab = 'games'"
                  v-if="gameMode === 'unlimited'"
                >
                  Games
                </button>
              </div>

              <!-- Stats Tab Content -->
              <div v-if="activeTab === 'stats'" class="tab-content">
                <div class="stats-checkboxes font-weight-[600]" v-if="gameMode === 'today'">
                  <label class="checkbox-label">
                    <input
                      type="checkbox"
                      :checked="statsPeriod === 'thisWeek'"
                      @change="statsPeriod = 'thisWeek'"
                    />
                    Last 7 days
                  </label>
                  <label class="checkbox-label">
                    <input
                      type="checkbox"
                      :checked="statsPeriod === 'allTime'"
                      @change="statsPeriod = 'allTime'"
                    />
                    All Time
                  </label>
                </div>

                <div v-if="statsPeriod === 'thisWeek'" class="weekly-stats">
                  <div class="week-grid-vertical">
                    <div v-for="(day, index) in weekDays" :key="index" class="day-row">
                      <div class="day-label">{{ day.label }}</div>
                      <div class="day-rank" :class="{ 'has-rank': day.rank !== '--' }">
                        {{ day.rank }}
                      </div>
                    </div>
                  </div>
                  <button class="share-btn mt-4">
                    <div class="flex items-center justify-center gap-2">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        width="16"
                        height="16"
                        fill="currentColor"
                      >
                        <path d="M18 16.08c-.76 0-1.44.3-1.96.77L8.91 12.7c.05-.23.09-.46.09-.7s-.04-.47-.09-.7l7.05-4.11c.54.5 1.25.81 2.04.81 1.66 0 3-1.34 3-3s-1.34-3-3-3-3 1.34-3 3c0 .24.04.47.09.7L8.04 9.81C7.5 9.31 6.79 9 6 9c-1.66 0-3 1.34-3 3s1.34 3 3 3c.79 0 1.5-.31 2.04-.81l7.12 4.16c-.05.21-.08.43-.08.65 0 1.61 1.31 2.92 2.92 2.92s2.92-1.31 2.92-2.92c0-1.61-1.31-2.92-2.92-2.92zM18 4c.55 0 1 .45 1 1s-.45 1-1 1-1-.45-1-1 .45-1 1-1zM6 13c-.55 0-1-.45-1-1s.45-1 1-1 1 .45 1 1-.45 1-1 1zm12 7.02c-.55 0-1-.45-1-1s.45-1 1-1 1 .45 1 1-.45 1-1 1z" />
                      </svg>
                      Share
                    </div>
                  </button>
                </div>

                <div v-else class="stats-list">
                  <ul>
                    <li>
                      <div>Words found</div>
                       <div style="font-size: 16px; color: black; font-weight: 600;">0</div>
                    </li>
                    <li>
                      <div>Games played</div>
                      <div  style="font-size: 16px; color: black; font-weight: 600;">0</div>
                    </li>
                    <li>
                      <div>Average guessed words</div>
                      <div   style="font-size: 16px; color: black; font-weight: 600;">0</div>
                    </li>
                    <li>
                      <div>Pangram found</div>
                      <div  style="font-size: 16px; color: black; font-weight: 600;">0</div>
                    </li>
                    <li>
                      <div>Most frequent ranking</div>
                        <div style="font-size: 16px; color: black; font-weight: 600;">0</div>
                    </li>
                  </ul>
                  <button class="share-btn mt-4">
                    <div class="flex items-center justify-center gap-2">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        width="16"
                        height="16"
                        fill="currentColor"
                      >
                        <path d="M18 16.08c-.76 0-1.44.3-1.96.77L8.91 12.7c.05-.23.09-.46.09-.7s-.04-.47-.09-.7l7.05-4.11c.54.5 1.25.81 2.04.81 1.66 0 3-1.34 3-3s-1.34-3-3-3-3 1.34-3 3c0 .24.04.47.09.7L8.04 9.81C7.5 9.31 6.79 9 6 9c-1.66 0-3 1.34-3 3s1.34 3 3 3c.79 0 1.5-.31 2.04-.81l7.12 4.16c-.05.21-.08.43-.08.65 0 1.61 1.31 2.92 2.92 2.92s2.92-1.31 2.92-2.92c0-1.61-1.31-2.92-2.92-2.92zM18 4c.55 0 1 .45 1 1s-.45 1-1 1-1-.45-1-1 .45-1 1-1zM6 13c-.55 0-1-.45-1-1s.45-1 1-1 1 .45 1 1-.45 1-1 1zm12 7.02c-.55 0-1-.45-1-1s.45-1 1-1 1 .45 1 1-.45 1-1 1z" />
                      </svg>
                      Share
                    </div>
                  </button>
                </div>
              </div>

              <!-- Games Tab Content -->
              <div v-if="activeTab === 'games'" class="tab-content">
                <div class="games-list">
                  <div v-for="(game, index) in games" :key="index" class="game-card">
                    <div class="game-content">
                      <div class="game-info">
                        <div class="letters-column">
                          <span 
                            v-for="(letter, i) in game.letters" 
                            :key="i"
                            :class="{ 'center-letter': letter === game.centerLetter }"
                            class="letter-tile"
                          >
                            {{ letter }}
                          </span>
                        </div>
                        <div class="status-column">
                          <span class="rank">{{ game.rank }}</span>
                          <span class="score">{{ game.score }}</span>
                        </div>
                        <button class="delete-game-btn" @click="confirmDelete(index)">
                          <i class="fas fa-trash"></i>
                        </button>
                      </div>
                    </div>
                  </div>
                </div>
                <button class="add-game-btn" @click="addNewGame" v-if="games.length < 5">+</button>
              </div>
            </div>
          </div>
        </div>

        <div class="right-section2">
          <!-- Add cursor animation to word display -->
          <div class="current-word">
            <span v-for="(letter, index) in currentWord" 
                  :key="index" 
                  :class="{ 'green-letter': letter === 'A' }"
                  class="word-letter">
              {{ letter }}
            </span>
            <span class="cursor">|</span>
          </div>

          <div class="hexagon-section">
            <div v-if="!isLoading" class="hexagon-grid">
              <div class="hexagon-row">
                <div 
                  v-for="(letter, index) in [letters[0], centerLetter, letters[2]]" 
                  :key="index"
                  :class="{ 'center-letter': letter === centerLetter }"
                  class="circle" 
                  @click="addLetter(letter)"
                >
                  {{ letter }}
                </div>
              </div>
              <div class="hexagon-row">
                <div 
                  v-for="(letter, index) in letters.slice(3)" 
                  :key="index + 3"
                  class="circle" 
                  @click="addLetter(letter)"
                >
                  {{ letter }}
                </div>
              </div>
            </div>
            <div v-else class="loading-spinner">
              Loading...
            </div>
            
            <div class="controls">
              <div class="flex justify-end gap-2">
                <a class="text-[14px] underline cursor-pointer" 
                   @click="clearWord" 
                   title="Reset all letters">Reset</a>
                <a class="text-[14px] underline cursor-pointer" 
                   @click="deleteLastLetter" 
                   title="Delete last letter">Delete</a>
              </div>
              <div>
                <button class="right-section-btn" 
                        style="width: 100%" 
                        @click="submitWord">Submit</button>
              </div>

              <div style="display: flex; gap: 10px; justify-content: center">
                
                <button class="right-section-btn " style="width: 100%; background: #28a745; color: white">
                  Ranking
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- How to Play popup -->
      <div v-if="showHowToPlay" class="popup-overlay" @click="showHowToPlay = false">
        <div class="popup-content" @click.stop>
          <div class="popup-header">
            <h2 class="text-[16px] font-[600] text-[#333]">How to Play</h2>
            <button class="close-btn" @click="showHowToPlay = false">×</button>
          </div>
          
          <div class="popup-body">
            <p class="text-[14px] mb-[16px] text-[#666]">Create words using the letters shown.</p>
            
            <div class="rules-section mb-[24px]">
              <ul class="text-[14px] space-y-[12px] text-[#444]">
                <li class="flex items-start gap-2">
                  <span class="bullet">•</span>
                  <span>Words must have at least four letters.</span>
                </li>
                <li class="flex items-start gap-2">
                  <span class="bullet">•</span>
                  <span>Words must include the center letter.</span>
                </li>
                <li class="flex items-start gap-2">
                  <span class="bullet">•</span>
                  <span>Letters can be used more than once.</span>
                </li>
                <li class="flex items-start gap-2">
                  <span class="bullet">•</span>
                  <span>Words with hyphens, proper nouns, vulgarities, and especially obscure words are not in the word list.</span>
                </li>
              </ul>
            </div>

            <div class="points-section mb-[24px]">
              <h3 class="text-[14px] font-[600] mb-[12px] text-[#333]">Every word earns points:</h3>
              <ul class="text-[14px] space-y-[12px] text-[#444]">
                <li class="flex items-start gap-2">
                  <span class="bullet">•</span>
                  <span>Four-letter words are worth 1 point.</span>
                </li>
                <li class="flex items-start gap-2">
                  <span class="bullet">•</span>
                  <span>Five-letter words and longer are worth 1 point for each letter.</span>
                </li>
                <li class="flex items-start gap-2">
                  <span class="bullet">•</span>
                  <span>Pangrams — words that use all seven letters — earn an additional bonus of 7 points. There is at least one pangram per puzzle.</span>
                </li>
              </ul>
            </div>

            <p class="text-[14px] text-[#666] mt-[20px] pt-[20px] border-t border-[#eee]">
              New puzzles generate at midnight.
            </p>
          </div>
        </div>
      </div>

      <!-- Rankings popup -->
      <div v-if="showRankings" class="popup-overlay" @click="showRankings = false">
        <div class="popup-content" @click.stop>
          <div class="popup-header">
            <h3>Rankings</h3>
            <button class="close-btn" @click="showRankings = false">×</button>
          </div>

          <p class="rankings-description">
            The points required for each rank are fixed percentages of the maximum number of points
            for each puzzle.
          </p>

          <p class="today-puzzle-text">Today's puzzle:</p>

          <div class="rankings-list">
            <ul>
              <li>Beginner (0 points)</li>
              <li>Novice (11 points)</li>
              <li>Okay (23 points)</li>
              <li>Good (34 points)</li>
              <li>Very Good (46 points)</li>
              <li>Superb (58 points)</li>
              <li>Fabulous (93 points)</li>
              <li>Exceptional (128 points)</li>
              <li>Genius (163 points)</li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Games popup -->
      <div v-if="showGames" class="popup-overlay" @click="showGames = false">
        <div class="popup-content games-popup" @click.stop>
          <div class="popup-header">
            <h3>Games</h3>
            <button class="close-btn" @click="showGames = false">×</button>
          </div>

          <div class="games-list">
            <div v-for="(game, index) in games" :key="index" class="game-card">
              <div class="game-content">
                <div class="game-info">
                  <div class="letters-column">
                    <span 
                      v-for="(letter, i) in game.letters" 
                      :key="i"
                      :class="{ 'center-letter': letter === game.centerLetter }"
                      class="letter-tile"
                    >
                      {{ letter }}
                    </span>
                  </div>
                  <div class="status-column">
                    <span class="rank">{{ game.rank }}</span>
                    <span class="score">{{ game.score }}</span>
                  </div>
                  <button class="delete-game-btn" @click="confirmDelete(index)">
                    <i class="fas fa-trash"></i>
                  </button>
                </div>
              </div>
            </div>
          </div>

          <button class="add-game-btn" @click="addNewGame" v-if="games.length < 5">+</button>
        </div>
      </div>

      <!-- Delete Confirmation popup -->
      <div v-if="showDeleteConfirm" class="popup-overlay" @click="showDeleteConfirm = false">
        <div class="delete-content" @click.stop>
          <p>Are you sure you want to delete this game?</p>
          <div class="confirm-buttons">
            <button class="ok-btn" @click="deleteGame">OK</button>
            <button class="cancel-btn" @click="cancelDelete">Cancel</button>
          </div>
        </div>
      </div>

      <!-- Add new rank popup -->
      <div v-if="showRankDetails" class="popup-overlay" @click="closeRankPopup">
        <div class="popup-content rank-popup" @click.stop>
          <div class="flex justify-between items-center mb-[10px] text-[14px] font-[600]">
            <div>{{ selectedRank.title }}</div>
            <button class="close-btn" @click="closeRankPopup">×</button>
          </div>
          <div class="popup-body">
            <div class="rank-details">
              <div class="points">Points needed: {{ selectedRank.points }}</div>
              <div class="description">{{ selectedRank.description }}</div> 
            </div>
          </div>
        </div>
      </div>

      <!-- Add this new popup component -->
      <div v-if="showRankStats" class="popup-overlay" @click="closeRankStats">
        <div class="popup-content rank-stats-popup" @click.stop>
          <div class="popup-header">
            <button class="back-btn" @click="closeRankStats">
              <i class="fas fa-arrow-left"></i>
            </button>
            <div class="tab-buttons">
              <button 
                :class="{ active: activeRankTab === 'stats' }"
                @click="activeRankTab = 'stats'"
              >
                Statics
              </button>
              <button 
                :class="{ active: activeRankTab === 'previous' }"
                @click="activeRankTab = 'previous'"
              >
                Previous
              </button>
            </div>
          </div>

          <div class="rank-stats-content">
            <!-- Stats Tab Content -->
            <div v-if="activeRankTab === 'stats'" class="stats-tab">
              <div class="letters-display">
                <span>D E K O R W Y</span>
                <div class="points-info">0 of 27 words | 0 of 117 points</div>
              </div>

              <div class="words-grid">
                <div v-for="word in rankWords" 
                     :key="word" 
                     class="word-item"
                     :class="{ found: foundWords.includes(word) }">
                  {{ word }}
                </div>
              </div>
            </div>

            <!-- Previous Tab Content -->
            <div v-if="activeRankTab === 'previous'" class="previous-tab">
              <!-- Add your previous tab content here -->
            </div>
          </div>
        </div>
      </div>
    </div> 
    <Play />
    <Faq />
    <Footer />

    <!-- Add this new popup for points -->
    <div v-if="showPointsPopup" class="points-popup">
      +{{ pointsEarned }} points
    </div>

    <!-- Add validation popup -->
    <div v-if="showValidationPopup" class="validation-popup">
      {{ validationMessage }}
    </div>
  </div>
</template>

<script setup>
import Play from './components/HowToPlay.vue'
import Faq from './components/Faq.vue'
import Footer from './components/Footer.vue'
import { ref, onMounted, onUnmounted, watch } from 'vue'
import axios from 'axios' // Make sure to import axios
import Cookies from 'js-cookie' // Make sure to install this package: npm install js-cookie
import LoginPopup from './components/LoginPopup.vue'
import SignupPopup from './components/SignupPopup.vue'
import ForgotPasswordPopup from './components/ForgotPasswordPopup.vue'
// At the top of the script section, import router
import { useRouter } from 'vue-router'

const showHowToPlay = ref(false)
const activeTab = ref('stats')
const statsPeriod = ref('thisWeek')
const showMenu = ref(false)
const showFoundWords = ref(false)
const showRankings = ref(false)
const showGames = ref(false)
const showDeleteConfirm = ref(false)
const foundWords = ref([])
const gameMode = ref('today')
const games = ref([
  { letters: 'BDEHLOR', centerLetter: 'H', rank: 'Beginner', score: 0 },
  { letters: 'DEGIKLO', centerLetter: 'G', rank: 'Beginner', score: 0 },
  { letters: 'EINORTX', centerLetter: 'X', rank: 'Beginner', score: 0 },
  { letters: 'BEILNTU', centerLetter: 'B', rank: 'Beginner', score: 0 },
  { letters: 'ACIPTVY', centerLetter: 'C', rank: 'Beginner', score: 0 },
])
const gameToDelete = ref(null)
const isDarkMode = ref(false)
const currentWord = ref('')


const showPointsPopup = ref(false)
const pointsEarned = ref(0)


const showValidationPopup = ref(false)
const validationMessage = ref('')


const totalPoints = ref(0)  // Initialize points to 0
const currentLevel = ref('B') // Starting level
const levelThresholds = {
  B: 0,    // Beginner
  N: 11,   // Novice 
  A: 23,   // Advanced
  V: 34,   // Very Good
  S: 46,   // Superb
  F: 93,   // Fabulous
  E: 128,  // Exceptional
  G: 163   // Genius
}


const username = ref(Cookies.get('username') || null) // Get username from cookie
const userPoints = ref(0)
const userWords = ref([])
const isLoading = ref(true) // This is the only isLoading declaration we'll keep


const letters = ref(['H', 'D', 'Z', 'N', 'X', 'V', 'W'])
const centerLetter = ref('X')


const router = useRouter()

const toggleMenu = () => {
  showMenu.value = !showMenu.value
}

const closeMenu = () => {
  showMenu.value = false
}

const toggleDarkMode = () => {
  isDarkMode.value = !isDarkMode.value
}


const vClickOutside = {
  mounted(el, binding) {
    el._clickOutside = (event) => {
      if (!(el === event.target || el.contains(event.target))) {
        binding.value(event)
      }
    }
    document.addEventListener('click', el._clickOutside, true)
  },
  unmounted(el) {
    document.removeEventListener('click', el._clickOutside, true)
  },
}

const statsData = {
  thisWeek: {
    wordsFound: 0,
    gamesPlayed: 0,
    avgGuessedWords: 0,
    pangramFound: 0,
    frequentRanking: 'none',
  },
  lastWeek: {
    wordsFound: 0,
    gamesPlayed: 0,
    avgGuessedWords: 0,
    pangramFound: 0,
    frequentRanking: 'none',
  },
  allTime: {
    wordsFound: 0,
    gamesPlayed: 0,
    avgGuessedWords: 0,
    pangramFound: 0,
    frequentRanking: 'none',
  },
  unlimited: {
    wordsFound: 1,
    gamesPlayed: 1,
    avgGuessedWords: 1.0,
    pangramFound: 0,
    frequentRanking: 'Beginner',
  },
}

const toggleFoundWords = () => {
  showFoundWords.value = !showFoundWords.value
}

const setGameMode = (mode) => {
  gameMode.value = mode
  activeTab.value = 'stats' // Always default to stats tab when changing modes
}

const generateRandomWord = () => {
  const alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
  let result = ''
  for (let i = 0; i < 6; i++) {
    result += alphabet.charAt(Math.floor(Math.random() * alphabet.length))
  }
  return result
}

const addNewGame = () => {
  if (games.value.length >= 5) return
  const newWord = generateRandomWord()
  const centerLetter = newWord[Math.floor(Math.random() * newWord.length)]
  games.value.push({
    letters: newWord,
    centerLetter,
    rank: 'Beginner',
    score: 0,
  })
}

const confirmDelete = (index) => {
  gameToDelete.value = index
  showDeleteConfirm.value = true
}

const cancelDelete = () => {
  showDeleteConfirm.value = false
  gameToDelete.value = null
  // Don't close the games popup
}

const deleteGame = () => {
  if (gameToDelete.value !== null) {
    games.value.splice(gameToDelete.value, 1)
    gameToDelete.value = null
    showDeleteConfirm.value = false
    // Don't close the games popup
  }
}

// Update the weekDays data
const weekDays = ref([
  { label: 'S', rank: '--' },
  { label: 'M', rank: 'Beginner' },
  { label: 'T', rank: '--' },
  { label: 'W', rank: '--' },
  { label: 'T', rank: '--' },
  { label: 'F', rank: '--' },
  { label: 'S', rank: '--' },
])

// Update click-outside handler for games popup
const closeGames = () => {
  if (!showDeleteConfirm.value) {
    showGames.value = false
  }
}

const showTooltip = (event, letter) => {
  // Remove any existing tooltips
  const existingTooltip = document.querySelector('.tooltip');
  if (existingTooltip) {
    existingTooltip.remove();
  }

  // Create and show new tooltip
  const tooltip = document.createElement('div');
  tooltip.className = 'tooltip';
  tooltip.textContent = `Letter ${letter}`;
  
  // Position tooltip above the circle
  const rect = event.target.getBoundingClientRect();
  tooltip.style.left = `${rect.left + (rect.width / 2)}px`;
  tooltip.style.top = `${rect.top - 40}px`;
  
  document.body.appendChild(tooltip);
  
  // Remove tooltip after 2 seconds
  setTimeout(() => {
    tooltip.remove();
  }, 2000);
}

const addLetter = (letter) => {
  currentWord.value += letter
}

const clearWord = () => {
  currentWord.value = ''
}

const deleteLastLetter = () => {
  currentWord.value = currentWord.value.slice(0, -1)
}

// Function to calculate points for a word
const calculatePoints = (word) => {
  if (word.length === 4) return 1
  return word.length // 1 point per letter for words longer than 4 letters
}

// Function to update points and level
const updatePointsAndLevel = () => {
  // Calculate total points from found words
  const points = foundWords.value.reduce((total, word) => {
    return total + calculatePoints(word)
  }, 0)
  
  totalPoints.value = points

  // Update level based on points
  for (const [level, threshold] of Object.entries(levelThresholds)) {
    if (points >= threshold) {
      currentLevel.value = level
    }
  }

 
  const currentThreshold = levelThresholds[currentLevel.value]
  const levels = Object.keys(levelThresholds)
  const currentIndex = levels.indexOf(currentLevel.value)
  const nextLevel = levels[currentIndex + 1]
  const pointsToNext = nextLevel ? levelThresholds[nextLevel] - points : 0

 
  const subtext = nextLevel ? 
    `( ${points} points | Just ${pointsToNext} more to reach ${nextLevel} level )` :
    `( ${points} points | Maximum level reached! )`

  // Update DOM element
  const subtextElement = document.querySelector('.subtext')
  if (subtextElement) {
    subtextElement.textContent = subtext
  }

  // Save to cookies
  saveGameState()
}

const saveGameState = () => {
  try {
    const gameState = {
      username: username.value,
      points: totalPoints.value,
      words: foundWords.value,
      level: currentLevel.value
    }
    localStorage.setItem('spellsbeeGameState', JSON.stringify(gameState))
  } catch (error) {
    console.error('Error saving game state:', error)
  }
}

const loadGameState = () => {
  try {
    const savedState = localStorage.getItem('spellsbeeGameState')
    if (savedState) {
      const state = JSON.parse(savedState)
      username.value = state.username || ''
      totalPoints.value = state.points || 0
      foundWords.value = state.words || []
      currentLevel.value = state.level || 'B'
    }
  } catch (error) {
    console.error('Error loading game state:', error)
  }
}

const fetchUserData = async () => {
  if (!username.value) {
    console.log('No username available, skipping fetch')
    return
  }
  
  isLoading.value = true
  try {
    const response = await axios.get(`${import.meta.env.VITE_API_URL}/api/stats/user/${username.value}`)
    if (response.data) {
      userPoints.value = response.data.points || 0
      userWords.value = response.data.words || []
      foundWords.value = userWords.value  
      totalPoints.value = userPoints.value
      updatePointsAndLevel()
    }
  } catch (error) {
    console.error('Error fetching user data:', error)
    userPoints.value = 0
    userWords.value = []
    foundWords.value = []
    totalPoints.value = 0
  } finally {
    isLoading.value = false
  }
}

const fetchDailyLetters = async () => {
  try {
    const response = await axios.get(`${import.meta.env.VITE_API_URL}/api/daily-letters`);
    letters.value = response.data.letters;
    centerLetter.value = response.data.letters[1];
    isLoading.value = false
  } catch (error) {
    console.error('Error fetching daily letters:', error)
    letters.value = ['A', 'B', 'C', 'D', 'E', 'F', 'G']
    centerLetter.value = 'B'
    isLoading.value = false
  }
}

const submitWord = async () => {
  if(username.value === null){
     username.value = Cookies.get('username');
  }
  try {
    if (!currentWord.value.includes(centerLetter.value)) {
      validationMessage.value = `Word must include the green letter "${centerLetter.value}"`
      showValidationPopup.value = true
      setTimeout(() => {
        showValidationPopup.value = false
      }, 1500)
      return
    }

    if (currentWord.value.length < 4) {
      validationMessage.value = 'Word must be at least 4 letters'
      showValidationPopup.value = true
      setTimeout(() => {
        showValidationPopup.value = false
      }, 1500)
      return
    }

    if (foundWords.value.includes(currentWord.value)) {
      validationMessage.value = 'Word already found'
      showValidationPopup.value = true
      setTimeout(() => {
        showValidationPopup.value = false
      }, 1500)
      return
    }

    const validateResponse = await axios.get(`${import.meta.env.VITE_API_URL}/api/spellsbee/validate/${currentWord.value}`)
    
    if (validateResponse.data.isValid) {
     
      const saveResponse = await axios.post(`${import.meta.env.VITE_API_URL}/api/stats/user/${username.value}`, {
        word: currentWord.value,
        points: calculatePoints(currentWord.value)
      })

      if (saveResponse.data) {
        foundWords.value.push(currentWord.value)
        
        const newPoints = calculatePoints(currentWord.value)
        pointsEarned.value = newPoints
        
        showPointsPopup.value = true
        setTimeout(() => {
          showPointsPopup.value = false
        }, 1500)

        await fetchUserData()
      }
    } else {
      validationMessage.value = 'Incorrect'
      showValidationPopup.value = true
      setTimeout(() => {
        showValidationPopup.value = false
      }, 1500)
    }
  } catch (error) {
    console.error('Error submitting word:', error)
    validationMessage.value = 'Error submitting word'
    showValidationPopup.value = true
    setTimeout(() => {
      showValidationPopup.value = false
    }, 1500)
  }
  
  currentWord.value = ''
}

onMounted(async () => {
  loadUserProfile() // Load user profile from cookies
  await fetchDailyLetters() // Wait for letters to load
  await fetchUserData() // Then fetch user data
  
  // Add event listeners
  window.addEventListener('keydown', handleKeyPress)
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  window.removeEventListener('keydown', handleKeyPress)
  document.removeEventListener('click', handleClickOutside)
})

const handleKeyPress = (event) => {
  const key = event.key.toUpperCase()
  
  if (event.key === 'Backspace' || event.key === 'Delete') {
    deleteLastLetter()
    return
  }

  if (event.key === 'Enter') {
    submitWord()
    return
  }

  if (letters.value.includes(key)) {
    addLetter(key)
  }
}

// Watch for changes in foundWords and update points
watch(foundWords, () => {
  updatePointsAndLevel()
}, { deep: true })


const getRankBackground = (letter, index) => {
  const ranks = ['B', 'N', 'A', 'V', 'S', 'F', 'E', 'G']
  const currentRankIndex = ranks.indexOf(currentLevel.value)
  
  if (index <= currentRankIndex) {
    return '#28a745' // Green for current and previous ranks
  } else if (index === currentRankIndex + 1) {
    return '#ffe999' // Yellow for next rank
  } else {
    return '#e6e6e6' // Gray for remaining ranks
  }
}

const getRankTooltip = (letter) => {
  const rankInfo = {
    'B': 'Beginner (0 points)',
    'N': 'Novice (11 points)',
    'A': 'Advanced (23 points)',
    'V': 'Very Good (34 points)',
    'S': 'Superb (46 points)',
    'F': 'Fabulous (93 points)',
    'E': 'Exceptional (128 points)',
    'G': 'Genius (163 points)'
  }
  return rankInfo[letter]
}

const activePopup = ref(null)

const showLoginPopup = () => {
  activePopup.value = 'login'
}

const closePopup = () => {
  activePopup.value = null
}

// Add these new functions
const loadUserProfile = () => {
  try {

    const userData = localStorage.getItem('userLoginData');

    if (userData) {
      userProfile.value = JSON.parse(userData)
      if (userProfile.value?.username) {
        username.value = userProfile.value.username
      }
    }
  } catch (error) {
    console.error('Error loading user profile:', error)
    userProfile.value = null
  }
}

const toggleProfileMenu = () => {
  showProfileMenu.value = !showProfileMenu.value
}

const logout = () => {
  Cookies.remove('username'); 
  localStorage.removeItem('userLoginData');
  userProfile.value = null
  showProfileMenu.value = false
  window.location.reload()
}

const handleClickOutside = (event) => {
  if (profileBtn.value && 
      profileMenu.value && 
      !profileBtn.value.contains(event.target) && 
      !profileMenu.value.contains(event.target)) {
    showProfileMenu.value = false
  }
}

const handleLoginSuccess = () => {
  activePopup.value = null
  router.push('/')
}

const handleSignupSuccess = () => {
  activePopup.value = null
  router.push('/')
}

// Add these refs for user profile and menu state
const userProfile = ref(null)
const showProfileMenu = ref(false)
const profileBtn = ref(null)
const profileMenu = ref(null)

// Add these new refs
const showRankDetails = ref(false)
const selectedRank = ref({
  title: '',
  points: 0,
  description: ''
})

// Add this new object for rank details
const rankDetails = {
  'B': {
    title: 'Beginner',
    points: 0,
    description: 'Starting rank for all players. Keep playing to improve!'
  },
  'N': {
    title: 'Novice',
    points: 11,
    description: 'You\'re getting the hang of it! Keep finding more words.'
  },
  'A': {
    title: 'Advanced',
    points: 23,
    description: 'You\'re making good progress. Your vocabulary is expanding!'
  },
  'V': {
    title: 'Very Good',
    points: 34,
    description: 'Impressive word skills! You\'re becoming a word master.'
  },
  'S': {
    title: 'Superb',
    points: 46,
    description: 'Outstanding performance! Your word knowledge is remarkable.'
  },
  'F': {
    title: 'Fabulous',
    points: 93,
    description: 'Exceptional word mastery! You\'re among the top players.'
  },
  'E': {
    title: 'Exceptional',
    points: 128,
    description: 'Amazing achievement! Your word skills are extraordinary.'
  },
  'G': {
    title: 'Genius',
    points: 163,
    description: 'Ultimate mastery! You\'ve reached the highest rank possible.'
  }
}

// Add these new functions
const showRankPopup = (letter) => {
  if (letter === 'B') { // Show stats popup only for Beginner rank
    showRankStats.value = true
  } else {
    selectedRank.value = rankDetails[letter]
    showRankDetails.value = true
  }
}

const closeRankPopup = () => {
  showRankDetails.value = false
}

const showRankStats = ref(false)
const activeRankTab = ref('stats')

// Sample words array - replace with your actual words
const rankWords = [
  'Dewy', 'Doddery', 'Dorky', 'Dory', 'Dowdy',
  'Dowry', 'Dryer', 'Dyed', 'Dyer', 'Eddy',
  'Eyed', 'Keyed', 'Keyword', 'Kooky',
  'Redye', 'Redyed', 'Reedy', 'Rookery',
  'Rowdy', 'Weedy', 'Woody', 'Wordy',
  'Worry', 'Wryer', 'Yoke', 'Yoked', 'Yore'
]

const closeRankStats = () => {
  showRankStats.value = false
}

// Add this new function
const openRankStatsPopup = () => {
  showRankStats.value = true
  activeRankTab.value = 'stats' // Set default tab to stats
}

</script>

<style scoped>
/* Updated styles for a more modern look */
.game-container {
  max-width: 800px;
  margin: 0 auto;
  min-height: 100vh;
  font-family: 'Inter', sans-serif;
  padding: 0 16px; /* Add padding for mobile */
  background-color: var(--bg-color, #ffffff); /* Use system default with fallback */
}

.nav-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: white;
  padding: 12px 20px;
  border-radius: 16px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.06);
  margin-bottom: 16px;
  min-height: 56px;
}

.left-section {
  flex: 1;
}

.help-section {
  display: flex;
  align-items: center;
  gap: 8px;
}

.center-section {
  flex: 2;
  display: flex;
  justify-content: center;
}

.right-section {
  flex: 1;
  display: flex;
  justify-content: flex-end;
}

.flower-icon {
  height: 24px;
  width: 24px;
}

.mode-toggles {
  display: flex;
  background: #f5f5f5;
  padding: 4px;
  border-radius: 24px;
  gap: 4px;
  align-items: center;
}

.mode-btn {
  padding: 8px 16px;
  border-radius: 20px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.mode-btn.active {
  background: #ffe999;
  color: #333;
  box-shadow: 0 2px 8px rgba(255, 233, 153, 0.4);
}

.nav-icon-btn {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  border: none;
  background: none;
  color: #666;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
}

/* Add specific styles for help button */

.nav-icon-btn:last-child:hover {
  opacity: 0.9;
}

/* Dark mode adjustments */
@media (prefers-color-scheme: dark) {
  .nav-icon-btn:last-child {
    background: #ffe999;
    color: #333; /* Keep text dark on yellow background */
  }
}

.nav-link {
  color: #666;
  text-decoration: none;
  font-size: 14px;
  font-weight: 500;
  transition: color 0.2s ease;
}

.nav-link:hover {
  color: #333;
}

.level-info {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.level-title {
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.subtext {
  margin-top: 2px;
  font-size: 14px;
  color: #666;
}

.progress-dots {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 4px 0;
}

.dot {
  cursor: pointer;
  border-radius: 50%;
  background: #e6e6e6;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-weight: 500;
  color: #333;
  transition: all 0.2s ease;
}

.dot.active:nth-child(1) {
  background: #28a745;
  color: white;
}

.dot.active:nth-child(2) {
  background: #ffe999;
  color: #333;
}

.category-selector select {
  width: 100%;
  padding: 12px 20px;
  border-radius: 25px;
  border: 1px solid #e6e6e6;
  background: white;
  font-size: 16px;
  appearance: none;
  background-image: url("data:image/svg+xml;charset=UTF-8,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='currentColor' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3e%3cpolyline points='6 9 12 15 18 9'%3e%3c/polyline%3e%3c/svg%3e");
  background-repeat: no-repeat;
  background-position: right 1rem center;
  background-size: 1em;
}

.hexagon-grid {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 5px;
  width: fit-content;
  margin: 0 auto;
}

.hexagon-row {
  display: flex;
  gap: 5px;
}

.circle {
  width: 60px;
  height: 60px;
  background: #e6e6e6;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  font-weight: bold;
  cursor: pointer;
  transition: all 0.2s ease;
  margin: 5px;
}

.circle:hover {
  transform: translateY(-2px);
}

.circle.green {
  background: #28a745;
  color: white;
}

.controls {
  display: flex;

  gap: 10px;

  display: flex;
  flex-direction: column;
  width: 100%;
  padding: 20px;
}

.controls button {
  font-weight: 500;
}

.controls button:hover {
  transform: translateY(-1px);
}

.right-section-btn {
  background: #ffe999;
  font-size: 14px;
  font-weight: 500;
  height: 40px;
  display: flex;
  border-radius: 8px;
  padding: 10px 16px;
  align-items: center;
  justify-content: center;
  color: #333;
}

.shuffle-btn {
  width: 40px;
  height: 40px;
  border-radius: 50% !important;
  padding: 0 !important;
  display: flex;
  align-items: center;
  justify-content: center;
}

.delete-btn {
  width: 40px;
  height: 40px;
  font-size: 14px;
  border-radius: 50%;
  border: none;
  background: white;
  font-weight: 500;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.delete-btn .delete-icon {
  color: #dc3545; /* Red color */
  font-size: 16px; /* Adjust size to match other icons */
}

.delete-btn:hover {
  transform: translateY(-1px);
}

.divider {
  border: none;
  height: 1px;
  background-color: #eee;
  margin: 0;
  width: 100%;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .divider {
    background-color: #333;
  }
}

/* Add these new styles */
.tabs-section {
  background: white;
  border-radius: 12px;
  overflow: hidden;
  border: 1px solid #eee;
}

.tabs-header {
  display: flex;
  border-bottom: 1px solid #eee;
}

.tab-button {
  flex: 1;
  padding: 8px;
  border: none;
  background: none;
  font-size: 14px;
  font-weight: 500;
  color: #666;
  cursor: pointer;
  transition: all 0.2s ease;
}

.tab-button.active {
  color: #333;
  background: #ffe999;
}

.tab-content {
  padding: 16px;
}

.stats-period-tabs {
  display: flex;
  gap: 8px;
  margin-bottom: 16px;
}

.stats-checkboxes {
  font-size: 13px;
  font-weight: 600;
  color: #666;
  display: flex;
  gap: 18px;
  margin-bottom: 16px;
}

.period-tab {
  flex: 1;
  padding: 6px;
  border: none;
  background: #f5f5f5;
  border-radius: 4px;
  font-size: 13px;
  color: #666;
  cursor: pointer;
  transition: all 0.2s ease;
}

.period-tab.active {
  background: #ffe999;
  color: #333;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .tabs-section {
    background: #2a2a2a;
    border-color: #333;
  }

  .tabs-header {
    border-color: #333;
  }

  .tab-button {
    color: #999;
  }

  .tab-button.active {
    color: #333;
    background: #ffe999;
  }

  .period-tab {
    background: #333;
    color: #999;
  }

  .period-tab.active {
    background: #ffe999;
    color: #333;
  }
}

.found-words-container {
  flex: 1; /* Take up remaining space */
}

.found-words-header {
  padding: 0;
  background: white;
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: pointer;
  user-select: none;
  font-size: 14px;
  color: #000000;
}

.found-words-list {
  padding: 12px 0 0 0;
  background: white;
  border-top: none;
}

.found-word {
  padding: 4px 0;
  font-size: 14px;
  color: #000000;
  font-weight: 500;
  color: #666;
}

.toggle-arrow {
  transition: transform 0.2s ease;
  color: rgb(102, 102, 102);
}

.toggle-arrow.rotated {
  transform: rotate(180deg);
}

.rankings-description {
  color: #666;
  font-size: 14px;
  margin-bottom: 16px;
}

.today-puzzle-text {
  font-weight: 500;
  margin-bottom: 8px;
}

.rankings-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.rankings-list ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.rankings-list li {
  padding: 8px 0;
  display: flex;
  align-items: flex-start;
  gap: 8px;
}

.rankings-list li::before {
  content: '•';
  color: #666;
}

.games-popup {
  background: white;
  border-radius: 16px;
  padding: 20px;
  width: 90%;
  max-width: 500px;
}

.games-list {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.game-card {
  background: white;
  border-radius: 16px;
  overflow: hidden;
  border: 1px solid #ffe999;
}

.game-content {
  padding: 16px;
  background: white;
  border-radius: 12px;
}

.game-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
}

.letters-column {
  display: flex;
  gap: 8px;
  align-items: center;
}

.letter-tile {
  font-size: 15px;
  font-weight: 500;
  padding: 4px;
}

.letter-tile.center-letter {
  background: #ffe999;
  border-radius: 4px;
}

.status-column {
  display: flex;
  align-items: center;
  gap: 12px;
}

.rank {
  font-size: 14px;
  color: #666;
}

.score {
  background: #ffe999;
  color: #333;
  width: 24px;
  height: 24px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 500;
  font-size: 14px;
}

.delete-game-btn {
  background: none;
  border: none;
  color: #666;
  padding: 4px 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: color 0.2s ease;
}

.delete-game-btn:hover {
  color: #dc3545; /* Red color on hover */
}

.add-game-btn {
  width: 48px;
  height: 48px;
  background: #ffe999;
  border: none;
  border-radius: 50%;
  font-size: 24px;
  color: #333;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 24px auto 0;
}

.delete-popup {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1100;
}

.delete-content {
  background: white;
  padding: 24px;
  border-radius: 12px;
  text-align: center;
}

.confirm-buttons {
  display: flex;
  gap: 16px;
  justify-content: center;
  margin-top: 20px;
}

.ok-btn,
.cancel-btn {
  padding: 8px 24px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
}

.ok-btn {
  background: #ffe999;
  color: #333;
}

.ok-btn:hover {
  opacity: 0.9;
}

.cancel-btn {
  background: #f8f9fa;
  color: #333;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  :root {
    --bg-color: #121212;
    --surface-color: #1e1e1e;
    --text-color: #ffffff;
  }

  .game-container {
    color: var(--text-color);
  }

  .nav-bar,
  .found-words-container,
  .popup-content,
  .game-card {
    background-color: var(--surface-color);
    color: var(--text-color);
  }

  .hexagon {
    background-color: var(--surface-color);
    color: var(--text-color);
  }

  .nav-bar {
    background: #2a2a2a;
  }

  .mode-toggles {
    background: #333;
  }

  .mode-btn {
    color: #fff;
  }

  .mode-btn.active {
    background: #ffe999;
    color: #333;
  }

  .nav-link {
    color: #999;
  }

  .nav-link:hover {
    color: #fff;
  }

  .info-button {
    background: #2a2a2a;
    border-color: #444;
    color: #fff;
  }

  .info-button:hover {
    background: #333;
    border-color: #555;
  }

  .menu-button {
    color: #999;
  }

  .menu-button:hover {
    color: #fff;
  }

  .share-btn {
    background: #ffe999;
  }

  .share-btn:hover {
    opacity: 0.9;
  }

  .toggle-indicator {
    color: #2ecc71;
  }

  .stats-content {
    background: #2a2a2a;
  }

  .stats-header {
    border-color: #333;
  }

  .stats-tabs {
    background: #222;
  }

  .tab-btn.active {
    background: #333;
    color: white;
  }

  .day-rank {
    background: #333;
    color: #999;
  }

  .day-rank.has-rank {
    background: #ffe999;
    color: #333;
  }

  .stats-summary {
    background: #222;
  }

  .summary-value {
    color: white;
  }

  .summary-label {
    color: #999;
  }

  .stats-list li {
    background: red;
    display: flex;
    flex-direction: column;
    padding:8px;
    color: white;
    border-bottom: none;
  }

  .stats-list li::before {
    color: #999;
  }

  .game-content {
    background: #2a2a2a;
    border-color: #444;
  }

  .delete-game-btn {
    color: #999;
  }

  .delete-game-btn:hover {
    background: #333;
    color: #fff;
  }
}

.stats-list ul {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.stats-list li {
  padding: 8px;
  border-radius: 4px;
  margin-bottom: 4px;
  background: #e6e6e6;
  display: flex;
  flex-direction: column;
  color: #333;
  font-size: 13px;
}

.stats-list li div:last-child {
  font-weight: 500;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .stats-list li {
    background: #333;
    color: #fff;
  }
}

/* Updated responsive styles */
@media (max-width: 768px) {
  .nav-bar {
    padding: 8px 12px;
  }

  .help-section {
    gap: 4px;
  }

  .flower-icon {
    height: 20px;
    width: 20px;
  }

  .mode-toggles {
    font-size: 13px; /* Smaller font on mobile */
  }

  .mode-btn {
    padding: 6px 12px; /* Smaller padding on mobile */
  }
 
  .hexagon {
    width: 60px; /* Smaller hexagons on mobile */
    height: 69px;
    font-size: 20px;
  }

  .level-info {
    width: 100% !important; /* Full width on mobile */
  }

  .found-words-container {
    width: 100% !important; /* Full width on mobile */
  }

  /* Adjust popup content for mobile */
  .popup-content {
    width: 95%;
    padding: 16px;
    margin: 16px;
  }

  .stats-content {
    max-width: 95%;
  }

  /* Adjust game cards for mobile */
  .game-content {
    flex-direction: column;
    gap: 12px;
    align-items: flex-start;
  }

  .game-info {
    width: 100%;
    justify-content: space-between;
  }
}

/* Add styles for very small screens */
@media (max-width: 480px) {
  .nav-bar {
    flex-direction: row; /* Change from column to row */
    align-items: center; /* Keep items centered */
  }

  .left-section {
    justify-content: flex-start; /* Align to the start instead of center */
  }

  .right-section {
    justify-content: flex-end; /* Align to the end instead of center */
  }

  .hexagon {
    width: 50px; /* Even smaller hexagons */
    height: 58px;
    font-size: 18px;
  }

  .controls {
    gap: 12px;
  }

  .controls button {
    padding: 10px 16px;
    font-size: 14px;
  }

  /* Adjust flex layout for level info and words section */
  .flex.justify-center.items-center.gap-8.mt-4.mb-8 {
    flex-direction: column;
  }

  .w-[50%] {
    width: 100% !important;
  }

  /* Adjust progress dots for mobile */
  .progress-dots {
    gap: 8px;
  }




}

/* Add tablet-specific adjustments */
@media (min-width: 769px) and (max-width: 1024px) {
  .game-container {
    max-width: 90%;
  }

  .hexagon {
    width: 65px;
    height: 75px;
  }
}

/* Adjust dark mode styles for better contrast on mobile */
@media (prefers-color-scheme: dark) {
  @media (max-width: 768px) {
    .game-content {
      background: #222;
      border-color: #333;
    }

    .popup-content {
      background: #222;
    }

    .nav-bar {
      background: #222;
    }
  }
}

/* Add landscape orientation support */
@media (max-height: 600px) and (orientation: landscape) {
  .game-container {
    height: auto;
    min-height: 100vh;
  }

  

  .controls {
    margin-bottom: 20px;
  }
}

/* Add these new styles */
.game-layout {
  display: flex;
  flex-direction: row;
  max-width: 1000px;
  margin: 0 auto;
  justify-content: center;
  min-height: 60vh;
}

.left-section {
  width: 300px;
  display: flex;
  justify-content: flex-start;
}

.right-section {
  width: 400px;
  display: flex;
  justify-content: flex-end;
}
.left-section2 {
  width: 50%;
  padding: 20px;
  border: 1px solid #f2f2f2;
  border-radius: 8px;
}

.right-section2 {
  width: 50%;
}

.right-content {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.hexagon-section {
  display: flex;
  flex-direction: column;
  gap: 24px;
  margin-top: 0;
  padding-top: 0;
  align-items: center;
}

.nav-links {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 8px;
}

.level-header {
  display: flex;
  align-items: center;
  gap: 8px;
}

.level-title {
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.subtext {
  color: #666;
  font-size: 14px;
}

.progress-dots {
  display: flex;
  gap: 8px;
}



.dot.active {
  
 
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-weight: 500;
}

/* Update responsive styles */
@media (max-width: 768px) {
  .game-layout {
    flex-direction: column;
    align-items: center;
    padding: 0 16px;
    gap: 32px;
  }

  .left-section,
  .right-section {
    width: 100%;
    justify-content: center;
    padding: 20px 0;
  }

  .right-content {
    margin: 0 auto;
  }
}

ul {
  list-style-type: none;
}

/* Dark mode styles */
@media (prefers-color-scheme: dark) {
  .nav-link {
    color: #999;
  }

  .nav-link:hover {
    color: #fff;
  }

  .level-title {
    color: #fff;
  }

  .subtext {
    color: #999;
  }
}

/* Add these new styles */
.circle-btn {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: none;
  background: white;
  color: #666;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
  margin-bottom: 8px;
}

.circle-btn:hover {
  transform: translateY(-2px);
}

/* Add dark mode support */
@media (prefers-color-scheme: dark) {
  .circle-btn {
    background: #2a2a2a;
    color: #999;
  }

  .circle-btn:hover {
    color: #fff;
  }
}

/* Add these new styles */
.bottom-buttons {
  display: flex;
  gap: 12px;
  margin-top: auto;
}

.circle-btn {
  width: 40px;
  height: 40px;
  font-size: 20px;
  border-radius: 50%;
  border: none;
  background: white;
  color: #666;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
}

.circle-btn:hover {
  transform: translateY(-2px);
}

/* Update dark mode support */
@media (prefers-color-scheme: dark) {
  .circle-btn {
    background: #2a2a2a;
    color: #999;
  }

  .circle-btn:hover {
    color: #fff;
  }
}

.divider {
  border: none;
  height: 1px;
  background-color: #eee;
  margin: 0;
  width: 100%;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .divider {
    background-color: #333;
  }
}

/* Add these new styles */
.tabs-section {
  background: white;
  border-radius: 12px;
  overflow: hidden;
  border: 1px solid #eee;
}

.tabs-header {
  display: flex;
  border-bottom: 1px solid #eee;
}

.tab-button {
  flex: 1;
  padding: 8px;
  border: none;
  background: none;
  font-size: 14px;
  font-weight: 500;
  color: #666;
  cursor: pointer;
  transition: all 0.2s ease;
}

.tab-button.active {
  color: #333;
  background: #ffe999;
}

.tab-content {
  padding: 16px;
}

.stats-period-tabs {
  display: flex;
  gap: 8px;
  margin-bottom: 16px;
}

.stats-checkboxes {
  font-size: 13px;
  font-weight: 600;
  color: #666;
  display: flex;
  gap: 18px;
  margin-bottom: 16px;
}

.period-tab {
  flex: 1;
  padding: 6px;
  border: none;
  background: #f5f5f5;
  border-radius: 4px;
  font-size: 13px;
  color: #666;
  cursor: pointer;
  transition: all 0.2s ease;
}

.period-tab.active {
  background: #ffe999;
  color: #333;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .tabs-section {
    background: #2a2a2a;
    border-color: #333;
  }

  .tabs-header {
    border-color: #333;
  }

  .tab-button {
    color: #999;
  }

  .tab-button.active {
    color: #333;
    background: #ffe999;
  }

  .period-tab {
    background: #333;
    color: #999;
  }

  .period-tab.active {
    background: #ffe999;
    color: #333;
  }
}

/* Add new styles for share button */
.share-btn {
  background: #ffe999;
  color: #333;
  padding: 8px 16px;
  border-radius: 4px;
  border: none;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  width: 100%;
  transition: background 0.2s ease;
}

.share-btn:hover {
  opacity: 0.9;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .share-btn {
    background: #ffe999;
    color: #333;
  }
  
  .share-btn:hover {
    opacity: 0.9;
  }
}

/* Update tooltip styles for better visibility */
[title] {
  position: relative;
}

[title]:hover::before {
  content: attr(title);
  position: absolute;
  top: 100%; /* Changed from bottom: 100% to top: 100% */
  left: 50%;
  transform: translateX(-50%);
  padding: 6px 10px;
  background: #e6e6e6; /* Changed from rgba(0, 0, 0, 0.8) to #e6e6e6 */
  color: #333; /* Changed from white to #333 */
  font-size: 12px;
  border-radius: 4px;
  white-space: nowrap;
  z-index: 100;
  pointer-events: none;
  margin-top: 8px; /* Changed from margin-bottom to margin-top */
}

[title]:hover::after {
  content: '';
  position: absolute;
  top: 100%; /* Changed from bottom: 100% to top: 100% */
  left: 50%;
  transform: translateX(-50%);
  border-left: 6px solid transparent;
  border-right: 6px solid transparent;
  border-bottom: 6px solid #e6e6e6; /* Changed from border-top to border-bottom and color */
  margin-top: 2px; /* Changed from margin-bottom to margin-top */
}

/* Dark mode support for tooltips - you can remove this if not needed */
@media (prefers-color-scheme: dark) {
  [title]:hover::before {
    background: #e6e6e6;
    color: #333;
  }
  
  [title]:hover::after {
    border-bottom-color: #e6e6e6;
  }
}

/* Add tooltip styles */
.tooltip {
  position: fixed;
  background: #333;
  color: white;
  padding: 8px 12px;
  border-radius: 4px;
  font-size: 14px;
  transform: translateX(-50%);
  z-index: 1000;
  pointer-events: none;
  animation: fadeIn 0.2s ease;
}

.tooltip::after {
  content: '';
  position: absolute;
  bottom: -5px;
  left: 50%;
  transform: translateX(-50%);
  border-left: 5px solid transparent;
  border-right: 5px solid transparent;
  border-top: 5px solid #333;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateX(-50%) translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateX(-50%) translateY(0);
  }
}

.word-display {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  min-height: 40px;
  background: white;
  border-radius: 8px;
  padding: 8px 16px;
  margin-bottom: 16px;
  width: 100%;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-top: 80px;
}

.current-word {
  font-size: 24px;
  font-weight: 500;
  letter-spacing: 2px;
  color: #333;
  text-align: center;
  margin: 80px auto 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;
  width: 100%;
  max-width: 80%;
  min-height: 60px;
  padding: 10px;
  position: relative;
}

.word-letter {
  display: inline-block;
  margin: 0 1px;
}

.current-word .green-letter {
  color: #28a745;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .word-display {
    background: #2a2a2a;
  }
  
  .current-word {
    color: #fff;
  }
  
  .current-word .green-letter {
    color: #28a745;
  }
}

.clear-btn {
  background: none;
  border: none;
  color: #666;
  cursor: pointer;
  padding: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: color 0.2s ease;
}

.clear-btn:hover {
  color: #dc3545;
}

.cursor {
  color: #28a745;
  animation: blink 1s step-end infinite;
  font-weight: normal;
  font-size: 32px;
  margin-left: 2px;
  line-height: 1;
  display: inline-block;
  vertical-align: middle;
  position: relative;
  top: -2px;
}

@keyframes blink {
  from, to { opacity: 1; }
  50% { opacity: 0; }
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .current-word {
    color: #fff;
  }
  
  .current-word .green-letter {
    color: #28a745;
  }

  .cursor {
    color: #28a745;
  }
}

.word-chips {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-top: 12px;
}

.word-chip {
  background: #f5f5f5;
  padding: 4px 12px;
  border-radius: 16px;
 
  color: #333;
  font-weight: 500;
}

.word-chip.pangram {
  background: #ffe999;
  color: #333;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .word-chip {
    background: #333;
    color: #fff;
  }
  
  .word-chip.pangram {
    background: #ffe999;
    color: #333;
  }
}

.week-grid-vertical {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 16px;
}

.day-row {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 4px 0;
}

.day-label {
  font-size: 12px;
  color: #666;
  width: 20px;
}

.day-rank {
  flex: 1;
  padding: 8px;
  background: #e6e6e6;
  border-radius: 4px;
  font-size: 12px;
  color: #666;
}

.day-rank.has-rank {
  background: #ffe999;
  color: #333;
  font-weight: 500;
  font-size: 12px;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .day-label {
    color: #999;
  }

  .day-rank {
    background: #333;
    color: #999;
  }

  .day-rank.has-rank {
    background: #ffe999;
    color: #333;
  }
}

.popup-overlay {
  position: fixed;
  inset: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  padding: 16px;
}

.popup-content {
  background: white;
  border-radius: 16px;
  width: 90%;
  max-width: 480px;
  max-height: 90vh;
  overflow-y: auto;
  padding: 24px;
  box-shadow: 0 4px 24px rgba(0, 0, 0, 0.12);
}

.popup-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.close-btn {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #666;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  transition: all 0.2s ease;
}

.close-btn:hover {
  background: #f5f5f5;
}

.popup-body {
  color: #444;
}

.bullet {
  color: #666;
  font-size: 16px;
  line-height: 1;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .popup-content {
    background: #2a2a2a;
  }

  .popup-header h2 {
    color: #fff;
  }

  .close-btn {
    color: #999;
  }

  .close-btn:hover {
    background: #333;
  }

  .popup-body {
    color: #e6e6e6;
  }

  .popup-body p,
  .popup-body li {
    color: #e6e6e6;
  }

  .popup-body .border-t {
    border-color: #444;
  }
}

/* Add these responsive styles for screens smaller than Nest Hub */
@media (max-width: 1023px) {
  .game-layout {
    flex-direction: column-reverse;
    gap: 32px;
  }

  .left-section2,
  .right-section2 {
    width: 100%;
    padding: 16px;
  }

  .right-section2 {
    margin-bottom: 20px;
  }

  .current-word {
    font-size: 20px;
    min-height: 40px;
  }

  .hexagon-section {
    gap: 16px;
  }

  .hexagon-grid {
    transform: scale(0.9);
  }

  .controls {
    padding: 12px;
  }

  /* Adjust tabs section for mobile */
  .tabs-section {
    margin-top: 20px;
  }

  .tab-content {
    padding: 12px;
  }

  .stats-checkboxes {
    flex-wrap: wrap;
    gap: 12px;
  }

  /* Adjust progress dots for mobile */
  .progress-dots {
    justify-content: center;
    flex-wrap: wrap;
  }

  .dot {
    width: 20px;
    height: 20px;
    font-size: 12px;
  }

  /* Adjust word chips for mobile */
  .word-chips {
    gap: 6px;
  }

  .word-chip {
    padding: 3px 10px;
    font-size: 12px;
  }
}

/* Additional adjustments for very small screens */
@media (max-width: 480px) {
  .hexagon-grid {
    transform-origin: center;
    transform: scale(0.8);
    margin: 0 auto;
  }

  .hexagon-section {
    width: 100%;
    padding: 0;
  }

  .nav-bar {
    padding: 8px;
  }

  .mode-toggles {
    font-size: 12px;
  }

  .mode-btn {
    padding: 6px 10px;
  }

  .right-section-btn {
    font-size: 13px;
    height: 36px;
  }

  .level-header {
    flex-direction: column;
    align-items: flex-start;
  }

  .subtext {
    font-size: 12px;
  }
}

/* Add these new styles */
.points-popup {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: #28a745;
  color: white;
  padding: 16px 32px;
  border-radius: 8px;
  font-size: 24px;
  font-weight: 600;
  animation: fadeInOut 1.5s ease-in-out;
  z-index: 1100;
}

@keyframes fadeInOut {
  0% {
    opacity: 0;
    transform: translate(-50%, -50%) scale(0.8);
  }
  20% {
    opacity: 1;
    transform: translate(-50%, -50%) scale(1.1);
  }
  30% {
    transform: translate(-50%, -50%) scale(1);
  }
  80% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .points-popup {
    background: #28a745;
    color: white;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  }
}

/* Add validation popup styles */
.validation-popup {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: #dc3545;
  color: white;
  padding: 16px 32px;
  border-radius: 8px;
  font-size: 18px;
  font-weight: 500;
  animation: fadeInOut 1.5s ease-in-out;
  z-index: 1100;
  text-align: center;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .validation-popup {
    background: #dc3545;
    color: white;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  }
}

.user-profile-section {
  position: relative;
}

.profile-btn {
  background: none;
  border: none;
  padding: 0;
  cursor: pointer;
  position: relative;
  margin-left: 6px;
}

.avatar {
  width: 32px;
  height: 32px;
  background: #ffe999;
  color: #333;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 14px;
  transition: all 0.2s ease;
}

.profile-btn:hover .avatar {
  transform: scale(1.05);
}

.profile-dropdown {
  position: absolute;
  top: calc(100% + 8px);
  right: 0;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
  width: 240px;
  padding: 16px;
  z-index: 1000;
  animation: slideIn 0.2s ease;
}

.profile-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
}

.avatar-large {
  width: 30px;
  height: 30px;
  background: #ffe999;
  color: #333;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 14px;
}

.user-info {
  flex: 1;
}

.username {
  font-weight: 600;
  color: #333;
  font-size: 14px;
  margin-bottom: 4px;
}

.email {
  color: #666;
  font-size: 12px;
}

.dropdown-divider {
  height: 1px;
  background: #eee;
  margin: 8px -16px;
}

.dropdown-item {
  width: 100%;
  padding: 8px 16px;
  border: none;
  background: none;
  color: #666;
  font-size: 14px;
  text-align: left;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  border-radius: 6px;
  transition: all 0.2s ease;
}

.dropdown-item:hover {
  background: #f5f5f5;
  color: #333;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(-8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .profile-dropdown {
    background: #2a2a2a;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
  }

  .username {
    color: #fff;
  }

  .email {
    color: #999;
  }

  .dropdown-divider {
    background: #444;
  }

  .dropdown-item {
    color: #999;
  }

  .dropdown-item:hover {
    background: #333;
    color: #fff;
  }
}

/* Add loading spinner styles */
.loading-spinner {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 200px;
  color: #666;
}

/* Update circle styles */
.circle.green {
  background: #28a745;
  color: white;
}

.circle.center-letter {
  background: #28a745;
  color: white;
}

/* Add these new styles */
.rank-popup {
  max-width: 400px;
  width: 90%;
}

 

.rank-details .points {
  font-size: 14px;
  font-weight: 600;
  color: #333;
  margin-bottom: 10px;
}

.rank-details .description {
  font-size: 14px;
  color: #666;
  line-height: 1.5;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .rank-details .points {
    color: #fff;
  }
  
  .rank-details .description {
    color: #999;
  }
}

/* Add these new styles */
.rank-stats-popup {
  width: 95%;
  max-width: 600px;
  padding: 20px;
}

.popup-header {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 24px;
}

.back-btn {
  background: none;
  border: none;
  font-size: 20px;
  color: #666;
  cursor: pointer;
  padding: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.tab-buttons {
  display: flex;
  gap: 16px;
  flex: 1;
  justify-content: center;
}

.tab-buttons button {
  background: none;
  border: none;
  padding: 8px 16px;
  font-size: 14px;
  font-weight: 500;
  color: #666;
  cursor: pointer;
  border-radius: 20px;
}

.tab-buttons button.active {
  background: #ffe999;
  color: #333;
}

.letters-display {
  text-align: center;
  margin-bottom: 24px;
}

.letters-display span {
  font-size: 18px;
  font-weight: 500;
  letter-spacing: 2px;
}

.points-info {
  font-size: 14px;
  color: #666;
  margin-top: 8px;
}

.words-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(100px, 1fr));
  gap: 12px;
  padding: 16px;
}

.word-item {
  padding: 8px 12px;
  background: #f5f5f5;
  border-radius: 4px;
  font-size: 14px;
  text-align: center;
  color: #666;
}

.word-item.found {
  background: #28a745;
  color: white;
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .back-btn {
    color: #999;
  }

  .tab-buttons button {
    color: #999;
  }

  .tab-buttons button.active {
    background: #ffe999;
    color: #333;
  }

  .points-info {
    color: #999;
  }

  .word-item {
    background: #333;
    color: #999;
  }

  .word-item.found {
    background: #28a745;
    color: white;
  }
}
</style>



