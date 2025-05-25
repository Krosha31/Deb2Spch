async function handleLogin(event) {
    event.preventDefault(); 

    const login = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const errorMessageElement = document.getElementById('error-message'); 
    errorMessageElement.textContent = ""; 

    try {
        const response = await fetch('/login/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include', // чтобы refresh токен в HttpOnly cookie дошел от сервера
            body: JSON.stringify({ login, password }),
        });

        if (response.ok) {
            const data = await response.json();

            if (data.token) {
                localStorage.setItem('token', data.token); // access токен сохраняем
                window.location.href = '/';
            } else {
                errorMessageElement.textContent = 'Сервер не вернул access токен.';
            }
        } else if (response.status === 404 || response.status === 401) {
            errorMessageElement.textContent = "Неверный логин или пароль";
        } else {
            errorMessageElement.textContent = 'Ошибка авторизации. Попробуйте позже.';
        }
    } catch (error) {
        console.error('Ошибка:', error);
        errorMessageElement.textContent = 'Ошибка соединения с сервером.';
    }
}



