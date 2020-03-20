export default {
  strategy: 'prefix_except_default',
  detectBrowserLanguage: {
    useCookie: true,
    cookieKey: 'i18n_redirected'
  },
  defaultLocale: 'zh-tw',
  vueI18n: {
    fallbackLocale: 'zh-tw',
    formatFallbackMessages: true
  },
  // vueI18nLoader: true,
  lazy: true,
  langDir: 'assets/locales/',
  locales: [
    {
      code: 'zh-tw',
      name: '中文 (台灣)',
      iso: 'zh-TW',
      file: 'zh_TW.json',
      description: 'Chinese (Taiwan)'
    },
    {
      code: 'ja',
      name: '日本語',
      iso: 'ja-JP',
      file: 'ja.json',
      description: 'Japanese'
    },
    {
      code: 'en',
      name: 'English',
      iso: 'en-US',
      file: 'en.json',
      description: 'English'
    },
    {
      code: 'zh-cn',
      name: '简体中文',
      iso: 'zh-CN',
      file: 'zh_CN.json',
      description: 'Simplified Chinese'
    },
    {
      code: 'ko',
      name: '한국어',
      iso: 'ko-KR',
      file: 'ko.json',
      description: 'Korean'
    },
    // #1126, #872 (comment)
    // ポルトガル語は訳が揃っていないため非表示
    // {
    //   code: 'pt-BR',
    //   name: 'Portuguese',
    //   iso: 'pt-BR',
    //   file: 'pt_BR.json',
    //   description: 'Portuguese'
    // },
    {
      code: 'vi',
      name: 'Tiếng Việt',
      iso: 'vi-VI',
      file: 'vi.json',
      description: 'Vietnamese'
    },
    {
      code: 'th',
      name: 'ไทย',
      iso: 'th-TH',
      file: 'th.json',
      description: 'Thai'
    },
    {
      code: 'ja-basic',
      name: 'やさしい にほんご',
      iso: 'ja-JP',
      file: 'ja-Hira.json',
      description: 'Easy Japanese'
    }
  ]
}
