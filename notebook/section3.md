- コンテナのライフサイクル
  - create
    - コンテナ作成から使用前まで
    ```
      docker create --name status-test -it alpine /bin/sh
    ```
  - running
    ```
      docker start status-test
    ```

  - pause(一時停止)
    ```
      docker pause status-test
    ```
    なお一時停止を解除するコマンドは以下の通り
    ```
      docker unpause status-test
    ```

  - exit(コンテナは終了したが残り続けている)
    ```
      docker stop status-test
    ```

- コンテナのシェルに接続するコマンド
  - docker attachを使用する場合
    抜ける場合はcontrol+p,control+qを順にクリックする。(この時コンテナは停止しない)
    exitで抜ける方法もある（これはコンテナを停止する）
  ```
    docker attach コンテナ名orコンテナID
  ```
  - docker execを使用する場合（こちらの方が安全）
    exitで抜けてもコンテナは停止しない。ß
  ```
    docker exec it コンテナ名orコンテナID/bin/bash
  ```

- dockerコンテナの動作確認
  - 以下のコマンドで動作しているコマンドがあるかを確認
    -a は詳細な情報を表示するためのオプション
  ```
    docker ps -a
  ```
