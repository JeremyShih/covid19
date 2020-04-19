<template>
  <v-col cols="12" md="6" class="DataCard">
    <time-stacked-bar-chart
      :title="$t('陽性患者数')"
      :title-id="'number-of-confirmed-cases'"
      :chart-id="'time-bar-chart-patients'"
      :chart-data="patientsGraph"
      :date="Patients.date"
      :items="patientsItems"
      :labels="patientsLabels"
      :unit="$t('人')"
      :data-labels="patientsDataLabels"
      :url="
        'https://docs.google.com/spreadsheets/d/1I9EXxe-pWLhcLosakg5TPt98ERY6tdpJn1KngIGY7oY/edit#gid=1441264486'
      "
    />
  </v-col>
</template>

<script>
import Patients from '@/data/patients.json'
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

    Patients.data.map(x => {
      const dateString = x['リリース日'].substr(-5, 5).replace('-', '/')

      if (typeof patientsData['境外'][dateString] === 'undefined') {
        // init array for first time of this date.
        patientsData['境外'][dateString] = 0
        patientsData['本土'][dateString] = 0
        patientsLabels.push(dateString)
      }

      patientsData[x.境外或本土][dateString] =
        patientsData[x.境外或本土][dateString] + 1
    })
    const patientsItems = [this.$t('境外'), this.$t('本土')]
    const patientsDataLabels = [this.$t('境外'), this.$t('本土')]

    const data = {
      Patients,
      patientsGraph: [
        Object.values(patientsData.境外),
        Object.values(patientsData.本土)
      ],
      patientsItems,
      patientsLabels,
      patientsDataLabels
    }
    return data
  }
}
</script>
