import json
import sys
import warnings
import dataProcessing

import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.decomposition import TruncatedSVD


# 데이터 전처리까지 같이 함 (원본데이터 필요)
def lsa(clear_item, topic_num):

    new_df = pd.DataFrame({'item': clear_item}).fillna("")

    # 알파벳 이외 문자 제거
    new_df['clean_doc'] = new_df['item'].str.replace("[^ㄱ-ㅎ|ㅏ-ㅣ|가-힣#]", " ")
    print(new_df)

    # 길이 3 이하 문자 제거
    new_df['clean_doc'] = new_df['clean_doc'].apply(lambda x: ' '.join([w for w in x.split()]))

    # tf-idf 벡터로 변환
    vectorizer = TfidfVectorizer(max_features=1000, max_df=0.5, smooth_idf=True)
    X = vectorizer.fit_transform(new_df['clean_doc'])

    # SVD 차원축소
    svd_model = TruncatedSVD(n_components=topic_num, algorithm='randomized', n_iter=100, random_state=12)
    svd_model.fit(X)

    terms = vectorizer.get_feature_names()
    components = svd_model.components_
    return terms, components


def main():

    data_path = sys.argv[1]
    amount = int(sys.argv[2])

    clear_name, clear_item = dataProcessing.do(data_path)
    terms, components = lsa(clear_item, amount)

    json_data = [[""]*amount]*amount
    for index, topic in enumerate(components):
        json_data[index] = [terms[i] for i in topic.argsort()[: -amount - 1: -1]]

    print(json.dumps(json_data, ensure_ascii=False))

    return


if __name__ == '__main__':
    warnings.filterwarnings(action='ignore')
    main()
