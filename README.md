# landscape

blog等に流用可能なテキストファイルを生成します。  
既存のファイルを指定した場合はblog等に流用できるようにテキストファイルに情報を付与します。

## 動作

landscape hello.md

* hello.mdが存在する場合

hello.md のファイルの先頭にFront Matterを追加します。

* hello.mdが存在しない場合

hello.mdを作成します。 Front Matterは記述されている状態です。


## 引数

landscape hello.md gatsby

typeにgatsby等を指定すると、それに合わせたFront Matterを記述します。
指定しない場合はgatsby用のFront Matterが記述されます。

現在はgatsbyにしか対応していません。