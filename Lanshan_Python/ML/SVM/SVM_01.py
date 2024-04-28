# -*- coding: utf-8 -*-
# Time：2019/4/1 16:48
# versions：Python 3.6

__author__ = "fpZRobert"

"""
    支持向量机实战-乳腺癌检测
"""

import warnings
warnings.filterwarnings("ignore", category=FutureWarning, module="sklearn", lineno=196)

from sklearn.svm import SVC
import numpy as np
import matplotlib.pyplot as plt
from sklearn.model_selection import GridSearchCV
from sklearn.model_selection import ShuffleSplit
from sklearn.model_selection import learning_curve
from sklearn.datasets import load_breast_cancer
from sklearn.model_selection import train_test_split

"""
    加载数据
"""
# 利用sklearn库导入乳腺癌数据集
cancer = load_breast_cancer()
X = cancer.data  # 特征矩阵
y = cancer.target   # 目标矩阵

# 数据形状和类别分布
print("Shape of X: {0}; positive example: {1}; negative: {2}".format(X.shape, y[y==1].shape[0], y[y==0].shape[0]))  # 查看数据的形状和类别分布

"""
    拆分数据集
"""
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2)

"""
    训练模型: RBF核函数
"""


clf = SVC(C=1.0, kernel="rbf", gamma=0.1)
# svm初始化，c为正则化参数，1.0表中等，kernel设置核函数为rbf，gamma设置核函数的系数为0.1，定义了单个训练样本的影响范围，值越小表示影响范围越远

# 计算训练集和测试集的得分
clf.fit(X_train, y_train)
train_score = clf.score(X_train, y_train)
test_score = clf.score(X_test, y_test)
print("train score: {0}; test score: {1}".format(train_score, test_score))

"""
    绘制学习曲线
"""

# 绘制学习曲线函数的定义
def plot_learning_curve(plt, estimator, title, X, y, ylim=None, cv=None,
                        n_jobs=1, train_sizes=np.linspace(.1, 1.0, 5)):
    plt.title(title)
    if ylim is not None:
        plt.ylim(*ylim)
    plt.xlabel("Training examples")
    plt.ylabel("Score")
    train_sizes, train_scores, test_scores = learning_curve(
        estimator, X, y, cv=cv, n_jobs=n_jobs, train_sizes=train_sizes)
    train_scores_mean = np.mean(train_scores, axis=1)
    train_scores_std = np.std(train_scores, axis=1)
    test_scores_mean = np.mean(test_scores, axis=1)
    test_scores_std = np.std(test_scores, axis=1)
    plt.grid()

    # 绘制学习曲线及其标准差的填充区域
    plt.fill_between(train_sizes, train_scores_mean - train_scores_std,
                     train_scores_mean + train_scores_std, alpha=0.1,
                     color="r")
    plt.fill_between(train_sizes, test_scores_mean - test_scores_std,
                     test_scores_mean + test_scores_std, alpha=0.1, color="g")
    plt.plot(train_sizes, train_scores_mean, 'o--', color="r",
             label="Training score")
    plt.plot(train_sizes, test_scores_mean, 'o-', color="g",
             label="Cross-validation score")

    plt.legend(loc="best")
    return plt

# 使用ShuffleSplit进行交叉验证
# ShuffleSplit函数用于将样本集合随机“打散”后划分为训练集、测试集
# 类似于交叉验证
cv = ShuffleSplit(n_splits=10, test_size=0.2, random_state=0)
title = "Learning Curves for Gaussian Kernel"

# 绘制学习曲线
plt.figure(figsize=(10, 4), dpi=144)
plot_learning_curve(plt, SVC(C=1.0, kernel="rbf", gamma=0.01), title, X, y, ylim=(0.5, 1.01), cv=cv)
plt.show()

"""
    模型调优
"""

# 定义要搜索的gamma值范围
gammas = np.linspace(0, 0.0003, 30)
param_grid = {"gamma": gammas}

# 使用GridSearchCV进行模型调优
# 自动调参，只要把参数输进去，就能给出最优化的结果和参数。
# 但是这个方法适合于小数据集，一旦数据的量级上去了，很难得出结果。
clf = GridSearchCV(SVC(), param_grid, cv=5)
clf.fit(X, y)

print("best param: {0}\n best score: {1}".format(clf.best_params_, clf.best_score_))

"""
    训练模型: 二阶多项式核函数
"""
clf = SVC(C=1.0, kernel="poly", degree=2)
clf.fit(X_train, y_train)
train_score = clf.score(X_train, y_train)
test_score = clf.score(X_test, y_test)
print("train score: {0}; test score: {1}".format(train_score, test_score))

"""
    绘制学习曲线
"""
cv = ShuffleSplit(n_splits=5, test_size=0.2, random_state=0)
title = "Learning Curves with degree={0}"
degrees = [1, 2]
plt.figure(figsize=(12, 4), dpi=144)
for i in range(len(degrees)):
    plt.subplot(1, len(degrees), i+1)
    plot_learning_curve(plt, SVC(C=1.0, kernel="poly", degree=degrees[i]), title.format(degrees[i]), X, y, ylim=(0.8, 1.01), cv=cv)
plt.show()
