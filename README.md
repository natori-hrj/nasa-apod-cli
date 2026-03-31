# nasa-apod-cli

NASA の Astronomy Picture of the Day (APOD) をターミナルから閲覧できる CLI ツール。

## 機能

- 今日の天文写真のタイトル・説明・画像URLを表示
- 指定した日付のAPODを取得 (`--date`)
- ランダムなAPODを取得 (`--random`)
- 説明文を日本語に翻訳して表示 (`--ja`)
- ターミナルにASCIIアートで画像を表示 (`--ascii`)

## セットアップ

### 1. APIキーの取得

[NASA API](https://api.nasa.gov/) から無料のAPIキーを取得してください。

### 2. インストール

```bash
go install github.com/natori/nasa-apod-cli@latest
```

または、リポジトリをクローンしてビルド:

```bash
git clone https://github.com/natori/nasa-apod-cli.git
cd nasa-apod-cli
go build -o apod .
```

### 3. APIキーの設定

環境変数で設定:

```bash
export NASA_API_KEY=your_api_key_here
```

または `.env` ファイルを作成:

```bash
cp .env.example .env
# .env を編集してAPIキーを設定
```

## 使い方

```bash
# 今日のAPODを表示
apod

# 特定の日付を指定
apod --date 2024-01-01

# ランダムなAPODを取得
apod --random

# 説明文を日本語で表示
apod --ja

# ASCIIアートで画像を表示
apod --ascii

# オプションは組み合わせ可能
apod --random --ja --ascii
```

### オプション一覧

| フラグ | 説明 |
|---|---|
| `--date YYYY-MM-DD` | 指定した日付のAPODを取得 |
| `--random` | ランダムなAPODを取得 |
| `--ja` | 説明文を日本語に翻訳 |
| `--ascii` | 画像をASCIIアートで表示 |
| `-h`, `--help` | ヘルプを表示 |

## 技術スタック

- Go
- [cobra](https://github.com/spf13/cobra) - CLIフレームワーク
- [godotenv](https://github.com/joho/godotenv) - .env ファイル読み込み
- [image2ascii](https://github.com/qeesung/image2ascii) - ASCIIアート変換

## ライセンス

MIT
