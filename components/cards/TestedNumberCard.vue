<template>
  <v-col cols="12" md="6" class="DataCard">
    <time-stacked-bar-chart
      :title="$t('検査実施件数')"
      :title-id="'number-of-tested'"
      :chart-id="'time-stacked-bar-chart-inspections'"
      :chart-data="inspectionsGraph"
      :date="Data.inspections_summary.date"
      :items="inspectionsItems"
      :labels="inspectionsLabels"
      :unit="$t('件.tested')"
      :data-labels="inspectionsDataLabels"
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
    // 検査実施日別状況
    const inspectionsGraph = [
      Data.inspections_summary.data['嚴重特殊傳染性肺炎通報'],
      Data.inspections_summary.data['居家檢疫送驗'],
      Data.inspections_summary.data['疑似新冠病毒感染送驗']
    ]
    const inspectionsItems = [
      this.$t('嚴重特殊傳染性肺炎通報'),
      this.$t('居家檢疫送驗'),
      this.$t('疑似新冠病毒感染送驗')
    ]
    const inspectionsLabels = Data.inspections_summary.labels
    const inspectionsDataLabels = [this.$t('都内'), this.$t('その他')]

    const data = {
      Data,
      inspectionsGraph,
      inspectionsItems,
      inspectionsLabels,
      inspectionsDataLabels
    }
    return data
  }
}
</script>
