from gensim.models import Word2Vec
import sys
import json
import warnings


# word2vec 할 문서집합과 유사도를 검색할 단어를 파라미터로 사용
def word2vec(corpus, word):
    nmodel = Word2Vec(sentences=corpus, vector_size=100, window=10, min_count=5, workers=10, sg=0)
    return nmodel.wv.most_similar(word)


def main():
    data_path = sys.argv[1]
    amount = int(sys.argv[2])
    words = sys.argv[3:]

    terms = [str, float]*len(words)
    for i, word in range(words):
        terms[i] = word2vec("", words)

    json_data = [[""] * amount] * len(words)
    for index, word in enumerate(words):
        json_data[index] = [terms[i] for i in word.argsort()[: -amount - 1: -1]]

    print(json.dumps(json_data, ensure_ascii=False))

    return


if __name__ == '__main__':
    warnings.filterwarnings(action='ignore')
    main()
