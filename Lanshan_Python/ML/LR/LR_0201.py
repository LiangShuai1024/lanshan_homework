import numpy as np
import pandas as pd
import matplotlib.pyplot as plt


data = pd.read_csv('../DataSet/Dry_Bean_Dataset.csv')
df = pd.DataFrame(data)



label = []
for i in df['Class'][0:3349]:
    if i == 'SEKER':
        label.append(0)
    else:
        label.append(1)
# 循环遍历df中的 'Class' 列前3349行数据
# 并进行判断，clas列中的字符为SEKER，则标签为0，其他为1，这里对数据进行也处理，把以其他字符表示的类别，变为0/1类别
x1 = df['MajorAxisLength'][0:3349]
x2 = df['MinorAxisLength'][0:3349]
# 取MajorAxisLength' 和 'MinorAxisLength' 列的前3349行数据，存进x1和x2中
train_data = list(zip(x1, x2, label))
# 使用 zip 函数将 'MajorAxisLength'、'MinorAxisLength' 和标签数据组合成一个列表。每个元素是一个包含两个特征和一个标签的元组，这样就形成了训练数据集。
# zip把（）中的元素对应合成

# 一些参数的初始化
class Logistic_Regression:
    def __init__(self, traindata, alpha=0.001, circle=1000, batchlength=40):
        self.traindata = traindata  # 训练数据集
        self.alpha = alpha  # 学习率
        self.circle = circle  # 学习次数
        self.batchlength = batchlength  # 把3349个数据分成多个部分，每个部分有batchlength个数据
        self.w = np.random.normal(size=(3, 1))  # 随机初始化参数w
        # np.random.normal从正态（高斯）分布中抽取随机样本
        # loc=, scale=, size=（x,y,z）
        # loc为分布的均值（中心）,scale为分布的标准差（宽度），size为输出的维度,会抽取x*y*z个样本
    def data_process(self):
        ''' 做随机梯度下降，打乱数据顺序，并把所有数据分成若干个batch '''
        np.random.shuffle(self.traindata)
        # np.random.shuffle对数据进行随机打乱，使每次迭代数据顺序不一，可增加模型的泛化能力。
        data = [self.traindata[i:i + self.batchlength]
                for i in range(0, len(self.traindata), self.batchlength)]
        # 这里是列表推导式，将打乱顺序后的数据集按照批次长度self.batchlength划分成若干个小批次，
        # self.traindata[i:i + self.batchlength]为切片操作，
        # 用于从训练数据集 self.traindata 中获取从索引 i 开始，到索引 i + self.batchlength 结束的一段数据。
        # 这样就得到了一个批次大小为 self.batchlength 的数据。
        # for从0开始，以elf.batchlength为步长，不超过len(self.traindata)
        # 每个小批次包含了一定数量的样本数据。这样做是为了在训练过程中能够批量地对数据进行处理，提高训练效率。
        return data

    def train1(self):
        '''根据损失函数（1）来进行梯度下降，这里采用随机梯度下降'''
        for i in range(self.circle):
            batches = self.data_process()  # 调用上述函数，批量获取数据，存入batches
            print('the {} epoch'.format(i))  # 程序运行时显示执行次数
            for batch in batches:
                d_w = np.zeros(shape=(3, 1))  # 初始化一个3*1的全零矩阵，用来累计w导数值
                for j in batch:  # 取batch中每一组数据
                    x0 = np.r_[j[0:2], 1]  # 把数据中指标取出，后面补1，np.r_用于将数组沿行方向链接
                    x = np.mat(x0).T  # 转化成列向量
                    y = j[2]  # 标签
                    dw = (self.sigmoid(self.w.T * x) - y)[0, 0] * x  # 获取参数的梯度
                    d_w += dw  # 梯度的累加
                self.w -= self.alpha * d_w / self.batchlength  # 根据累计梯度更新步长
        w = regr.w
        w1 = w[0, 0]
        w2 = w[1, 0]
        w3 = w[2, 0]
        x = np.arange(190, 500)
        y = -w1 * x / w2 - w3 / w2
        plt.plot(x, y)  # 绘制一条直线，这里用于绘制决策边界
        color = []
        for i in df['Class'][0:3349]:
            if i == 'SEKER':
                color.append('red')
            else:
                color.append('blue')
        plt.scatter(df['MajorAxisLength'][0:3349], df['MinorAxisLength'][0:3349], color=color)  # 绘制散点图
        plt.xlabel('MajorAxisLength')
        plt.ylabel('MinorAxisLength')
        plt.show()


    # 均方差似乎不佳，这里就不细写了

    def train2(self):
        '''用均方损失函数来进行梯度下降求解'''
        for i in range(self.circle):
            batches = self.data_process()
            print('the {} epoch'.format(i))  # 程序运行时显示执行次数
            for batch in batches:
                d_w = np.zeros(shape=(3, 1))  # 用来累计w导数值
                for j in batch:  # 取batch中每一组数据
                    x0 = np.r_[j[0:2], 1]  # 把数据中指标取出，后面补1
                    x = np.mat(x0).T  # 转化成列向量
                    y = j[2]  # 标签
                    dw = \
                    ((self.sigmoid(self.w.T * x) - y) * self.sigmoid(self.w.T * x) * (1 - self.sigmoid(self.w.T * x)))[
                        0, 0] * x
                    d_w += dw
                self.w -= self.alpha * d_w / self.batchlength

    def sigmoid(self, x):
        return 1 / (1 + np.exp(-x))

    def predict(self, x):
        '''测试新数据属于哪一类，x是2维列向量'''
        s = self.sigmoid(self.w.T * x)
        if s >= 0.5:
            return 1
        elif s < 0.5:
            return 0


if __name__ == '__main__':
    regr = Logistic_Regression(traindata=train_data)
    regr.train1()  # 采用1的方式进行训练
