# go-sample-todo

参考サイト：
https://yyh-gl.github.io/tech-blog/blog/go_web_api/

# 参考サイトを元に作ってみた

### ドメイン層はシステムが扱う業務領域に関するコードを置く場所である。

### リポジトリはDBはKVSとのCRUD処理を記述する場所
#### ドメイン層には技術的関心ごとを持ち込まないというルールがある
#### ここではinterfaceのみ定義する
#### 実装は infra で行う