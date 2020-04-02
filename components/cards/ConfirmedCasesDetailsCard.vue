<template>
  <v-col cols="12" md="6" class="DataCard">
    <confirmed-cases-card
      :title="$t('検査陽性者の状況')"
      :title-id="'details-of-confirmed-cases'"
      :date="Patients.date"
    >
      <confirmed-cases-table
        :aria-label="$t('検査陽性者の状況')"
        v-bind="confirmedCases"
      />
    </confirmed-cases-card>
  </v-col>
</template>

<script>
import Data from '@/data/data.json'
import Patients from '@/data/patients.json'
import formatConfirmedCases from '@/utils/formatConfirmedCases'
import ConfirmedCasesCard from '@/components/ConfirmedCasesCard.vue'
import ConfirmedCasesTable from '@/components/ConfirmedCasesTableWithoutSymptom.vue'

export default {
  components: {
    ConfirmedCasesCard,
    ConfirmedCasesTable
  },
  data() {
    // 検査陽性者の状況
    const confirmedCases = formatConfirmedCases(Data.main_summary)

    const data = {
      Patients,
      confirmedCases
    }
    return data
  }
}
</script>
