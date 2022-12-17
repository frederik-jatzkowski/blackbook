<script lang="ts">
    import Button3 from "$lib/interact/Button3.svelte";
    import { createEventDispatcher } from "svelte";
    import type { TaskData } from "./types";

    export let task: TaskData;
    export let edit: boolean = false;

    const editStart = createEventDispatcher()
    const editSave = createEventDispatcher()
    const editDone = createEventDispatcher()
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<task on:click={()=>editStart("editStart")}>
    {#if !edit}
        <h3>{task.title}</h3>
        <p>{task.notice}</p>
    {:else}
        <input type="text" bind:value={task.title}/>
        <textarea rows="5" bind:value={task.notice}/>
        <Button3 span={1} on:click={()=>editDone("editDone")}>Erledigen</Button3>
        <Button3 span={1} on:click={()=>editSave("editSave")}>Speichern</Button3>
    {/if}
</task>

<style>
    task {
        grid-column: span 2;
        display: grid;
        grid-template-columns: 50%, 50%;
        grid-gap: .5rem;
        padding: .5rem;
        background-color: var(--COLOR-2);
    }
    input, textarea {
        grid-column: span 2;
        background-color: var(--COLOR-1);
        color: var(--COLOR-4);
        font-size: 1rem;
        font-family: inherit;
        padding: .5rem;
        outline: none;
        border: none;
    }
    
    input {
        width: 100%;
        
    }

    textarea {
        width: 100%;
        resize: none;
    }
</style>