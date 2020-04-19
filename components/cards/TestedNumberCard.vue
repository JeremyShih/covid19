<template>
  <v-col cols="12" md="6" class="DataCard">
    <time-stacked-bar-chart
      :title="$t('検査実施件数')"
      :title-id="'number-of-tested'"
      :chart-id="'time-stacked-bar-chart-inspections'"
      :chart-data="inspectionsGraph"
      :date="InspectionsSummary.date"
      :items="inspectionsItems"
      :labels="inspectionsLabels"
      :unit="$t('件.tested')"
      :data-labels="inspectionsDataLabels"
      :url="
        'https://docs.google.com/spreadsheets/d/e/2PACX-1vRM7gTCUvuCqR3zdcLGccuGLv1s7dpDcQ-MeH_AZxnCXtW4iqVmEzUnDSKR7o8OiMLPMelEpxE7Pi4Q/pubhtml#'
      "
    >
      <!-- 件.tested = 検査数 -->
      <template v-if="$i18n.locale !== 'ja-basic'" v-slot:additionalNotes />
    </time-stacked-bar-chart>
  </v-col>
</template>

<script>
import InspectionsSummary from '@/data/inspections_summary.json'
import TimeStackedBarChart from '@/components/TimeStackedBarChart.vue'

export default {
  components: {
    TimeStackedBarChart
  },
  data() {
    // 検査実施日別状況
    const inspectionsGraph = [
      InspectionsSummary.data['嚴重特殊傳染性肺炎通報'],
      InspectionsSummary.data['居家檢疫送驗'],
      InspectionsSummary.data['疑似新冠病毒感染送驗']
    ]
    const inspectionsItems = [
      this.$t('嚴重特殊傳染性肺炎通報'),
      this.$t('居家檢疫送驗'),
      this.$t('疑似新冠病毒感染送驗')
    ]
    const inspectionsLabels = InspectionsSummary.labels
    const inspectionsDataLabels = inspectionsItems

    const data = {
      InspectionsSummary,
      inspectionsGraph,
      inspectionsItems,
      inspectionsLabels,
      inspectionsDataLabels
    }
    return data
  }
}
</script>

<style module lang="scss">
.Graph {
  &Desc {
    margin: 0;
    margin-top: 1rem;
    padding-left: 0 !important;
    font-size: 12px;
    list-style: none;
  }
}
</style>
