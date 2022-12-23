<script lang="ts">
  import type client from "$lib/api/client";
  import Snackbar, { Label, Actions } from "@smui/snackbar";
  import IconButton from "@smui/icon-button";
  import type { Readable } from "svelte/store";

  export let store: Readable<client.types.Response<any>>;

  let messages: string[] = [];
  let index = -1;
  let snackbar: Snackbar;
  function next() {
    if (!snackbar) return;
    (snackbar.close as Function)();

    if (index < messages.length - 1) {
      index++;
      (snackbar.open as Function)();
    }
  }
  $: {
    messages = [];
    index = -1;

    if ($store.success) messages.push($store.success);
    if ($store.errors) for (const err of $store.errors) messages.push(err);

    setTimeout(next);
  }
</script>

<Snackbar bind:this={snackbar}>
  <Label>{messages[index]}</Label>
  <Actions>
    <IconButton class="material-icons" title="Dismiss" on:click={next}
      >close</IconButton
    >
  </Actions>
</Snackbar>
