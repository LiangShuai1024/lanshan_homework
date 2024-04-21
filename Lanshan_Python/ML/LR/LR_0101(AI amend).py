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
# np.random.multivariate_normal根据均值和协方差矩阵的情况生成一个多元正态分布矩阵,这里为一个二维的数组
# y1和y2分别为0和1标签，
# 合并样本和标签
X = np.concatenate((X1, X2), axis=0)
y = np.concatenate((y1, y2))
# np.concatenate实现对（）内数组的拼接，
# axis指定拼接的方向，默认是 axis = 0，也就是说对0轴的数组对象进行纵向的拼接（纵向的拼接沿着axis= 1方向）
# 定义激活函数 sigmoid
def sigmoid(x):
    return 1.0 / (1.0 + np.exp(-x))


# 定义LogisticRegression类
# 一些参数的初始化
class LogisticRegression:
    def __init__(self, learning_rate=0.01, num_iterations=1000):
        self.learning_rate = learning_rate  # 学习率
        self.num_iterations = num_iterations  # 学习次数
        self.weights = None
        self.bias = None

    def fit(self, X, y):
        num_samples, num_features = X.shape  # 获取训练数据 X 的形状，其中 num_samples 是样本数量，num_features 是特征数量。

        # 初始化权重和偏置
        self.weights = np.zeros(num_features)
        # np.zeros生成一个给定形状和类型的用0填充的数组，这里以初始化权重，
        # 因为每个特征都有一个对应的权重参数，这里取X，也就是特征数据的形状

        self.bias = 0  # 初始化模型的偏置参数为零

        # 梯度下降
        for _ in range(self.num_iterations):
            linear_model = np.dot(X, self.weights) + self.bias
            # 计算线性模型的预测结果，即将输入特征 X 与权重相乘，然后加上偏置。
            # np.dot实现向量点积与矩阵乘法
            # 若两个参数a和b都是一维向量则是计算的点积，
            # 但是当其中有一个是矩阵时（包括一维矩阵），dot便进行矩阵乘法运算。
            # 所以如果是一个向量和一个矩阵相乘，这个向量会自动转换为一维矩阵进行计算。
            y_pred = sigmoid(linear_model)
            # 将预测结果传入 Sigmoid 函数中进行转换，得到预测的类别概率

            dw = (1 / num_samples) * np.dot(X.T, (y_pred - y))
            # 计算权重参数的梯度。首先计算预测类别与真实类别的差异(y_pred - y)，然后乘以输入特征X的转置，最后除以样本数量。
            db = (1 / num_samples) * np.sum(y_pred - y)
            # 计算偏置参数的梯度。计算预测类别与真实类别的差异，然后np.sum求和，最后除以样本数量。

            self.weights -= self.learning_rate * dw
            self.bias -= self.learning_rate * db
            # 将上述dw与db乘以学习率，以更新模型权重与偏置

    def predict_prob(self, X):
        linear_model = np.dot(X, self.weights) + self.bias
        # 特征 X 与权重相乘，然后加上偏置。
        y_pred = sigmoid(linear_model)
        # 结果通过Sigmoid进行转换，得到预测的类别概率
        return y_pred

    def predict(self, X, threshold=0.5):
        y_pred_prob = self.predict_prob(X)
        # 用上面predict_prob取得预测的类别概率
        y_pred = np.zeros_like(y_pred_prob)
        # 初始化一个与y_pred_prob形状相同的全零数组，用于存储预测的类别
        # np.zeros_like返回相同形状和类型的一个全零数组
        y_pred[y_pred_prob >= threshold] = 1
        # 这里采用阈值分类的方法，当预测概率y_pred_prob大于等于阈值threshold时，预测为正类（1），否则预测为负类（0）。
        # 存入y_pred中
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
