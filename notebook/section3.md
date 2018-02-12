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
