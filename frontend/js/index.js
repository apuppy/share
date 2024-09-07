const THEME_DARK = 'dark';
const THEME_LIGHT = 'light';


// manually toggle theme
document.querySelector('#theme-toggle').addEventListener('click', function (event) {
    event.preventDefault()
    let theme = ensureTheme()
    wantTheme = flipTheme(theme)
    showWantTheme(wantTheme)
    setStorageTheme(wantTheme)
    applyTheme(wantTheme)
})

// auto detect system theme when page loads
document.addEventListener('DOMContentLoaded', function (event) {
    let theme = ensureTheme()
    newTheme = flipTheme(theme)
    showWantTheme(newTheme)
    setStorageTheme(theme)
    applyTheme(theme)
})

function ensureTheme(){
    theme = getTheme()
    return theme
}

function getTheme() {
    let theme = ''
    let storageTheme = getStorageTheme();
    console.log("storageTheme:", storageTheme)
    switch (storageTheme) {
        case THEME_DARK:
            theme=THEME_DARK
            break
        case THEME_LIGHT:
            theme = THEME_LIGHT
            break
        default:
            let sysTheme = querySysTheme()
            if (sysTheme === THEME_DARK) {
                theme = THEME_DARK
            } else {
                theme = THEME_LIGHT
            }
    }
    return theme
}

function flipTheme(theme) {
    return theme === THEME_DARK ? THEME_LIGHT : THEME_DARK
}

function showWantTheme(wantTheme) {
    const themeEl = document.querySelector("#theme-name")
    console.log("wantTheme:", wantTheme)
    themeEl.textContent = wantTheme
}

function applyTheme(theme) {
    console.log('apply relevant theme:', theme)
    // document.body.classList.toggle("dark-theme")
    if (theme === THEME_DARK) {
        document.body.classList.add("dark-theme")
    } else {
        document.body.classList.remove("dark-theme")
    }
}

function querySysTheme() {
    const prefersDarkMode = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
    if (prefersDarkMode) {
        return THEME_DARK
    } else {
        return THEME_LIGHT
    }
}

function getStorageTheme() {
    return localStorage.getItem('theme');
}

function setStorageTheme(theme) {
    localStorage.setItem('theme', theme);
}