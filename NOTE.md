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