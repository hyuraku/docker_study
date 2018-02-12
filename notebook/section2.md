- イメージのプッシュ
docker hubにて作ったレポジトリに
作成したイメージをプッシュする

  - まず下のコマンドでdocker hubにログインする
  ```
  docker login
  ```
  - docker hubにおけるタグ付けルールを確認
  dockerID/imageName:tagName

  - 下のコマンドでpushする
  ```
  docker push dockerID/imageName:tagName
  ```

  - 下のコマンドでpullする(おさらい)
```
   docker pull dockerID/imageName:tagName
```

- nginxのコマンドを立ち上げるコマンド
  - -dはデタッチモードを表す。このモードを設定しないと同じターミナル内で他の作業ができない。
```
  docker run --name コンテナ名　-d -p　ホストのポート番号:コンテナ側のポート番号 イメージ名
```
  コマンド例
```
  docker run --name some_nginx -d -p 8080:80 nginx
```

- nginxのコマンドを立ち上げるコマンド(バインドマウントを使用するケース)
  - マウント
    コンテナの外にあるデータを、コンテナの中で利用できる状態にすること
  - バインドマウント
    Dockerホストのファイルやディレクトリを利用できるようにする機能
```
  docker run --name コンテナ名 -d -v ホストのディレクトリ:コンテナのマウントポイント:オプション -p ホストのポート番号:コンテナのポート番号 イメージ名
```

コマンド例
```
  docker run --name some-nginx -v /some/content:/usr/share/nginx/html:ro -d nginx
  #/some/content:コンテナにマウントするホスト側のディレクトリ
  #/usr/share/nginx/html:マウント先のコンテナのディレクトリ
  #ro:read only、読み取り専用のオプション
```

macの場合のコマンド
```
  -v /Users/ユーザー名/指定のフォルダ:/usr/share/nginx/html:ro
```

- dockerファイルをコピーするコマンド
　- ホストマシンのファイルをコンテナ内にコピーする場合
```
  docker cp ホスト上のコピー死体ファイルのパス コンテナ名orID:コピー先のパス
```
　- コンテナ内のファイルをホストマシンにこぴーする場合
```
  docker cp コンテナ名orID:コンテナ上のコピーしたいファイルのパス コピー先のパス
```
