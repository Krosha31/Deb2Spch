
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

async function handleLogin(event) {
    event.preventDefault(); // Предотвращаем отправку формы по умолчанию

    const login = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const errorMessageElement = document.getElementById('error-message'); // Получаем элемент для ошибок
    errorMessageElement.textContent = ""; // Очищаем предыдущее сообщение об ошибке

    // const res = validatePassword(password);
    // if (res) {
    //     errorMessageElement.textContent = res; // Устанавливаем текст ошибки
    //     return;
    // }
    const response = await fetch('/login/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ login, password }),
    }).then(response => {
        if (response.ok) {
            window.location.href = '/';
            document.cookie = `user=${login}; path=/`;
        } else if (response.status == 404) {
            errorMessageElement.textContent = "Incorrect login or password";
        } else {
            errorMessageElement.textContent = 'Ошибка регистрации. Попробуйте еще раз.';
        }
    })
    .catch(error => {
        console.error('Ошибка:', error);
    });;

    if (response.ok) {
        const data = await response.json();
        alert('Успешно авторизованы!');
        console.log(data);
    } else {
        errorMessageElement.textContent = 'Ошибка авторизации. Проверьте имя пользователя и пароль.'; 
    }
}



