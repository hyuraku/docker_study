1. Hello world

  次の順にコマンドを入力  
  dockerコマンド　サブコマンド　イメージの指定

  下のコマンドはhello-worldというイメージを指定して起動(run)する
  ```
  docker run hello-world
  ```
  runのコマンドは以下の3つのコマンドを同時に実行している。
  - docker pull:イメージの取得
  - dicker create:コンテナの作成
  - docker start:コンテナの起動

2. Docker image  
   Dockerイメージとは
    - コンテナ実行に必要なファイルをまとめたファイルシステム(OSのライブラリ、webサーバーなど)
    - AUFSなどの特殊なファイルシステムでで使用されている。
    - イメージ上のデータはレイヤで構成され読み取り専用

3. Whalesay  
  次のコマンドを使うとクジラがセリフを言う画像がみれる。
    ```
    docker run docker/whalesay cowsay Hello!!
    ```
