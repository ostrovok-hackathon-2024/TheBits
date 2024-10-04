<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated class="brand-colors-primary">
      <q-toolbar>
        <q-toolbar-title>
          The Bits. Хакатон Островок 2024. 1 трэк - Динамическое ценоообразование.
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>
defineOptions({
  name: "MainLayout",
});
</script>
