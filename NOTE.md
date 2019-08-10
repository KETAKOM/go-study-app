## 色々メモ
  
### Goのimport宣言  
```
import (
  "fmt"
  . "fmt" // Println(.は使ってるのみたことない、、、)
  _ "fmt" // 利用なし
  f "fmt" // f.Println 
)
```

```_``` はそのファイルでは使わないけど、他のパッケージで使用する場合にかく  

参考URL:
https://qiita.com/taizo/items/56f8639260eb2a0fa999  

### Wireを使用したDIについて
``` wire gen ./... ``` 
でDIされたコードを自動生成してくれる 
  
生成元のコード(このリポジトリだとinjecter.go)はビルド実行時には含まれない

``` // +build wireinject ``` 
の次は一行、改行が必要。

