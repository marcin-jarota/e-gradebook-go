import { type App, type Plugin } from 'vue'

// Recursive type for nested translation options
type TranslateOptions = {
  [key: string]: string | TranslateOptions
}

const TranslatePlugin: Plugin = {
  install: (app: App, options: TranslateOptions) => {
    const translate = (key: string): string | undefined => {
      //@ts-ignore
      const value = key.split('.').reduce((o, i) => {
        if (o) return o[i]
      }, options)

      return typeof value === 'string' ? value : `missing value for key: ${key}`
    }
    app.config.globalProperties.$translate = translate
    app.provide('translate', translate)
  }
}

declare module '@vue/runtime-core' {
  //Bind to `this` keyword
  interface ComponentCustomProperties {
    $translate: (key: string) => string | undefined
  }
}

export default TranslatePlugin
