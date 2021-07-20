import gensim
from gensim import corpora

#word2vec할 문서집합과 유사도를 검색할 단어, Topic 개수를 파라미터로 사용
def LDA(corpus, word, topicNum):
    #단어마다 고유번호를 매겨서, 어떤 단어인지 알 수 있는 사전을 만듬
    dictionary = gensim.corpora.Dictionary(corpus)
    textCorpus = [dictionary.doc2bow(text) for text in corpus]

    #단어를 벡터화해서 LDA분석
    ldamodel = gensim.models.ldamodel.LdaModel(textCorpus, num_topics=topicNum, id2word=dictionary, passes=15)
    return ldamodel.print_topics(num_words=5)