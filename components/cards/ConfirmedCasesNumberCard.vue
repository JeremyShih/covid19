<template>
  <v-col cols="12" md="6" class="DataCard">
    <time-stacked-bar-chart
      :title="$t('陽性患者数')"
      :title-id="'number-of-confirmed-cases'"
      :chart-id="'time-bar-chart-patients'"
      :chart-data="patientsGraph"
      :date="Data.patients.date"
      :items="patientsItems"
      :labels="patientsLabels"
      :unit="$t('人')"
      :data-labels="patientsDataLabels"
    />
    <!-- 件.tested = 検査数 -->
  </v-col>
</template>

<script>
import Data from '@/data/data.json'
import TimeStackedBarChart from '@/components/TimeStackedBarChart.vue'

export default {
  components: {
    TimeStackedBarChart
  },
  data() {
    const patientsData = {
      境外: {},
      本土: {}
    }
    const patientsLabels = []

    Data.patients.data.map(x => {
      const dateString = x['リリース日'].substr(-5, 5).replace('-', '/')

      patientsData.境外 = {
        [dateString]: 0,
        ...patientsData.境外
      }
      patientsData.本土 = {
        [dateString]: 0,
        ...patientsData.本土
      }

      patientsData[x.境外或本土] = {
        ...patientsData[x.境外或本土],
        [dateString]: patientsData[x.境外或本土][dateString] + 1
      }

      if (!patientsLabels.includes(dateString)) {
        patientsLabels.push(dateString)
      }
    })
    const patientsItems = [this.$t('境外'), this.$t('本土')]
    const patientsDataLabels = [this.$t('境外'), this.$t('本土')]
    patientsLabels.reverse()
    const data = {
      Data,
      patientsGraph: [
        Object.values(patientsData.境外).reverse(),
        Object.values(patientsData.本土).reverse()
      ],
      patientsItems,
      patientsLabels,
      patientsDataLabels
    }
    return data
  }
}
</script>
