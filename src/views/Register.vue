<!-- Register.vue -->
<template>
    <div class="container">
        <div class="register-container">
            <h2>Register</h2>
            <!-- 注册表单 -->
            <div class="form-group">
                <label for="newUsername">Username:</label>
                <input type="text" id="newUsername" v-model="newUsername" placeholder="Enter your new username">
            </div>
            <div class="form-group">
                <label for="newPassword">Password:</label>
                <input type="password" id="newPassword" v-model="newPassword" placeholder="Enter your new password">
            </div>
            <button class="login-button" @click="goToLogin">Login</button>
            <button class="register-button" @click="register">Register</button>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            newUsername: '',
            newPassword: ''
        };
    },
    methods: {
        register() {
            // 发送注册请求到后端
            axios.post('http://localhost:8888/register', {
                newUsername: this.newUsername,
                newPassword: this.newPassword
            })
                .then(response => {
                    alert(response.data.message);
                    this.$router.push('/login');
                })
                .catch(error => {
                    this.error = 'Failed to register user';
                    console.error(error);
                });
        },
        goToLogin() {
            this.$router.push('/login');
        }
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

.register-container {
    max-width: 400px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
    margin-top: 18%;
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
    background-color: #28a745;
    /* 使用绿色作为注册按钮颜色 */
    color: #fff;
    border: none;
    border-radius: 3px;
    cursor: pointer;
}

.login-button {
    background-color: transparent;
    border: 1px solid #28a745;
    color: #28a745;
    /* 注册按钮无色背景，蓝色边框和文字 */
    margin-right: 10px;
}

.register-button {
    background-color: #28a745;

}
</style>