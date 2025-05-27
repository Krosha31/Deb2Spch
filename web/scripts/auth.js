async function handleLogin(event) {
    event.preventDefault(); 

    const login = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const errorMessageElement = document.getElementById('error-message'); 
    errorMessageElement.textContent = ""; 

    try {
        const response = await fetch('login/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include', 
            body: JSON.stringify({ login, password }),
        });

        if (response.ok) {
            const data = await response.json();

            if (data.token) {
                localStorage.setItem('token', data.token);
                window.location.href = '';
            } else {
                console.log('Сервер не вернул access токен.');
            }
        } else if (response.status === 404 || response.status === 401) {
            errorMessageElement.textContent = "Incorrect login or password";
        } else {
            errorMessageElement.textContent = 'Authorization failed. Try again later';
        }
    } catch (error) {
        console.error('Ошибка:', error);
        errorMessageElement.textContent = 'Internal Server Error';
    }
}



