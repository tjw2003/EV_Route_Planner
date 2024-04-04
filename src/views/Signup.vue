<template>
    <div class="container">
        <div class="login-wrapper">
            <!-- 头部提示 -->
            <div class="header">用户注册</div>
  
            <!--表单包裹-->
            <div class="form-wrapper">
                <!--用户名-->
                <input type="text" v-model="username" placeholder="请输入用户名" class="input-item">
                <!--密码-->
                <input type="password" v-model="password" placeholder="请输入密码" class="input-item">
                <button type="button" class="btn" @click="register">注册</button>
  
            </div>
            <div class="msg">
                已有账号？<a href="../login">立即登录</a>
            </div>
        </div>
    </div>
  </template>
  <script>
  import axios from 'axios'
  
  export default{
    name:'SignUp',
    props:{
        msg:String
    },
    data(){
        return{
            username:"",
            password:"",
        }
    },
  
    methods:{
        register() {
              try {
                  const response = await axios.post('/register', { username: this.username, password: this.password });
                  console.log(response.data);
                  // 如果注册成功，进行页面重定向到/login
                this.$router.push({ path: '/login' });
              } catch (error) {
                  console.error(error);
              }
          }
      }
  }
  </script>
  <style scoped lang="scss">
  /*容器样式*/
  .container {
    height: 100vh;
    width: 100vw;
    background-image: linear-gradient(to right, #fbc2eb, #a6c1ee);
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
  /* 表单样式 */
  .login-wrapper {
    background-color: #fff;
    width: 250px;
    padding: 30px 20px;
    border-radius: 15px;
    box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
  }
  
  /*标题样式*/
  .login-wrapper .header{
    font-size: 30px;
    font-weight: bold;
    color: black;
    text-align: center;
    line-height: 200px;
  }
  
  /*输入框样式 */
  .login-wrapper .form-wrapper .input-item{
    display: block;
    width:100%;
    margin-bottom: 20px;
    border: none;
    padding: 10px;
    border-bottom: 2px solid rgb(128, 125, 125);
    font-size: 15px;
    outline: none;
  }
  
  .login-wrapper .form-wrapper .input-item::placeholder{
    text-transform: uppercase;
  }
  
  /*登录按钮*/
  .login-wrapper .form-wrapper .btn{
    display: inline-block;
    text-align: center;
    width: 100%;
    border: none;
    padding: 5px;
    margin: 40px auto 0;
    background-image: linear-gradient(to right,#fbc2eb,#a6c1ee);
    color: #fff;
    border-radius: 5px;
  }
  
  /*鼠标悬停*/
  .login-wrapper .form-wrapper .btn:hover{
    cursor: pointer;
    color: rgba(13, 36, 36, 0.76);
  }
  
  /*底部提示文字颜色*/
  .login-wrapper .msg{
    text-align: center;
    color: black;
    font-weight: bold;
    line-height: 80px;
  }
  
  .login-wrapper .msg a{
    text-decoration: none;
    color: #a6c1ee;
  }
  </style>