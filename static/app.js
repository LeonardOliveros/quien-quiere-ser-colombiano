// Quiz App - Main JavaScript

// Global variables
let currentUser = null;
let currentSession = null;
let currentQuestion = null;
let questions = [];
let questionIndex = 0;
let timer = null;
let timeRemaining = 0;
let startTime = null;
let correctAnswers = 0;
let incorrectAnswers = 0;
let flaggedQuestions = [];
let gameMode = null;
let lifelines = {
    fifty: true,
    hint: true,
    skip: 3
};

// API Base URL
const API_URL = '/api';

// Helper function to add auth header to fetch requests
async function fetchWithAuth(url, options = {}) {
    const token = localStorage.getItem('token');

    if (!options.headers) {
        options.headers = {};
    }

    if (token) {
        options.headers['Authorization'] = token;
    }

    const response = await fetch(url, options);

    // Handle 401 Unauthorized - token is invalid
    if (response.status === 401) {
        // Clear invalid token from localStorage
        localStorage.removeItem('token');
        localStorage.removeItem('userId');
        currentUser = null;

        // Show alert and redirect to login
        alert('Sesión expirada o inválida. Por favor, inicia sesión de nuevo.');
        location.reload();
    }

    return response;
}

// Initialize app
document.addEventListener('DOMContentLoaded', () => {
    checkAuth();
    setupEventListeners();
});

// Setup event listeners
function setupEventListeners() {
    // Auth forms - Check if elements exist before adding listeners
    const loginForm = document.getElementById('loginForm');
    const registerForm = document.getElementById('registerForm');
    
    if (loginForm) {
        loginForm.addEventListener('submit', handleLogin);
    }
    if (registerForm) {
        registerForm.addEventListener('submit', handleRegister);
    }
}

// Authentication functions
function checkAuth() {
    const userId = localStorage.getItem('userId');
    const token = localStorage.getItem('token');
    
    if (userId && token) {
        currentUser = { id: userId, token: token };
        showMainMenu();
    } else {
        showAuthModal();
    }
}

function showAuthModal() {
    const authModal = document.getElementById('authModal');
    if (authModal) {
        const modal = new bootstrap.Modal(authModal);
        modal.show();
    }
}

async function handleLogin(e) {
    e.preventDefault();
    
    const username = document.getElementById('loginUsername').value;
    const password = document.getElementById('loginPassword').value;
    
    try {
        const response = await fetch(`${API_URL}/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password })
        });
        
        const data = await response.json();
        
        if (response.ok) {
            localStorage.setItem('userId', data.user_id);
            localStorage.setItem('token', data.token);
            currentUser = { id: data.user_id, token: data.token };
            
            // Close modal
            const authModal = bootstrap.Modal.getInstance(document.getElementById('authModal'));
            if (authModal) {
                authModal.hide();
            }
            
            showMainMenu();
        } else {
            alert(data.error || 'Error al iniciar sesión');
        }
    } catch (error) {
        console.error('Login error:', error);
        alert('Error de conexión');
    }
}

async function handleRegister(e) {
    e.preventDefault();

    const username = document.getElementById('regUsername').value;
    const email = document.getElementById('regEmail').value;
    const password = document.getElementById('regPassword').value;

    try {
        const response = await fetch(`${API_URL}/register`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, email, password })
        });

        const data = await response.json();

        if (response.ok) {
            // Automatically login after successful registration
            await autoLoginAfterRegister(username, password);
        } else {
            alert(data.error || 'Error al registrarse');
        }
    } catch (error) {
        console.error('Register error:', error);
        alert('Error de conexión');
    }
}

async function autoLoginAfterRegister(username, password) {
    try {
        const response = await fetch(`${API_URL}/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password })
        });

        const data = await response.json();

        if (response.ok) {
            localStorage.setItem('userId', data.user_id);
            localStorage.setItem('token', data.token);
            currentUser = { id: data.user_id, token: data.token };

            // Clear registration form
            const registerForm = document.getElementById('registerForm');
            if (registerForm) {
                registerForm.reset();
            }

            // Close modal
            const authModal = bootstrap.Modal.getInstance(document.getElementById('authModal'));
            if (authModal) {
                authModal.hide();
            }

            showMainMenu();
        } else {
            alert(data.error || 'Error al iniciar sesión automático');
            // Switch to login tab so user can try manually
            const loginTab = document.querySelector('#authTabs .nav-link[href="#login"]');
            if (loginTab) {
                loginTab.click();
            }
        }
    } catch (error) {
        console.error('Auto-login error:', error);
        alert('Error de conexión. Por favor, inicia sesión manualmente.');
    }
}

// Navigation functions
function showMainMenu() {
    hideAllScreens();
    const mainMenu = document.getElementById('mainMenu');
    if (mainMenu) {
        mainMenu.style.display = 'block';
    }
}

function showGameScreen() {
    hideAllScreens();
    const gameScreen = document.getElementById('gameScreen');
    if (gameScreen) {
        gameScreen.style.display = 'block';
    }
}

function showResultsScreen() {
    hideAllScreens();
    const resultsScreen = document.getElementById('resultsScreen');
    if (resultsScreen) {
        resultsScreen.style.display = 'block';
    }
}

function hideAllScreens() {
    const screens = ['mainMenu', 'gameScreen', 'resultsScreen'];
    screens.forEach(screenId => {
        const screen = document.getElementById(screenId);
        if (screen) {
            screen.style.display = 'none';
        }
    });
}

// Game functions
async function startGame(mode) {
    gameMode = mode;

    // Determine game configuration
    let config = {
        mode: mode,
        question_count: 1000, // Use all available questions
        time_limit: 0,
        categories: ['CULTURA', 'GEOGRAFIA', 'HISTORIA', 'CONSTITUCION'],
        difficulty: 'MIXED',
        focus_weak_areas: false
    };

    switch (mode) {
        case 'TIMED':
            config.question_count = 1000; // Use all available questions
            config.time_limit = 10800; // 3 hours in seconds
            break;
        case 'WEAK_AREAS':
            config.focus_weak_areas = true;
            config.question_count = 1000; // Use all available questions
            break;
        case 'CATEGORY':
            const category = await selectCategory();
            if (!category) return;
            config.categories = [category];
            config.question_count = 1000; // Use all available questions
            break;
    }
    
    try {
        const response = await fetchWithAuth(`${API_URL}/game/start`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(config)
        });
        
        const data = await response.json();
        
        if (response.ok) {
            currentSession = data.session_id;
            timeRemaining = config.time_limit;
            resetGameState();
            showGameScreen();
            loadNextQuestion();
            // Initialize timer display
            if (timeRemaining > 0) {
                startTimer();
            } else {
                // No time limit - show placeholder
                const timerEl = document.getElementById('timer');
                if (timerEl) {
                    timerEl.textContent = '--:--';
                }
            }
        }
    } catch (error) {
        console.error('Start game error:', error);
        alert('Error al iniciar el juego');
    }
}

async function selectCategory() {
    return new Promise((resolve) => {
        const categories = ['CULTURA', 'GEOGRAFIA', 'HISTORIA', 'CONSTITUCION'];
        const categoryNames = {
            'CULTURA': 'Cultura y Sociedad',
            'GEOGRAFIA': 'Geografía',
            'HISTORIA': 'Historia Patria',
            'CONSTITUCION': 'Constitución Política'
        };
        
        let html = '<div class="category-selection">';
        categories.forEach(cat => {
            html += `<button class="btn btn-outline-light m-2" onclick="resolveCategory('${cat}')">${categoryNames[cat]}</button>`;
        });
        html += '</div>';
        
        // Show category selection modal
        const modalHtml = `
            <div class="modal fade" id="categoryModal">
                <div class="modal-dialog">
                    <div class="modal-content">
                        <div class="modal-header">
                            <h5 class="modal-title">Selecciona una Categoría</h5>
                        </div>
                        <div class="modal-body">
                            ${html}
                        </div>
                    </div>
                </div>
            </div>
        `;
        
        // Remove any existing category modal
        const existingModal = document.getElementById('categoryModal');
        if (existingModal) {
            existingModal.remove();
        }
        
        document.body.insertAdjacentHTML('beforeend', modalHtml);
        const modal = new bootstrap.Modal(document.getElementById('categoryModal'));
        modal.show();
        
        window.resolveCategory = (cat) => {
            modal.hide();
            setTimeout(() => {
                const modalElement = document.getElementById('categoryModal');
                if (modalElement) {
                    modalElement.remove();
                }
                resolve(cat);
            }, 300);
        };
    });
}

function resetGameState() {
    questionIndex = 0;
    correctAnswers = 0;
    incorrectAnswers = 0;
    flaggedQuestions = [];
    lifelines = { fifty: true, hint: true, skip: 3 };
    startTime = Date.now();
    updateScoreDisplay();
}

async function loadNextQuestion() {
    try {
        const response = await fetchWithAuth(`${API_URL}/game/${currentSession}/question`, {
            headers: {}
        });

        const data = await response.json();

        if (response.ok && data.question) {
            currentQuestion = data.question;
            questionIndex = data.question_number;
            displayQuestion();
        } else {
            // No more questions, end game
            endGame();
        }
    } catch (error) {
        console.error('Load question error:', error);
        alert('Error al cargar la pregunta');
    }
}

function displayQuestion() {
    // Update question counter
    const currentQuestionEl = document.getElementById('currentQuestion');
    const totalQuestionsEl = document.getElementById('totalQuestions');
    const categoryBadgeEl = document.getElementById('categoryBadge');
    const questionTextEl = document.getElementById('questionText');
    const progressBarEl = document.getElementById('progressBar');
    
    if (currentQuestionEl) {
        currentQuestionEl.textContent = questionIndex;
    }
    if (totalQuestionsEl) {
        totalQuestionsEl.textContent = getQuestionCount();
    }
    
    // Update category badge
    if (categoryBadgeEl && currentQuestion.category) {
        categoryBadgeEl.textContent = currentQuestion.category;
        categoryBadgeEl.className = `badge bg-${getCategoryColor(currentQuestion.category)}`;
    }
    
    // Display question text
    if (questionTextEl && currentQuestion.text) {
        questionTextEl.textContent = currentQuestion.text;
    }
    
    // Display choices
    if (currentQuestion.choices && Array.isArray(currentQuestion.choices)) {
        currentQuestion.choices.forEach((choice, index) => {
            const btn = document.getElementById(`answer-${index}`);
            if (btn) {
                const textEl = btn.querySelector('.answer-text');
                if (textEl) {
                    textEl.textContent = choice.text;
                }
                btn.className = 'answer-btn';
                btn.disabled = false;
                btn.style.display = 'block'; // Reset display after 50:50
            }
        });
    }
    
    // Update progress bar
    if (progressBarEl) {
        const progress = (questionIndex / getQuestionCount()) * 100;
        progressBarEl.style.width = progress + '%';
    }
}

async function selectAnswer(index) {
    if (!currentQuestion || currentQuestion.answered) return;
    
    const choice = currentQuestion.choices[index];
    if (!choice) return;
    
    const timeSpent = Math.floor((Date.now() - startTime) / 1000);
    
    try {
        const response = await fetchWithAuth(`${API_URL}/game/${currentSession}/answer`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                question_id: currentQuestion.id,
                choice_id: choice.id,
                time_spent: timeSpent
            })
        });
        
        const data = await response.json();
        
        if (response.ok) {
            currentQuestion.answered = true;
            showAnswerResult(index, data.correct, data.explanation);
            
            if (data.correct) {
                correctAnswers++;
            } else {
                incorrectAnswers++;
            }
            updateScoreDisplay();
            
            // Load next question after delay
            setTimeout(() => {
                loadNextQuestion();
            }, 2000);
        }
    } catch (error) {
        console.error('Submit answer error:', error);
        alert('Error al enviar respuesta');
    }
}

function showAnswerResult(selectedIndex, isCorrect, explanation) {
    const selectedBtn = document.getElementById(`answer-${selectedIndex}`);
    
    if (selectedBtn) {
        if (isCorrect) {
            selectedBtn.classList.add('correct');
        } else {
            selectedBtn.classList.add('incorrect');
            // Show correct answer
            if (currentQuestion.choices) {
                currentQuestion.choices.forEach((choice, index) => {
                    if (choice.is_correct) {
                        const correctBtn = document.getElementById(`answer-${index}`);
                        if (correctBtn) {
                            correctBtn.classList.add('correct');
                        }
                    }
                });
            }
        }
    }
    
    // Disable all buttons
    for (let i = 0; i < 4; i++) {
        const btn = document.getElementById(`answer-${i}`);
        if (btn) {
            btn.disabled = true;
        }
    }
}

async function flagQuestion(event) {
    if (!currentQuestion) return;

    const isFlagged = flaggedQuestions.includes(currentQuestion.id);

    try {
        await fetchWithAuth(`${API_URL}/game/${currentSession}/flag`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ question_id: currentQuestion.id })
        });

        // Find the flag button
        let flagBtn = null;
        if (event && event.target) {
            flagBtn = event.target.closest('button');
        } else {
            flagBtn = document.querySelector('button:has(i.fa-flag)');
        }

        if (isFlagged) {
            // Unmark the question
            flaggedQuestions = flaggedQuestions.filter(id => id !== currentQuestion.id);
            if (flagBtn) {
                flagBtn.classList.remove('btn-warning');
                flagBtn.innerHTML = '<i class="fas fa-flag"></i> Marcar';
            }
        } else {
            // Mark the question
            flaggedQuestions.push(currentQuestion.id);
            if (flagBtn) {
                flagBtn.classList.add('btn-warning');
                flagBtn.innerHTML = '<i class="fas fa-flag"></i> Marcada';
            }
        }

        // Update the flagged count
        const flaggedCountEl = document.getElementById('flaggedCount');
        if (flaggedCountEl) {
            flaggedCountEl.textContent = flaggedQuestions.length;
        }
    } catch (error) {
        console.error('Error flagging question:', error);
    }
}

// Lifeline functions
function use50_50() {
    if (!lifelines.fifty || !currentQuestion || currentQuestion.answered) return;
    
    lifelines.fifty = false;
    const lifelineBtn = document.getElementById('lifeline-50');
    if (lifelineBtn) {
        lifelineBtn.classList.add('used');
    }
    
    // Hide two incorrect answers
    let incorrectHidden = 0;
    if (currentQuestion.choices) {
        currentQuestion.choices.forEach((choice, index) => {
            if (!choice.is_correct && incorrectHidden < 2) {
                const btn = document.getElementById(`answer-${index}`);
                if (btn) {
                    btn.style.display = 'none';
                    incorrectHidden++;
                }
            }
        });
    }
}

function useHint() {
    if (!lifelines.hint || !currentQuestion || currentQuestion.answered) return;
    
    lifelines.hint = false;
    const lifelineBtn = document.getElementById('lifeline-hint');
    if (lifelineBtn) {
        lifelineBtn.classList.add('used');
    }
    
    if (currentQuestion.hint) {
        alert('Pista: ' + currentQuestion.hint);
    } else {
        alert('No hay pista disponible para esta pregunta');
    }
}

function skipQuestion() {
    if (lifelines.skip <= 0 || !currentQuestion || currentQuestion.answered) return;
    
    lifelines.skip--;
    const lifelineBtn = document.getElementById('lifeline-skip');
    if (lifelineBtn) {
        lifelineBtn.textContent = `Saltar (${lifelines.skip})`;
        if (lifelines.skip === 0) {
            lifelineBtn.classList.add('used');
        }
    }
    
    loadNextQuestion();
}

// Timer functions
function startTimer() {
    if (timer) {
        clearInterval(timer);
    }
    
    timer = setInterval(() => {
        timeRemaining--;
        updateTimerDisplay();
        
        if (timeRemaining <= 0) {
            clearInterval(timer);
            endGame();
        }
    }, 1000);
}

function updateTimerDisplay() {
    const hours = Math.floor(timeRemaining / 3600);
    const minutes = Math.floor((timeRemaining % 3600) / 60);
    const seconds = timeRemaining % 60;
    
    const display = hours > 0 ? 
        `${hours}:${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}` :
        `${minutes}:${seconds.toString().padStart(2, '0')}`;
    
    const timerEl = document.getElementById('timer');
    if (timerEl) {
        timerEl.textContent = display;
    }
}

function pauseGame() {
    // Disable all answer buttons to pause the game
    for (let i = 0; i < 4; i++) {
        const btn = document.getElementById(`answer-${i}`);
        if (btn) {
            btn.disabled = true;
        }
    }

    // Disable all lifeline buttons
    const lifelineButtons = document.querySelectorAll('.btn-lifeline');
    lifelineButtons.forEach(btn => {
        btn.disabled = true;
    });

    // Stop the timer if it's running
    if (timer) {
        clearInterval(timer);
        timer = null;
    }

    // Show pause modal and wait for user to continue
    const pauseModal = document.createElement('div');
    pauseModal.className = 'modal fade';
    pauseModal.id = 'pauseModal';
    pauseModal.setAttribute('tabindex', '-1');
    pauseModal.innerHTML = `
        <div class="modal-dialog modal-dialog-centered">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">Juego Pausado</h5>
                </div>
                <div class="modal-body">
                    <p>El juego está pausado. Presiona "Continuar" para reanudar.</p>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-primary" id="resumeBtn">Continuar</button>
                </div>
            </div>
        </div>
    `;

    document.body.appendChild(pauseModal);
    const modal = new bootstrap.Modal(pauseModal);
    modal.show();

    // Handle resume
    document.getElementById('resumeBtn').addEventListener('click', () => {
        modal.hide();
        setTimeout(() => {
            pauseModal.remove();
            resumeGame();
        }, 300);
    });
}

function resumeGame() {
    // Re-enable all answer buttons that haven't been answered yet
    for (let i = 0; i < 4; i++) {
        const btn = document.getElementById(`answer-${i}`);
        if (btn) {
            // Only keep disabled if the current question was already answered
            if (currentQuestion && !currentQuestion.answered) {
                btn.disabled = false;
            }
        }
    }

    // Re-enable all lifeline buttons that haven't been used
    const lifelineButtons = document.querySelectorAll('.btn-lifeline');
    lifelineButtons.forEach(btn => {
        if (!btn.classList.contains('used')) {
            btn.disabled = false;
        }
    });

    // Restart the timer if time limit exists
    if (timeRemaining > 0) {
        startTimer();
    }
}

// Score and stats functions
function updateScoreDisplay() {
    const correctCountEl = document.getElementById('correctCount');
    const incorrectCountEl = document.getElementById('incorrectCount');
    const flaggedCountEl = document.getElementById('flaggedCount');
    
    if (correctCountEl) {
        correctCountEl.textContent = correctAnswers;
    }
    if (incorrectCountEl) {
        incorrectCountEl.textContent = incorrectAnswers;
    }
    if (flaggedCountEl) {
        flaggedCountEl.textContent = flaggedQuestions.length;
    }
}

async function endGame() {
    if (timer) {
        clearInterval(timer);
    }

    try {
        // End the session
        await fetchWithAuth(`${API_URL}/game/${currentSession}/end`, {
            method: 'POST',
            headers: {}
        });

        // Get results
        const response = await fetchWithAuth(`${API_URL}/game/${currentSession}/results`, {
            headers: {}
        });
        
        const results = await response.json();
        displayResults(results);
    } catch (error) {
        console.error('End game error:', error);
        alert('Error al finalizar el juego');
    }
}

function displayResults(results) {
    showResultsScreen();
    
    // Display overall score
    const percentage = Math.round(results.percentage || 0);
    const finalScoreEl = document.getElementById('finalScore');
    if (finalScoreEl) {
        finalScoreEl.textContent = percentage + '%';
    }
    
    const statusElement = document.getElementById('resultStatus');
    if (statusElement) {
        if (percentage >= 60) {
            statusElement.textContent = '¡APROBADO!';
            statusElement.className = 'result-status passed';
        } else {
            statusElement.textContent = 'NO APROBADO';
            statusElement.className = 'result-status failed';
        }
    }
    
    // Display category scores
    const categoryResultsDiv = document.getElementById('categoryResults');
    if (categoryResultsDiv && results.category_scores) {
        categoryResultsDiv.innerHTML = '';
        
        for (const [category, score] of Object.entries(results.category_scores)) {
            const categoryHtml = `
                <div class="category-result">
                    <h5>${category}</h5>
                    <div class="category-score-bar">
                        <div class="category-score-fill" style="width: ${score.percentage}%">
                            ${score.correct_answers}/${score.total_questions} (${Math.round(score.percentage)}%)
                        </div>
                    </div>
                    <small class="${score.passed ? 'text-success' : 'text-danger'}">
                        ${score.passed ? 'Aprobado' : 'Necesita mejorar'}
                    </small>
                </div>
            `;
            categoryResultsDiv.innerHTML += categoryHtml;
        }
    }
    
    // Display incorrect answers
    const incorrectDiv = document.getElementById('incorrectAnswers');
    if (incorrectDiv) {
        incorrectDiv.innerHTML = '';
        
        if (results.incorrect_answers && results.incorrect_answers.length > 0) {
            results.incorrect_answers.forEach(item => {
                const reviewHtml = `
                    <div class="review-item">
                        <div class="review-question">${item.question.text}</div>
                        <div class="review-answer user-answer">
                            <i class="fas fa-times"></i> Tu respuesta: ${item.user_choice ? item.user_choice.text : 'No respondida'}
                        </div>
                        <div class="review-answer correct-answer">
                            <i class="fas fa-check"></i> Respuesta correcta: ${item.correct_choice.text}
                        </div>
                        ${item.explanation ? `<div class="explanation">${item.explanation}</div>` : ''}
                    </div>
                `;
                incorrectDiv.innerHTML += reviewHtml;
            });
        } else {
            incorrectDiv.innerHTML = '<p class="text-success">¡Todas las respuestas fueron correctas!</p>';
        }
    }
    
    // Display flagged questions
    const flaggedDiv = document.getElementById('flaggedQuestions');
    if (flaggedDiv) {
        flaggedDiv.innerHTML = '';
        
        if (results.flagged_questions && results.flagged_questions.length > 0) {
            results.flagged_questions.forEach(question => {
                const flaggedHtml = `
                    <div class="review-item flagged">
                        <div class="review-question">${question.text}</div>
                        <div class="explanation">Marcaste esta pregunta para revisar más tarde.</div>
                    </div>
                `;
                flaggedDiv.innerHTML += flaggedHtml;
            });
        } else {
            flaggedDiv.innerHTML = '<p>No marcaste ninguna pregunta para revisión.</p>';
        }
    }
    
    // Display recommendations
    const recommendationsDiv = document.getElementById('studyRecommendations');
    if (recommendationsDiv) {
        recommendationsDiv.innerHTML = '';
        
        if (results.recommendations && results.recommendations.length > 0) {
            results.recommendations.forEach(rec => {
                const recHtml = `
                    <div class="recommendation-item">
                        <h6><i class="fas fa-book-reader"></i> Recomendación</h6>
                        <p>${rec}</p>
                    </div>
                `;
                recommendationsDiv.innerHTML += recHtml;
            });
        }
    }
}

// Statistics functions
async function showStats() {
    try {
        const response = await fetchWithAuth(`${API_URL}/user/${currentUser.id}/stats`, {
            headers: {}
        });
        
        const stats = await response.json();
        displayStats(stats);
    } catch (error) {
        console.error('Show stats error:', error);
        alert('Error al cargar estadísticas');
    }
}

function displayStats(stats) {
    const statsContent = document.getElementById('statsContent');
    if (!statsContent) return;
    
    let html = `
        <div class="stats-overview">
            <h4>Resumen General</h4>
            <div class="row">
                <div class="col-md-3">
                    <div class="stat-card">
                        <i class="fas fa-gamepad fa-2x mb-2"></i>
                        <p class="stat-label">Total Juegos</p>
                        <p class="stat-value">${stats.total_games || 0}</p>
                    </div>
                </div>
                <div class="col-md-3">
                    <div class="stat-card">
                        <i class="fas fa-percentage fa-2x mb-2"></i>
                        <p class="stat-label">Promedio</p>
                        <p class="stat-value">${Math.round(stats.average_score || 0)}%</p>
                    </div>
                </div>
                <div class="col-md-3">
                    <div class="stat-card">
                        <i class="fas fa-trophy fa-2x mb-2"></i>
                        <p class="stat-label">Mejor Puntuación</p>
                        <p class="stat-value">${stats.best_score || 0}</p>
                    </div>
                </div>
                <div class="col-md-3">
                    <div class="stat-card">
                        <i class="fas fa-question-circle fa-2x mb-2"></i>
                        <p class="stat-label">Preguntas Totales</p>
                        <p class="stat-value">${stats.total_questions || 0}</p>
                    </div>
                </div>
            </div>
        </div>
        
        <div class="category-stats mt-4">
            <h4>Estadísticas por Categoría</h4>
            <div class="row">
    `;
    
    if (stats.category_stats) {
        for (const [category, catStats] of Object.entries(stats.category_stats)) {
            const improvement = catStats.improvement > 0 ? 
                `<span class="text-success">↑${catStats.improvement.toFixed(1)}%</span>` :
                catStats.improvement < 0 ?
                `<span class="text-danger">↓${Math.abs(catStats.improvement).toFixed(1)}%</span>` :
                '<span class="text-muted">--</span>';
            
            html += `
                <div class="col-md-6 mb-3">
                    <div class="category-stat-card">
                        <h5>${category}</h5>
                        <div class="progress mb-2">
                            <div class="progress-bar" style="width: ${catStats.average_percentage || 0}%">
                                ${Math.round(catStats.average_percentage || 0)}%
                            </div>
                        </div>
                        <small>
                            ${catStats.correct_answers || 0}/${catStats.total_questions || 0} correctas
                            ${improvement}
                        </small>
                    </div>
                </div>
            `;
        }
    }
    
    html += `
            </div>
        </div>
        
        <div class="areas-summary mt-4">
            <div class="row">
                <div class="col-md-6">
                    <h5 class="text-danger">Áreas Débiles</h5>
                    <ul class="list-unstyled">
    `;
    
    if (stats.weak_areas && stats.weak_areas.length > 0) {
        stats.weak_areas.forEach(area => {
            html += `<li><i class="fas fa-exclamation-triangle"></i> ${area}</li>`;
        });
    } else {
        html += `<li class="text-muted">No hay áreas débiles identificadas</li>`;
    }
    
    html += `
                    </ul>
                </div>
                <div class="col-md-6">
                    <h5 class="text-success">Áreas Fuertes</h5>
                    <ul class="list-unstyled">
    `;
    
    if (stats.strong_areas && stats.strong_areas.length > 0) {
        stats.strong_areas.forEach(area => {
            html += `<li><i class="fas fa-check-circle"></i> ${area}</li>`;
        });
    } else {
        html += `<li class="text-muted">Continúa practicando para fortalecer tus conocimientos</li>`;
    }
    
    html += `
                    </ul>
                </div>
            </div>
        </div>
    `;
    
    statsContent.innerHTML = html;
    
    // Show modal
    const statsModal = document.getElementById('statsModal');
    if (statsModal) {
        const modal = new bootstrap.Modal(statsModal);
        modal.show();
    }
}

async function showHistory() {
    try {
        const response = await fetchWithAuth(`${API_URL}/user/${currentUser.id}/history`, {
            headers: {}
        });
        
        const history = await response.json();
        displayHistory(history);
    } catch (error) {
        console.error('Show history error:', error);
        alert('Error al cargar historial');
    }
}

function displayHistory(sessions) {
    let html = '<h4>Historial de Juegos</h4><div class="history-list">';
    
    if (!sessions || sessions.length === 0) {
        html += '<p class="text-muted">No hay juegos registrados aún.</p>';
    } else {
        sessions.forEach(session => {
            const date = new Date(session.created_at).toLocaleDateString('es-CO');
            const percentage = session.total_questions > 0 ? 
                (session.correct_answers / session.total_questions * 100).toFixed(1) : 0;
            const statusClass = percentage >= 60 ? 'success' : 'danger';
            
            html += `
                <div class="history-item">
                    <div class="row align-items-center">
                        <div class="col-md-3">${date}</div>
                        <div class="col-md-3">${getModeLabel(session.mode)}</div>
                        <div class="col-md-3">
                            ${session.correct_answers}/${session.total_questions} correctas
                        </div>
                        <div class="col-md-3">
                            <span class="badge bg-${statusClass}">${percentage}%</span>
                        </div>
                    </div>
                </div>
            `;
        });
    }
    
    html += '</div>';
    
    // Show in modal
    showModal('Historial de Juegos', html);
}

async function showRecommendations() {
    try {
        const response = await fetchWithAuth(`${API_URL}/recommendations/${currentUser.id}`, {
            headers: {}
        });
        
        const recommendations = await response.json();
        displayRecommendations(recommendations);
    } catch (error) {
        console.error('Show recommendations error:', error);
        alert('Error al cargar recomendaciones');
    }
}

function displayRecommendations(recommendations) {
    let html = '<h4>Recomendaciones de Estudio Personalizadas</h4>';
    
    if (!recommendations || recommendations.length === 0) {
        html += '<p class="text-success">¡Excelente! No hay recomendaciones específicas. Continúa con tu práctica regular.</p>';
    } else {
        html += '<div class="recommendations-list">';
        
        recommendations.sort((a, b) => b.priority - a.priority);
        
        recommendations.forEach(rec => {
            const priorityClass = rec.priority >= 4 ? 'danger' : 
                                 rec.priority >= 3 ? 'warning' : 'info';
            
            html += `
                <div class="recommendation-card mb-3">
                    <div class="d-flex justify-content-between align-items-start">
                        <div>
                            <h5>${rec.category} - ${rec.subcategory}</h5>
                            <p>${rec.description}</p>
                            <small class="text-muted">Prioridad: 
                                <span class="badge bg-${priorityClass}">${rec.priority}/5</span>
                            </small>
                        </div>
                    </div>
                </div>
            `;
        });
        
        html += '</div>';
    }
    
    // Show in modal
    showModal('Recomendaciones de Estudio', html);
}

// Helper functions
function getQuestionCount() {
    // Return a large number since we want to use all available questions
    return 1000;
}

function getCategoryColor(category) {
    const colors = {
        'CULTURA': 'info',
        'GEOGRAFIA': 'success',
        'HISTORIA': 'warning',
        'CONSTITUCION': 'danger'
    };
    return colors[category] || 'secondary';
}

function getModeLabel(mode) {
    const labels = {
        'PRACTICE': 'Práctica',
        'TIMED': 'Contrarreloj',
        'WEAK_AREAS': 'Áreas Débiles',
        'CATEGORY': 'Por Categoría'
    };
    return labels[mode] || mode;
}

function showModal(title, content) {
    const modalHtml = `
        <div class="modal fade" id="dynamicModal">
            <div class="modal-dialog modal-lg">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title">${title}</h5>
                        <button type="button" class="btn-close" data-bs-dismiss="modal"></button>
                    </div>
                    <div class="modal-body">
                        ${content}
                    </div>
                </div>
            </div>
        </div>
    `;
    
    // Remove any existing modal
    const existing = document.getElementById('dynamicModal');
    if (existing) {
        existing.remove();
    }
    
    document.body.insertAdjacentHTML('beforeend', modalHtml);
    const modal = new bootstrap.Modal(document.getElementById('dynamicModal'));
    modal.show();
    
    // Clean up modal after hiding
    document.getElementById('dynamicModal').addEventListener('hidden.bs.modal', function () {
        this.remove();
    });
}

function backToMenu() {
    showMainMenu();
}

function startNewGame() {
    if (gameMode) {
        startGame(gameMode);
    }
}

// Logout function
function logout() {
    localStorage.removeItem('userId');
    localStorage.removeItem('token');
    currentUser = null;
    location.reload();
}

// Export functions for use in HTML onclick handlers
window.startGame = startGame;
window.selectAnswer = selectAnswer;
window.flagQuestion = flagQuestion;
window.use50_50 = use50_50;
window.useHint = useHint;
window.skipQuestion = skipQuestion;
window.pauseGame = pauseGame;
window.resumeGame = resumeGame;
window.showStats = showStats;
window.showHistory = showHistory;
window.showRecommendations = showRecommendations;
window.backToMenu = backToMenu;
window.startNewGame = startNewGame;
window.logout = logout;
