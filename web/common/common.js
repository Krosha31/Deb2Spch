document.addEventListener("DOMContentLoaded", function() {
    fetch('/common/header.html')
        .then(response => response.text())
        .then(data => {
            document.getElementById('header-container').innerHTML = data;
        });
});