import json
import sys
import warnings
import dataProcessing

import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.decomposition import TruncatedSVD


def lsa(clear_item, topic_num):

    new_df = pd.DataFrame({'item': clear_item}).fillna("")

    # 토큰화된 단어를 빈 칸으로 묶음
    detokenized_doc = list()
    for token in clear_item:
        detokenized_doc.append(' '.join(token))
    new_df['clean_doc'] = detokenized_doc

    # tf-idf 벡터로 변환
    vectorizer = TfidfVectorizer(max_features=1000, max_df=0.5, smooth_idf=True)
    X = vectorizer.fit_transform(new_df['clean_doc'])

    # SVD 차원축소
    svd_model = TruncatedSVD(n_components=topic_num, algorithm='randomized', n_iter=100, random_state=12)
    svd_model.fit(X)

    terms = vectorizer.get_feature_names()
    components = svd_model.components_

    json_data = list()
    for topic in components:
        json_data.append([terms[i] for i in topic.argsort()[: -topic_num - 1: -1]])

    return json_data


def main():

    data_path = sys.argv[1]
    amount = int(sys.argv[2])

    clear_name, clear_item = dataProcessing.do(data_path)
    json_data = lsa(clear_item, amount)

    print(json.dumps(json_data, ensure_ascii=False))

    return


if __name__ == '__main__':
    warnings.filterwarnings(action='ignore')
    main()
