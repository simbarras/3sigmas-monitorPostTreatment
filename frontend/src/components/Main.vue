<template>
  <v-container class="fill-height">

    <v-responsive class="align-center text-center fill-height">
      <div v-for="action in actions" :key="action.idUuid">
        <ActionItem :existAction="action" :bucketList="buckets" :equationList="equations"/>
      </div>
      <v-btn prepend-icon="mdi-plus-circle" variant="tonal" size="x-large" elevation="24" rounded="xl" style="margin-bottom: 5rem" @click="actionStore.preAddAction()">
        Ajouter une action
      </v-btn>
    </v-responsive>
  </v-container>
</template>

<script lang="ts" setup>
import ActionItem from "@/components/ActionItem.vue";
import {useActionStore, useBucketStore, useEquationStore} from "@/store/app";
import {computed} from "vue";

const actionStore = useActionStore();
actionStore.fetchActions();

const bucketStore = useBucketStore();
bucketStore.fetchBuckets();

const equationStore = useEquationStore();
equationStore.fetchEquations();

const actions = computed(() => actionStore.actionList);
const buckets = computed(() => bucketStore.bucketList);
const equations = computed(() => equationStore.equationList);

</script>
