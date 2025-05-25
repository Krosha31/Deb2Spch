document.addEventListener("DOMContentLoaded", async function() {
    await fetch('/common/header.html')
        .then(response => response.text())
        .then(data => {
            document.getElementById('header-container').innerHTML = data;

        const accessToken = localStorage.getItem("token");
        console.log(accessToken, isTokenValid(accessToken))
        if (accessToken && isTokenValid(accessToken)) {
            const loginItem = document.getElementById('login-item');
            console.log(loginItem)
            loginItem.innerHTML = `
        <li class="nav-item dropdown">
                <a class="nav-link dropdown-toggle d-flex align-items-center" href="#" id="userDropdown" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                    <img src="/addons/user-icon.png" alt="User" class="rounded-circle" width="48" height="48" style="object-fit: cover; vertical-align: middle;">
                </a>
                <ul class="dropdown-menu dropdown-menu-end custom-dropdown-size" aria-labelledby="userDropdown">
                    <li><a class="dropdown-item" href="/profile">Профиль</a></li>
                    <li><hr class="dropdown-divider"></li>
                    <li><a class="dropdown-item text-danger" href="#" id="logoutBtn">Выйти</a></li>
                </ul>
            </li>
            `;
        }
        const logoutBtn = document.getElementById("logoutBtn");
        if (logoutBtn) {
            logoutBtn.addEventListener("click", function (e) {
                e.preventDefault();
                localStorage.removeItem("token");
                window.location.href = "/";
            });
        }
    });
});


function isTokenValid(token) {
    try {
      const payload = JSON.parse(atob(token.split(".")[1]));
      const exp = payload.exp;
      const now = Math.floor(Date.now() / 1000);
      return exp > now;
    } catch (e) {
      return false;
    }
  }


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



function validatePassword(password) {
    const minLength = 10;
    const hasUpperCase = /[A-Z]/.test(password);
    const hasLowerCase = /[a-z]/.test(password);
    const hasSpecialChar = /[!@#$%^&*(),.?":{}|<>]/.test(password);

    if (password.length < minLength) {
        return "Пароль должен содержать хотя бы 10 символов.";
    }
    if (!hasUpperCase) {
        return "Пароль должен содержать хотя бы одну заглавную букву.";
    }
    if (!hasLowerCase) {
        return "Пароль должен содержать хотя бы одну строчную букву.";
    }
    if (!hasSpecialChar) {
        return "Пароль должен содержать хотя бы один специальный символ.";
    }

    return "";
}
