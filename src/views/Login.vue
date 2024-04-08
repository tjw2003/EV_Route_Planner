<template>
    <div class="container">
        <div class="login-container">
            <h2>Login</h2>
            <div class="form-group">
                <label for="username">Username:</label>
                <input type="text" id="username" v-model="username" placeholder="Enter your username">
            </div>
            <div class="form-group">
                <label for="password">Password:</label>
                <input type="password" id="password" v-model="password" placeholder="Enter your password">
            </div>
            <button class="login-button" @click="login">Login</button>
            <button class="register-button" @click="goToRegister">Register</button>
            <p v-if="error" class="error-message">{{ error }}</p>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
export default {
    data() {
        return {
            username: '',
            password: '',
            error: ''
        };
    },
    methods: {
        login() {
            if (this.username === '' || this.password === '') {
                this.error = 'Please enter username and password.';
            } else {
                // 发送登录请求到后端
                axios.post('http://localhost:8888/login', {
                    username: this.username,
                    password: this.password
                })
                    .then(response => {
                        alert(response.data.message);
                        this.$router.push('/questions');
                    })
                    .catch(error => {
                        this.error = 'Invalid credentials';
                        console.error(error);
                    });
            }
        },
        goToRegister() {
            this.$router.push('/register');
        },
    }
};
</script>

<style scoped>
.container {
    position: absolute;
    width: 100%;
    height: 100%;
    /* 设置容器高度为视窗高度 */
    background-image: linear-gradient(to right, #fbc2eb, #ef0909eb);
}

.login-container {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 400px;
    /* 设置容器宽度为400px */
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
    background-color: #f9f9f9;
}

.form-group {
    margin-bottom: 15px;
}

label {
    font-weight: bold;
}

input[type="text"],
input[type="password"] {
    width: 100%;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 3px;
}

button {
    padding: 10px 15px;
    color: #fff;
    border: none;
    border-radius: 3px;
    cursor: pointer;
    margin-right: 30px;
}

.login-button {
    background-color: #007bff;
    /* 登录按钮蓝色背景 */
}

.register-button {
    background-color: transparent;
    border: 1px solid #007bff;
    color: #007bff;
    /* 注册按钮无色背景，蓝色边框和文字 */
    margin-right: 10px;
}


.error-message {
    color: red;
    margin-top: 10px;
}
</style>