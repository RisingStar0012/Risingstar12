<template>
  <BBModal
    :title="
      isCreatingTable
        ? $t('schema-editor.actions.create-table')
        : $t('schema-editor.actions.rename')
    "
    class="shadow-inner outline outline-gray-200"
    @close="dismissModal"
  >
    <div class="w-72">
      <p>{{ $t("schema-editor.table.name") }}</p>
      <BBTextField
        class="my-2 w-full"
        :required="true"
        :focus-on-mount="true"
        :value="state.tableName"
        @input="handleTableNameChange"
      />
    </div>
    <div class="w-full flex items-center justify-end mt-2 space-x-3 pr-1 pb-1">
      <button type="button" class="btn-normal" @click="dismissModal">
        {{ $t("common.cancel") }}
      </button>
      <button class="btn-primary" @click="handleConfirmButtonClick">
        {{ isCreatingTable ? $t("common.create") : $t("common.save") }}
      </button>
    </div>
  </BBModal>
</template>

<script lang="ts" setup>
import { computed, reactive } from "vue";
import { useI18n } from "vue-i18n";
import { generateUniqueTabId, useNotificationStore } from "@/store";
import {
  convertColumnMetadataToColumn,
  convertTableMetadataToTable,
} from "@/types";
import { ColumnMetadata, TableMetadata } from "@/types/proto/store/database";
import { Engine } from "@/types/proto/v1/common";
import { SchemaEditorTabType, useSchemaEditorContext } from "../common";

const tableNameFieldRegexp = /^\S+$/;

interface LocalState {
  tableName: string;
}

const props = defineProps<{
  schemaId: string;
  tableId?: string;
}>();

const emit = defineEmits<{
  (event: "close"): void;
}>();

const { t } = useI18n();
const { engine, editableSchemas, addTab } = useSchemaEditorContext();
const notificationStore = useNotificationStore();
const state = reactive<LocalState>({
  tableName: "",
});

const isCreatingTable = computed(() => {
  return props.tableId === undefined;
});

const handleTableNameChange = (event: Event) => {
  state.tableName = (event.target as HTMLInputElement).value;
};

const handleConfirmButtonClick = async () => {
  if (!tableNameFieldRegexp.test(state.tableName)) {
    notificationStore.pushNotification({
      module: "bytebase",
      style: "CRITICAL",
      title: t("schema-editor.message.invalid-table-name"),
    });
    return;
  }

  const schema = editableSchemas.value.find(
    (schema) => schema.id === props.schemaId
  );
  if (!schema) {
    notificationStore.pushNotification({
      module: "bytebase",
      style: "CRITICAL",
      title: t("schema-editor.message.schema-not-found"),
    });
    return;
  }
  const tableNameList = schema.tableList.map((table) => table.name);
  if (tableNameList.includes(state.tableName)) {
    notificationStore.pushNotification({
      module: "bytebase",
      style: "CRITICAL",
      title: t("schema-editor.message.duplicated-table-name"),
    });
    return;
  }

  if (isCreatingTable.value) {
    const table = TableMetadata.fromPartial({});
    table.name = state.tableName;
    const column = ColumnMetadata.fromPartial({});
    column.name = "id";
    if (engine.value === Engine.POSTGRES) {
      column.type = "INTEGER";
    } else {
      column.type = "INT";
    }
    column.comment = "ID";
    const columnEdit = convertColumnMetadataToColumn(column);
    columnEdit.status = "created";
    const tableEdit = convertTableMetadataToTable(table);
    tableEdit.status = "created";
    tableEdit.columnList.push(columnEdit);
    tableEdit.primaryKey.columnIdList.push(columnEdit.id);
    schema.tableList.push(tableEdit);
    addTab({
      id: generateUniqueTabId(),
      type: SchemaEditorTabType.TabForTable,
      schemaId: props.schemaId,
      tableId: tableEdit.id,
    });
  } else {
    const table = schema.tableList.find((table) => table.id === props.tableId);
    if (table) {
      table.name = state.tableName;
    }
  }
  dismissModal();
};

const dismissModal = () => {
  emit("close");
};
</script>