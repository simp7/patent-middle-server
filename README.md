# patent-middle-server

---

## 사양

---

- 리눅스(bash 쉘)
- go 1.6 이상(최신 버전을 다운받으면 이상 없음) -> `wget -q -O - https://git.io/vQhTU`
- mongoDB 4.0.1 이상(최신 버전을 다운받으면 이상 없음) -> https://docs.mongodb.com/manual/administration/install-on-linux/ 참조
- pip3 -> `sudo apt install pip3`
- python3 (리눅스에는 기본적으로 설치되어 있음)

## 설치

---

1. 깃허브에서 다운로드
2. 터미널에서 프로젝트를 다운로드한 폴더로 이동
3. `go build`
4. `sudo mv ./patent-middle-server /usr/local/bin/` (만약 해당 경로가 없는 경우 $PATH 내의 경로 중 아무 경로에 넣으면 됨)

## 실행

---

### 서버

1. mongoDB 실행
2. 터미널에서 `patent-middle-server` 실행

### 클라이언트

- 주소/KR/`검색어` : LDA 방식 검색
- 주소/EN/`검색어` : word2vec 방식 검색

## 설정

---

설정 파일: $HOME/patent-server/conf.yaml (첫 실행시 자동으로 생성)

### rest-server
    word-search: 단어 -> 출원번호 검색 URL
    claim-search: 출원번호 -> 청구항 검색 URL
    api-key: 위의 검색에 쓰일 api 키 (없는 경우 환경 변수 $KIPRIS 참조, 환경 변수에 두는 것을 권장)
    row: 단어 -> 출원번호 검색에서 한 페이지에 받을 항목 갯수(500이 최대, 높을 수록 api 호출 빈도 적어짐, 낮을 수록 속도 살짝 향상)

### database
    address: MongoDB 데이터베이스 URL
    database: 특허를 저장할 MongoDB 데이터베이스 이름
    collection: 특허를 저장할 MongoDB 콜렉션 이름

### port
해당 서버에서 열 포트 번호

## 기타

---

### 기본 쉘로 bash 쉘이 아닌 다른 쉘 사용시

$HOME/patent-server의 .sh 파일의 첫 줄을 `#!/{사용자 쉘 절대경로}`로 변경
ex) #!/bin/zsh