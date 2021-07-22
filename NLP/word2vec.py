from gensim.models import Word2Vec


# word2vec 할 문서집합과 유사도를 검색할 단어를 파라미터로 사용
def word2vec(corpus, word):
    nmodel = Word2Vec(sentences=corpus, vector_size=100, window=10, min_count=5, workers=10, sg=0)
    return nmodel.wv.most_similar(word)
