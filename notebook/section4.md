- docker machine
  - 現在存在しているdocker-machineのリストを表示する
  ```
    docker-machine ls
  ```

  - docker-machineの作成
  ```
    #virtualboxをドライバーにdefaultという名前のdockerホストを作成する。
    docker-machine create --driver virtualbox default
  ```
  - 作成したdockerホストに接続する。
  ```
    docker-machine env default
  ```

  - dockerホストの内部を見る
  ```
    docker-machine ssh default
  ```

  - dockerホストのidを調べる
  ```
    docker-machine ip default
  ```

  - dockerホストを停止する
  ```
    docker-machine stop default
  ```
