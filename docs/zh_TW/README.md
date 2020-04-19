# 台灣非官方 COVID-19 資料整理站 | 皮丘版

![](https://github.com/PichuChen/covid19/workflows/pichu%20master%20deploy/badge.svg)

[![台灣非官方 COVID-19 資料整理站 | 皮丘版](https://stopcovid19.pichuchen.tw/ogp.png)](https://stopcovid19.pichuchen.tw/)

### [日本語](./../../README.md) | [English](./../en/README.md) | [Español](./../es/README.md) | [한국어](./../ko/README.md) | 繁體中文 | [简体中文](./../zh_CN/README.md) | [Tiếng Việt](./../vi/README.md) | [ภาษาไทย](./../th/README.md) | [Français](./../fr/README.md) | [Português](./../pt_BR/README.md)

## 這個網站

這個網站主要是從東京都新型冠狀肺炎對策網站分支出來的台灣版網站，工作語言以中文以及英文為主。
原則上我個人鼓勵自己動手建立自己的版本，但是也很歡迎發送 PR 到這邊來改這個版本，或者是到 g0v 的 slack 上面討論可以新增什麼資料。

目前已經有其他同類型的資料整理站或是平台了，因此這個站主要定位的功能會在於可用性（對不同平台、語系、IPv6、載入速度）。資料更新的部分採用手動更新的方式以降低對 API 提供者的負擔。

## 如何貢獻
如果您能對 Issues 中做出各式各樣的修正協助，我們將不勝感激。

詳情請洽[如何貢獻](./CONTRIBUTING.md)。


## 行動原則
詳情請洽[建立網站的行動原則](./CODE_OF_CONDUCT.md)。

## 授權
本軟體採 [MIT 授權條款](./LICENSE.txt)釋出。

## 從這個網站衍生出來的東西

請參考[此連結](./FORKED_SITES.md)

## 給翻譯者的資訊

有要幫忙翻譯的話，請參考 [這個文件](./TRANSLATION.md) 。

## 給開發者的資訊

### 開發環境建置

- Node.js 版本最低需求：10.19.0 以上

**使用 yarn 的做法**
```bash
# install dependencies
$ yarn install

# serve with hot reload at localhost:3000
$ yarn dev
```

**取消可用性檢查（vue-axe）的方法**

- 如果本機開發伺服器負擔重時，以下方法可在關閉可用性檢查的狀況下進行啟動。

```bash
# serve with hot reload at localhost:3000
$ yarn dev-no-axe
```

**使用 docker compose 的做法**
```bash
# serve with hot reload at localhost:3000
$ docker-compose up --build
```

**使用 Vagrant 的做法**
```bash
# serve with hot reload at localhost:3000
$ vagrant up
```

### 被 `Cannot find module ****` 卡住時

**使用 yarn 的做法**
```
$ yarn install
```

**使用 docker compose 的做法**
```bash
$ docker-compose run --rm app yarn install
```

### VSCode + Remote Containers 的開發環境

1. 安裝 VSCode 的擴充套件「[Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack)」。
2. 如同 [這個圖像（外部連結）](https://code.visualstudio.com/docs/remote/containers#_quick-start-try-a-dev-container)點擊左下角的 「Open Folder in Container」 後選擇 Repository 的資料夾路徑開始建立環境。

#### 提示
- 如果想要變更設定，請更改 `.devcontainer/devcontainer.json` 這隻檔案。
詳細請參考 [devcontainer.json的參考值](https://code.visualstudio.com/docs/remote/containers#_devcontainerjson-reference)。
- Remote Container 啟動時擴充套件只有導入 「ESlint」、如果有必要，請在 `devcontainer.json` 的 `extensions` 中新增。
詳細的步驟請參考 [這裡（外部連結）](https://code.visualstudio.com/docs/remote/containers#_managing-extensions)。
- 如果要重新建立開發環境，請執行左下角的 「Rebuild Container」。

### 生產環境/其他環境的判定

關於 `process.env.GENERATE_ENV` 這個值，生產環境為 `'production'` ，除此之外為 `'development'` 。  
如果只想要在測試環境中執行的話，請利用這個值作為參考。

### Deploy 到 Staging 環境以及正式環境的方法

當 `pichu-master` 分支被更新時，HTML 檔案將會在 `pichu-master-pages` 分支中被 build 起來，然後正式版兼開發版網站 https://stopcovid19.pichuchen.tw/ 會被更新。

### 分支規則

只允許推送 Pull Request 到 `development` 跟 `dev-hotfix` 。
在推送 Pull Request 時，請依照以下命名規則為您的分支命名

新增功能: feature/#{ISSUE_ID}-#{branch_title_name}  
Hotfix: hotfix/#{ISSUE_ID}-{branch_title_name}

#### 基本分支
| 目的 | 分支 | 預覽用 URL | 備註 |
| ---- | -------- | ---- | ---- |
| 開發=正式 | pichu-master | https://stopcovid19.pichuchen.tw/ | 基本上請推送 Pull Request 到這裡 |

#### 系統所使用的分支
| 目的 | 分支 | 預覽用 URL | 備註 |
| ---- | -------- | ---- | ---- |
| 正式網站 HTML | pichu-master-pages | https://stopcovid19.pichuchen.tw/ | 生成靜態網站 HTML 的位置 |
| OGP 工作用 | deploy / new_ogp | 無 | OGP 更新用 |

#### 其他分支
其他分支是從東京都版本 Merge 回來或者是發送 PR 給東京都版本時用的。
