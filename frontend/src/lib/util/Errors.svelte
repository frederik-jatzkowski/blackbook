<script lang="ts">
  import type client from "./api/client";
  import Error from "./Error.svelte";
  import Success from "./Success.svelte";


    export let response: client.types.Response<any>;
</script>

{#if response.errors || response.success}
    {#if response.success}
        <Success>{response.success}</Success>
    {/if}
    {#if response.errors}
        {#each response.errors as err}
            <Error>{err}</Error>
        {/each}
    {/if}
{/if}