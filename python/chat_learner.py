"""
Quick and Dirty. First try to predict the owner of a given peace of text.
I am sure, there's a lot of code and a huge amount of papers out there doing 
way better.
"""
from sklearn.ensemble import RandomForestClassifier
from nltk import ngrams as create_ngrams
import random

N = 4


def read_log(log_path: str):
    with open(log_path, 'r') as f:
        log = f.readlines()
    random.shuffle(log)
    data = []
    for l in log:
        if l is not None and l.strip() is not "":
            user, msg = l.split(":", 1)
            ngrams = ngram_from_msg(msg)
            for gram in ngrams:
                data.append((user, [hash(i) for i in gram]))
    return data


def ngram_from_msg(msg: str):
    lines = []
    ngrams = list(create_ngrams(msg.split(), N))
    for gram in ngrams:
        lines.append([hash(i) for i in gram])
    return lines


if __name__ == '__main__':
    log_path = "<path-to-log>"
    data = read_log(log_path)
    train = data[:(len(data)//3) * 2]
    predict = data[(len(data)//3) * 2:]
    train_ngrams = [i[1] for i in train]
    train_label = [i[0] for i in train]
    predict_ngrams = [i[1] for i in predict]
    predict_label = [i[0] for i in predict]

    randomForest = RandomForestClassifier()
    randomForest.fit(train_ngrams, train_label)
    score = randomForest.score(predict_ngrams, predict_label)
    print(score)







