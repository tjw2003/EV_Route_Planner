# -*- coding: utf-8 -*-

from flask import Flask, request, jsonify
from flask_cors import CORS

app = Flask(__name__)
CORS(app, resources={r"/*": {"origins": "*"}})  # 允许所有来源的请求访问你的应用

# 预定义的APIKEY
PREDEFINED_APIKEY = "DBUICFAOIUDFBASDBNFAI"

# 定义权重
weights = [3, 2, 1, 1, 2, 1, 3, 2, 1, 3]  # 根据问题重要性设定的权重

# 计算总分函数
def calculate_score(answers):
    total_score = sum(answer * weight for answer, weight in zip(answers, weights))
    return total_score

# 分类用户风险态度
def classify_risk(total_score):
    if 20 <= total_score <= 45:
        return 1  # 激进型
    elif 46 <= total_score <= 70:
        return 2  # 中立型
    else:
        return 3  # 保守型

@app.route('/classify', methods=['POST'])
def classify_user():
    data = request.json
    api_key = data.get('APIKEY')
    user_id = data.get('USERID')
    answers = data.get('answers')
    
    # APIKEY验证逻辑
    if api_key != PREDEFINED_APIKEY:
        response = jsonify({"error": "Invalid APIKEY"})
        response.status_code =401
    if not all([user_id, answers]):
        response = jsonify({"error": "Missing USERID or answers"})
        response.status_code =400
    else:
        total_score = calculate_score(answers)
        risk_type = classify_risk(total_score)
        response = jsonify({"USERID": user_id, "RiskType": risk_type})
    
     # 设置 CORS 头部信息
    response.headers['Access-Control-Allow-Origin']='*'
    response.headers['Access-Control-Allow-Headers'] ='*'
    response.headers['Access-Control-Allow-Methods']= 'POST'

    return response

if __name__ == '__main__':
    app.run(debug=True)
