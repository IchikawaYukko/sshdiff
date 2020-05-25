# sshdiff

[English](README.md)

sshdiff は ssh によるリモートコマンド実行結果の差分を表示するツールです。

サーバ移行作業に便利かもしれない！？

# インストール
1. release から実行ファイルをダウンロード
1. 実行ファイルを /usr/local/bin 等に置く
1. パーミッションを `rwxr-xr-x` (755) に設定する

# 使い方
## 簡単な使い方
```
$ sshdiff server1 server2 ls
```

`ls` を server1 と server2 上で実行して、出力の差分を表示。

## sshユーザ名を指定する場合
```
$ sshdiff yuriko@server1 yuriko@server2 ls
```

## 引数付きコマンドを実行する場合
```
$ sshdiff yuriko@server1 yuriko@server2 'rpm -ql|sort'
````

引数付きで実行する場合は ' ' " " などでクォート必須。

ポート番号(22) や 鍵(id_rsa, id_ecdsa, etc..) がデフォルト出ない場合は ~/.ssh/config に以下の様に書くことで指定可能。

```
Host server1
HostName server1.example.com
Port 8022
IdentityFile ~/.ssh/id_ecdsa_server1

Host server2
HostName server2.example.com
Port 1022
IdentityFile ~/.ssh/id_ecdsa_server2
```