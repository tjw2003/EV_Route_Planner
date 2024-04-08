<template>
  <div class="questionnaire">
    <h2>新能源汽车补能行为机理挖掘调查问卷</h2>
    <p>您好，欢迎您填写本问卷</p>
    <p>'1非常不同意', '2比较不同意', '3无所谓', '4比较同意', '5非常同意'</p>
    <!-- 问题列表 -->
    <div v-for="(question, index) in questions" :key="index" class="question">
      <p class="question-text">{{ question.text }}</p>
      <!-- 根据问题类型展示不同的输入方式 -->
      <template v-if="question.type === 'singleChoice'">
        <div class="options">
          <label v-for="(option, optionIndex) in question.options" :key="optionIndex" class="option">
            <input type="radio" :id="'option_' + index + '_' + optionIndex" :value="option" v-model="responses[index]">
            {{ option }}
          </label>
        </div>
      </template>
      <template v-else-if="question.type === 'multipleChoice'">
        <div class="options">
          <label v-for="(option, optionIndex) in question.options" :key="optionIndex" class="option">
            <input type="checkbox" :id="'option_' + index + '_' + optionIndex" :value="option"
              v-model="responses[index]">
            {{ option }}
          </label>
        </div>
      </template>
      <template v-else-if="question.type === 'text'">
        <textarea v-model="responses[index]" placeholder="请输入您的回答"></textarea>
      </template>
    </div>



    <!-- 提交按钮 -->
    <button @click="submitQuestionnaire">提交问卷</button>
  </div>
</template>

<script>
import axios from 'axios';
export default {

  data() {
    return {
      questions: [
        {
          text: '您倾向于在新能源汽车电池电量较低的情况下才会充电，以充分利用电池[单选题]:',
          type: 'singleChoice',
          options: ['1', '2', '3', '4', '5']
        },
        {
          text: '您会在电池电量低于50%时就开始尝试寻找充电设备，或开始留意周边的补能设施 [单选题] ：',
          type: 'singleChoice',
          options: ['1', '2', '3', '4', '5']
        },
        {
          text: '您一旦决定前去充电，那么到达充电桩前我不会去做其他任何事情 [单选题]:',
          type: 'singleChoice',
          options: ['1', '2', '3', '4', '5']
        },
        {
          text: '您每到达一个新的地点，您会刻意查询周边可用的充电设施 [单选题]：',
          type: 'singleChoice',
          options: ['1', '2', '3', '4', '5']
        },
        {
          text: '在当前描述场景下，当电动车的里程表显示剩的余公里数与到目的地的公里数差不多时，您会感到焦躁不安 [单选题] ：',
          type: 'singleChoice',
          options: ['1', '2', '3', '4', '5']
        },
        {
          text: '在当前描述场景下，如果可能，您会尽量让电动车电量保持在（ ）以上 [单选题]  ：',
          type: 'singleChoice',
          options: ['1', '2', '3', '4', '5']
        },
        {
          text: '在当前描述场景下，当有合适的机会充电时，您总是选择充电 [单选题] ：',
          type: 'singleChoice',
          options: ['1', '2', '3', '4', '5']
        },
        {
          text: '在当前描述场景下，当您的电动汽车的表现续航里程与到目的地的剩余里程差不多时，（目的地有补能设施），您还会坚持前往目的地 [单选题]:',
          type: 'singleChoice',
          options: ['1', '2', '3', '4', '5']
        },
        {
          text: '在当前描述场景下，当您要去一个不太熟悉的地方，相较于平常，您会给电动车充更多的电 [单选题]：',
          type: 'singleChoice',
          options: ['1', '2', '3', '4', '5']
        },
        {
          text: '在当前描述场景下，当电动车剩余里程和下一段路程差不多时，您会选择不充电而直接出发 [单选题]  ：',
          type: 'singleChoice',
          options: ['1', '2', '3', '4', '5']
        }
      ],
      responses: []
    };
  },
  methods: {
    submitQuestionnaire() {
      // 收集用户的回答
      const answers = this.responses.map(response => parseInt(response, 10));
      console.log('用户的答案：', answers);
      // 构建请求数据
      const requestData = {
        "APIKEY": "DBUICFAOIUDFBASDBNFAI", // 固定的 APIKEY
        "USERID": "1", 
        "answers": answers
      };

      try {
        // 发送 POST 请求
        const response = axios.post('http://150.158.136.241:8901/classify', requestData, { withCredentials: true })
        // 处理成功响应
        if (response.status === 200) {
          const responseData = response.data;
          console.log('后端返回的数据：', responseData);

          if (responseData.RiskType !== undefined && responseData.USERID !== undefined) {
            // 根据返回的数据做相应处理
            console.log('风险类型:', responseData.RiskType);
            console.log('用户ID:', responseData.USERID);
          }

          // 提示用户问卷提交成功
          alert('问卷提交成功！');

          // 清空用户的回答
          this.responses = [];

          // 页面跳转到 Planner 页面
          this.$router.push('/Planner');
        } else {
          throw new Error('请求失败'); // 抛出错误
        }
      } catch (error) {
        console.error('请求失败:', error);
        // 处理请求失败的情况，例如提示用户提交失败
        alert('问卷提交失败，请重试！');
      }
    }

  }
};
</script>

<style scoped>
.questionnaire {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f0f0f0;
  border-radius: 8px;
  font-family: Arial, sans-serif;
}

.question {
  margin-bottom: 20px;
}

.question-text {
  font-weight: bold;
}

.options {
  display: flex;
  flex-direction: column;
}

.option {
  margin-bottom: 8px;
}

textarea {
  width: 100%;
  height: 100px;
  padding: 8px;
  border-radius: 4px;
  border: 1px solid #ccc;
}

button {
  padding: 10px 20px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  opacity: 0.8;
}
</style>