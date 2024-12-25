function toggleTheme() {
    fetch('/toggle-theme', {
        method: 'POST',
        credentials: 'same-origin'
    }).then(() => {
        document.body.classList.toggle('dark-theme');
    });
}

// Check for saved theme preference
document.addEventListener('DOMContentLoaded', () => {
    const theme = document.cookie.split('; ').find(row => row.startsWith('theme='));
    if (theme && theme.split('=')[1] === 'dark') {
        document.body.classList.add('dark-theme');
    }
});