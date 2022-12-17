<script lang="ts">
    import Button3 from "$lib/interact/Button3.svelte";
    import Error from "$lib/util/Error.svelte";
    import { createEventDispatcher } from "svelte";

    export let name : string;
    export let btn : string;

    const submit = createEventDispatcher()
</script>

<form on:submit|preventDefault={()=>{}}>
    <h2>{name}</h2>
    <slot/>
    <Button3 on:click={()=>submit("submit")}>{btn}</Button3>
</form>

<style>
    * {
        grid-column: span 2;
    }
    form {
        display: grid;
        grid-gap: 1rem;
        text-align: justify;
    }
</style>