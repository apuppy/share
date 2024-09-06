const THEME_DARK = 'dark';
const THEME_LIGHT = 'light';


// manually toggle theme
document.querySelector('#theme-toggle').addEventListener('click', function (event) {
    event.preventDefault()
    manuallyUseTheme()
})

// auto detect system theme when page loads
document.addEventListener('DOMContentLoaded', function (event) {
    const theme = autoDetectTheme()
    useTheme(theme)
})

function autoDetectTheme() {
    let theme = ''
    let storageTheme = getStorageTheme();
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

function manuallyUseTheme() {
    const theme = switchTheme()
    useTheme(theme)
}

function switchTheme() {
    const themeEl = document.querySelector("#theme-name")
    const themeName = themeEl.getAttribute('data-theme-name')
    wantTheme = themeName === THEME_DARK ? THEME_LIGHT : THEME_DARK
    setStorageTheme(wantTheme);
    themeEl.setAttribute('data-theme-name', wantTheme)
    themeEl.innerHTML = wantTheme
    return wantTheme
}

function useTheme(theme) {
    //TODO use relevant theme related CSS
    console.log('Use relevant theme:', theme)
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