##dockerファイルの書き方

FROM dockerイメージ名:タグ名

RUN apt-get -y update  && apt-get install -y fortunes

CMD /uer/games/fortune | cowsay

- apt-get
パッケージを取得してインストール/アップデートする
	- -y(オプション)
	インタラクティブ（ユーザーへの問い合わせ）に「yes」と答えます。

	- update(コマンド)
	サーバーからパッケージ・リストを入手する

	- install(コマンド)
	パッケージをインストールする
