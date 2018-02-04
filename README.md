# CryptocurrencyServer
サーバーサイド

# 環境
MySQL : 5.7.21
Beego : 1.9.2
Bee CLI : 1.9.1
GoVersion : 1.9.3


# ローカル開発環境の起動方法

## 依存パッケージのインストール
```bash
glide install
```

## DBのマイグレーション
```bash
bee migrate
```

## APIの実行
```bash
bee run
```

### glideの使い方（参考リンク）
https://liginc.co.jp/305623