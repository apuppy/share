document.addEventListener('DOMContentLoaded', function() {
    let storedTheme = localStorage.getItem('theme');
    switch(storedTheme) {
        case 'dark':
            useTheme('dark')
            break
        case 'light':
            useTheme('dark')('light')
            break
        default:
            let sysTheme = querySysTheme()
            if (sysTheme === 'dark') {
                useTheme('dark')
            }else {
                useTheme('light')
            }
    }
})

function useTheme(theme) {
    //TODO use relevant theme
    console.log('Theme:', theme)
}

function querySysTheme() {
    const prefersDarkMode = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
    if (prefersDarkMode) {
        return 'dark'
    } else {
        return 'light'
    }
}