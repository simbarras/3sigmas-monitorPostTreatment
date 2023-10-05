<template>
  <v-card title="Action" variant="tonal" style="margin-bottom: 2rem;">
    <v-card-subtitle>
      {{existAction.id}}
    </v-card-subtitle>
    <v-form>
      <v-container>
        <v-row>
          <v-col
            cols="12"
            md="4"
          >
            <v-combobox
              label="Bucket"
              :items="bucketList"
              v-model="existAction.bucketName"
            ></v-combobox>
          </v-col>

          <v-col
            cols="12"
            md="4"
          >
            <v-combobox
              label="Fonction"
              :items="equationList"
              v-model="existAction.equationName"
            ></v-combobox>
          </v-col>

          <v-col
            cols="12"
            md="2"
          >
            <v-checkbox label="Active" v-model="existAction.active"></v-checkbox>
          </v-col>

          <v-col
            cols="12"
            md="1"
          >
            <v-btn icon="mdi-upload" @click="actionStore.uploadAction(existAction)" >
            </v-btn>
          </v-col>
          <v-col
            cols="12"
            md="1"
          >
            <v-btn icon="mdi-delete-forever" @click="actionStore.deleteAction(existAction)">
            </v-btn>
          </v-col>
        </v-row>
        <v-row>
          <v-col
            cols="12"
            md="12"
          >
            <v-textarea label="Variables" v-model="existAction.listVariables"></v-textarea>
          </v-col>
        </v-row>
      </v-container>
    </v-form>
  </v-card>
</template>

<script lang="ts" setup>

import {defineProps} from "vue";
import {Action} from "@/plugins/data";
import {useActionStore} from "@/store/app";

const actionStore = useActionStore();

defineProps({
  existAction: {
    type: Object as Action,
    required: true
  },
  bucketList: {
    type: Array,
    required: true
  },
  equationList: {
    type: Array,
    required: true
  }
})

</script>
