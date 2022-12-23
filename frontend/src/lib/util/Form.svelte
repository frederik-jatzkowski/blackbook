<script lang="ts">
  import { createEventDispatcher } from "svelte";

  import Button, { Label } from "@smui/button";

  export let name: string;
  export let btn: string;

  const submit = createEventDispatcher();
</script>

<form on:submit|preventDefault={() => {}}>
  <h2>{name}</h2>
  <slot />
  <Button on:click={() => submit("submit")}>
    <Label>{btn}</Label>
  </Button>
</form>

<style>
  * {
    grid-column: span 2;
  }
  form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    text-align: justify;
  }
</style>
