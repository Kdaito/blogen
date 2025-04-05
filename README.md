# blogen

**blogen** は、Note API から記事を取得し、HTML にパースしてローカルファイルとして保存する CLI ツールです。  
記事の ID を指定することで、任意の記事を取得・保存できます。

---

## 🛠 インストール方法

### ✅ 1. `go install` を使う方法

以下のコマンドで `blogen` をインストールできます。

```sh
go install github.com/Kdaito/blogen@latest
```

### ✅ 2. リリースバイナリを使う方法

GitHub Releases からお使いの OS に対応するバイナリをダウンロードして実行してください。

---

## 📘 利用可能なコマンド

| コマンド | 説明 |
| - | - |
| `indiv` | 指定した ID の Note 記事を取得して HTML ファイルとして出力します |
| `help` | コマンドのヘルプを表示します |

### 🧪 `indiv`

指定した ID の Note 記事を取得して HTML ファイルとして出力します

```sh
blogen indiv -w ./output note_id_1 note_id_2
```

- idは複数指定することができます。
- idは一回の実行で最大50個まで指定することができます。

#### オプション

| フラグ | 説明 |
| - | - |
| `-w` / `--workspace` | HTMLファイルを出力するディレクトリを指定します。指定しない場合はカレントディレクトリに出力します。 |
| `-h` / `--help` | `indiv`コマンドのヘルプを表示します |

#### 実行例

```sh
# 単一記事を保存
blogen indiv -w ./html_output n123abc456
```

```sh
# 複数記事を一括保存
blogen indiv -w ./html_output n123abc456 n789def101
```

#### 💡 補足
- ディレクトリが存在しない場合はエラーになります。事前に作成してください。
- Note API の仕様変更により、取得できない場合があります。