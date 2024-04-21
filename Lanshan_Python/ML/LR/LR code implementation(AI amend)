import numpy as np
import matplotlib.pyplot as plt

# 设置随机种子，以便结果可复现
np.random.seed(42)

# 生成随机数据
# 两个特征的均值和方差
mean_1 = [2, 2]
cov_1 = [[2, 0], [0, 2]]
mean_2 = [-2, -2]
cov_2 = [[1, 0], [0, 1]]

# 生成类别1的样本
X1 = np.random.multivariate_normal(mean_1, cov_1, 50)
y1 = np.zeros(50)

# 生成类别2的样本
X2 = np.random.multivariate_normal(mean_2, cov_2, 50)
y2 = np.ones(50)

# 合并样本和标签
X = np.concatenate((X1, X2), axis=0)
y = np.concatenate((y1, y2))


# 定义激活函数 sigmoid
def sigmoid(x):
    return 1.0 / (1.0 + np.exp(-x))


# 定义LogisticRegression类
class LogisticRegression:
    def __init__(self, learning_rate=0.01, num_iterations=1000):
        self.learning_rate = learning_rate
        self.num_iterations = num_iterations
        self.weights = None
        self.bias = None

    def fit(self, X, y):
        num_samples, num_features = X.shape

        # 初始化权重和偏置
        self.weights = np.zeros(num_features)
        self.bias = 0

        # 梯度下降
        for _ in range(self.num_iterations):
            linear_model = np.dot(X, self.weights) + self.bias
            y_pred = sigmoid(linear_model)

            dw = (1 / num_samples) * np.dot(X.T, (y_pred - y))
            db = (1 / num_samples) * np.sum(y_pred - y)

            self.weights -= self.learning_rate * dw
            self.bias -= self.learning_rate * db

    def predict_prob(self, X):
        linear_model = np.dot(X, self.weights) + self.bias
        y_pred = sigmoid(linear_model)
        return y_pred

    def predict(self, X, threshold=0.5):
        y_pred_prob = self.predict_prob(X)
        y_pred = np.zeros_like(y_pred_prob)
        y_pred[y_pred_prob >= threshold] = 1
        return y_pred


# 创建 Logistic 回归模型
logreg = LogisticRegression()

# 训练模型
logreg.fit(X, y)

# 预测样本
X_new = np.array([[2.5, 2.5], [-6.0, -4.0]])
y_pred_prob = logreg.predict_prob(X_new)
y_pred = logreg.predict(X_new)

print("Predicted Probabilities:", y_pred_prob)
print("Predicted Labels:", y_pred)

# 绘制散点图
plt.scatter(X[:, 0], X[:, 1], c=y, cmap=plt.cm.Set1, edgecolor='k')
plt.xlabel('Feature 1')
plt.ylabel('Feature 2')
plt.title('Logistic Regression Dataset')
plt.show()
