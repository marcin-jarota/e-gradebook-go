import type { App } from 'vue'
import VueApexCharts from 'vue3-apexcharts'

export const setupCharts = (app: App<Element>) => {
  // app.use(VueApexCharts)
  app.component('apex-chart', VueApexCharts)
}
