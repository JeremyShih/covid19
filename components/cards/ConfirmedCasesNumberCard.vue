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
      本土: {},
      敦睦遠訓支隊: {}
    }
    const patientsLabels = []
    const firstPatient = Patients.data[0]
    const firstDateString = firstPatient['リリース日']
      .substr(-5, 5)
      .replace('-', '/')
    const today = new Date()
    const todayDateString = '' + (today.getMonth() + 1) + '/' + today.getDate()

    const firstDate = new Date(firstDateString)
    const todayDate = new Date(todayDateString)

    for (firstDate; ; firstDate.setDate(firstDate.getDate() + 1)) {
      if (firstDate <= todayDate) {
        break
      }
      const dateString =
        '' + (firstDate.getMonth() + 1) + '/' + firstDate.getDate()
      patientsData['境外'][dateString] = 0
      patientsData['本土'][dateString] = 0
      patientsData['敦睦遠訓支隊'][dateString] = 0
      patientsLabels.push(dateString)
    }

    Patients.data.map(x => {
      const oriDateString = x['リリース日'].substr(-5, 5).replace('-', '/')
      const date = new Date(oriDateString)
      const dateString = '' + (date.getMonth() + 1) + '/' + date.getDate()

      if (typeof patientsData['境外'][dateString] === 'undefined') {
        // init array for first time of this date.
        patientsData['境外'][dateString] = 0
        patientsData['本土'][dateString] = 0
        patientsData['敦睦遠訓支隊'][dateString] = 0
        patientsLabels.push(dateString)
      }

      patientsData[x.境外或本土][dateString] =
        patientsData[x.境外或本土][dateString] + 1
    })
    const patientsItems = [
      this.$t('境外'),
      this.$t('本土'),
      this.$t('敦睦遠訓支隊')
    ]
    const patientsDataLabels = [
      this.$t('境外'),
      this.$t('本土'),
      this.$t('敦睦遠訓支隊')
    ]

    const data = {
      Patients,
      patientsGraph: [
        Object.values(patientsData.境外),
        Object.values(patientsData.本土),
        Object.values(patientsData.敦睦遠訓支隊)
      ],
      patientsItems,
      patientsLabels,
      patientsDataLabels
    }
    return data
  }
}
</script>
