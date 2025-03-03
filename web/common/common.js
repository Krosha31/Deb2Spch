document.addEventListener("DOMContentLoaded", async function() {
    await fetch('/common/header.html')
        .then(response => response.text())
        .then(data => {
            document.getElementById('header-container').innerHTML = data;
        });
    const user = getCookie('user');

    if (user) {
        const loginItem = document.getElementById('login-item');
        console.log(loginItem)
        loginItem.innerHTML = '<div class="user-icon"><img src="/addons/user-icon.png" alt="User" class="user-icon"></div>';
        const userIcon = document.querySelector('.user-icon');
        userIcon.addEventListener('click', clearCookies);
    }
});


function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}


function clearCookies() {
    const cookies = document.cookie.split("; ");
    for (let cookie of cookies) {
        const eqPos = cookie.indexOf("=");
        const name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie;
        document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/";
    }
    window.location.href = '/';
}

