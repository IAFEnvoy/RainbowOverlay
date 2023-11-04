window.onload = _ => {
    let text = document.getElementById('text')
    window.runtime.WindowSetSize(text.offsetWidth + 40, text.offsetHeight + 40)
}