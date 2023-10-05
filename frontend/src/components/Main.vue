<template>
  <v-container class="fill-height">

    <v-responsive class="align-center text-center fill-height">
      <div v-for="action in actions" :key="action.id">
        <ActionItem :existAction="action" :bucketList="['bucket1', 'bucket2', 'bucket3']" :equationList="['Gauche', 'Droite', 'Fleche']"/>
      </div>
      <v-btn prepend-icon="mdi-plus-circle" variant="tonal" size="x-large" elevation="24" rounded="xl" style="margin-bottom: 5rem">
        Ajouter une action
      </v-btn>
    </v-responsive>
  </v-container>
</template>

<script lang="ts" setup>
import ActionItem from "@/components/ActionItem.vue";
import {useActionStore} from "@/store/app";
import {computed} from "vue";

const actionStore = useActionStore();

actionStore.fetchActions();

const actions = computed(() => actionStore.actionList);

</script>
