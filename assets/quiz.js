(() => {
  function selectAnswer(button) {
    const quiz = button.closest('[data-quiz]');
    const feedback = quiz.querySelector('.feedback');
    const correct = button.dataset.correct === 'true';

    quiz.querySelectorAll('.answer').forEach((answer) => {
      answer.classList.remove('correct', 'incorrect');
      answer.setAttribute('aria-pressed', 'false');
    });

    button.classList.add(correct ? 'correct' : 'incorrect');
    button.setAttribute('aria-pressed', 'true');
    feedback.className = `feedback ${correct ? 'good' : 'bad'}`;
    feedback.textContent = correct ? button.dataset.right : button.dataset.wrong;
  }

  document.querySelectorAll('[data-quiz] .answer').forEach((button) => {
    button.addEventListener('click', () => selectAnswer(button));
  });
})();
