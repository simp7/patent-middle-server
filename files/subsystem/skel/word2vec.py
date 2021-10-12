from gensim.models import Word2Vec
import sys
import json
import warnings
import dataProcessing
from konlpy.tag import Okt


# word2vec 할 문서집합과 유사도를 검색할 단어를 파라미터로 사용
def word2vec(corpus, word):
    n_model = Word2Vec(sentences=corpus, vector_size=100, window=10, min_count=5, workers=10, sg=0)
    return n_model.wv.most_similar(word)


def main():

    data_path = sys.argv[1]
    okt = Okt()

    name, item = dataProcessing.do(data_path)
    amount = int(sys.argv[2])
    words = sys.argv[3:]

    divided = set()
    for word in words:
        nouns = okt.nouns(word)
        for noun in nouns:
            divided.add(noun)

    keywords = list()
    for word in divided:
        try:
            result = word2vec(name, word)
            result.insert(0, (word, 1.0))
            keywords.append(result)
        except KeyError:
            pass

    json_data = list()

    for keyword in keywords:
        words = list()
        for word_bundle in keyword:
            words.append(word_bundle[0])
        json_data.append(words)

    print(json.dumps(json_data, ensure_ascii=False))

    return


if __name__ == '__main__':
    warnings.filterwarnings(action='ignore')
    main()
