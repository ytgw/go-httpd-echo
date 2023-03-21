# go-httpd-echo

受信したHTTPリクエストの内容を表示するWEBアプリです。
[TCP Exposer](https://www.tcpexposer.com/)のデモや挙動確認のために作成しました。

## 表示内容

- メソッド・ホスト名・パスなどの基本情報
- TCP Exposer(リバースプロキシとして振る舞います)が付与したHTTPヘッダー
- その他HTTPヘッダー
- リクエストボディ


## アクセス方法

- [HTTPSでのアクセス](https://echo.tcpexposer.com/)

- [HTTPでのアクセス](http://echo.tcpexposer.com/)

- [TCPでのアクセス](http://tcpexposer.com:18080/)

    この場合は、TCP ExposerはHTTPリバースプロキシとして振る舞わず、TCP通信を転送します。

## ローカルでの動作方法

```bash
docker build -t go-httpd-echo .
docker run --rm -p 8080:8080 go-httpd-echo
curl localhost:8080
```
