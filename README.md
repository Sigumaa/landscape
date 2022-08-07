# landscape

blog等に流用可能なテキストファイルを生成します。  
既存のファイルを指定した場合はblog等に流用できるようにテキストファイルに情報を付与します。

## 動作

landscape hello.md

* hello.mdが存在する場合

hello.md のファイルの先頭にFront Matterを追加します。

* hello.mdが存在しない場合

hello.mdを作成します。 Front Matterは記述されている状態です。


### check

- [x] 既存のファイルに書き込む
- [ ] 存在しない場合ファイルを作成する