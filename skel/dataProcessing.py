#!/usr/bin/python
# -*- coding: utf-8 -*-

import os
import pandas as pd
import json
from konlpy.tag import Okt


# 데이터 전처리
def do(datapath):
    # 데이터 업로드
    data = pd.read_csv(datapath, sep='\t', error_bad_lines=False)

    # 출원명과 청구항 데이터
    name = list(data['name'])
    item = list(data['item'])

    # konlpy를 이용해서 명사만 추출
    okt = Okt()
    name_noun = list(map(okt.nouns, name))
    item_noun = list()
    for i in item:
        if type(i) != str:
            item_noun += ''
        else:
            item_noun.append(okt.nouns(i))

    # 불용어(쓸모없는 단어) 불러오기
    stopwords = list()
    with open(os.environ.get('HOME') + '/patent-server/stopwords.json', encoding='utf-8') as file:
        stopwords = json.load(file)

    # 불용어 제거하는 함수
    def clear(word_tokenize):
        result = list()
        for w in word_tokenize:
            if w not in stopwords:
                result.append(w)
        return result

    # 불용어 제거
    clear_name = list(map(clear, name_noun))
    clear_item = list(map(clear, item_noun))
    return clear_name, clear_item
