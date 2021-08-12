from gensim.models import Word2Vec
import sys
import json
import warnings
import dataProcessing


# word2vec 할 문서집합과 유사도를 검색할 단어를 파라미터로 사용
def word2vec(corpus, word):
    nmodel = Word2Vec(sentences=corpus, vector_size=100, window=10, min_count=5, workers=10, sg=0)
    return nmodel.wv.most_similar(word)


def main():

    datapath = sys.argv[1]

    name, item = dataProcessing.do(datapath)
    amount = int(sys.argv[2])
    words = sys.argv[3:]

    terms = list()
    for i, word in enumerate(words):
        terms.append(word2vec(name, word))

    # json_data = [[""] * amount] * len(words)
    # for index, word in enumerate(words):
    #     json_data.insert(index, terms[word])

    print(json.dumps(terms, ensure_ascii=False))

    return


if __name__ == '__main__':
    warnings.filterwarnings(action='ignore')
    main()
