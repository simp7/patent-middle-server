import gensim
import warnings
import sys
import dataProcessing
import json


# word2vec 할 문서집합과 유사도를 검색할 단어, Topic 개수를 파라미터로 사용
def lda(corpus, topic_num):
    # 단어마다 고유번호를 매겨서, 어떤 단어인지 알 수 있는 사전을 만듬
    dictionary = gensim.corpora.Dictionary(corpus)
    text_corpus = [dictionary.doc2bow(text) for text in corpus]

    # 단어를 벡터화해서 LDA 분석
    lda_model = gensim.models.ldamodel.LdaModel(text_corpus, num_topics=topic_num, id2word=dictionary, passes=15)
    topic = lda_model.print_topics(num_words=5)
    return topic


def main():

    data_path = sys.argv[1]
    topic_num = sys.argv[2]
    words = sys.argv[3:]

    clear_name, clear_item = dataProcessing.do(data_path)
    topics = lda(clear_item, topic_num)

    print(json.dumps(topics, ensure_ascii=False))

    return


if __name__ == '__main__':
    warnings.filterwarnings(action='ignore')
    main()
