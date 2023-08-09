<template>
  <BBModal
    :title="$t('task.check-result.title', { name: task.title })"
    class="!w-[56rem]"
    header-class="whitespace-pre-wrap break-all gap-x-1"
    @close="$emit('close')"
  >
    <div class="space-y-4">
      <PlanCheckBadgeBar
        :plan-check-run-list="planCheckRunList"
        :selected-type="selectedType"
        :task="task"
      />
      <TabFilter
        v-if="selectedPlanCheckRunUID"
        v-model:value="selectedPlanCheckRunUID"
        :items="tabItemList"
      />

      <PlanCheckDetail
        v-if="selectedPlanCheckRun"
        :plan-check-run="selectedPlanCheckRun"
        :task="task"
      />
    </div>
  </BBModal>
</template>

<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { first, orderBy } from "lodash-es";
import { useI18n } from "vue-i18n";

import {
  PlanCheckRun,
  PlanCheckRun_Type,
  Task,
} from "@/types/proto/v1/rollout_service";
import { TabFilter, TabFilterItem } from "@/components/v2";
import PlanCheckBadgeBar from "./PlanCheckBadgeBar.vue";
import PlanCheckDetail from "./PlanCheckDetail.vue";
import { humanizeDate } from "@/utils";

const props = defineProps<{
  planCheckRunList: PlanCheckRun[];
  selectedType: PlanCheckRun_Type;
  task: Task;
}>();

defineEmits<{
  (event: "close"): void;
}>();

const { t } = useI18n();

const selectedPlanCheckRunList = computed(() => {
  return orderBy(
    props.planCheckRunList.filter(
      (checkRun) => checkRun.type === props.selectedType
    ),
    (checkRun) => parseInt(checkRun.uid, 10),
    "desc"
  );
});

const selectedPlanCheckRunUID = ref(first(selectedPlanCheckRunList.value)?.uid);

const selectedPlanCheckRun = computed(() => {
  const uid = selectedPlanCheckRunUID.value;
  if (!uid) return undefined;
  return selectedPlanCheckRunList.value.find(
    (checkRun) => checkRun.uid === uid
  );
});

const tabItemList = computed(() => {
  return selectedPlanCheckRunList.value.map<TabFilterItem<string>>(
    (checkRun, i) => {
      const label =
        i === 0
          ? t("common.latest")
          : checkRun.createTime
          ? humanizeDate(checkRun.createTime)
          : `UID(${checkRun.uid})`;
      return {
        label,
        value: checkRun.uid,
      };
    }
  );
});

watch(selectedPlanCheckRunList, (list) => {
  selectedPlanCheckRunUID.value = first(list)?.uid;
});
</script>