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

4. Images  
ローカル上にあるイメージ一覧を表示するコマンド
imageとはイメージ一覧を表示するサブコマンドである。  
    ```
    docker images
    ```

5. Tag,remove  
  - タグ付けとは、既存のイメージに別名をつけることである。  
  docker タグ付けするサブコマンド　既存のイメージ名　新規のイメージ名:新規のタグ名  
      ```
      docker tag docker/whalesay my_whalesay:ver1
      ```  

  - 次のコマンドでイメージの詳細情報を表示する  
    docker イメージの詳細情報を表示するサブコマンド　対象のイメージ名orイメージID
      ```
      docker inspect my_whalesay
      ```

  - 次のコマンドでイメージを削除する  
    docker rmi 対象のイメージ名orイメージID
      ```
      docker rmi docker/whalesay
      ```
    強制削除の時は-fをつける
      ```
      docker rmi -f docker/whalesay
      ```
